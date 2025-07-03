package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"module.resume/internal/auth"
	"module.resume/internal/domain/user"
)

type AuthService interface {
	Login(context context.Context, user *user.User) (string, error)
	Logout(context context.Context, token string) error
	Authenticate(ctx context.Context, token string) (*auth.Claims, error)
}

type authService struct {
	userRepo user.Repository
	cache    Cache
	secret   []byte
}

func NewAuthService(userRepo user.Repository, cache Cache, secret string) AuthService {
	return &authService{
		userRepo: userRepo,
		cache:    cache,
		secret:   []byte(secret),
	}
}

func (a *authService) Login(context context.Context, user *user.User) (string, error) {
	storedUser, err := a.userRepo.FindByEmail(context, user.Email)
	if err != nil {
		return "", err
	}
	matched := storedUser.CheckPassword(user.Password)
	if !matched {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"sub": storedUser.Email,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iss": "module-resume-server",
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(a.secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authService) Logout(context context.Context, token string) error {
	claims, err := a.parseToken(token)

	var remainingTime time.Duration
	if err != nil {
		remainingTime = 24 * time.Hour
	} else {
		remainingTime = time.Until(claims.ExpiresAt.Time)
	}

	if remainingTime < 0 {
		return nil
	}

	redisErr := a.cache.Set(context, "blocklist:"+token, "true", remainingTime)
	if redisErr != nil {
		return redisErr
	}

	return nil
}

func (a *authService) Authenticate(ctx context.Context, token string) (*auth.Claims, error) {
	key := "blocklist:" + token
	val, err := a.cache.Get(ctx, key)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if val != "" {
		return nil, errors.New("token is blocklisted")
	}

	return a.parseToken(token)
}

func (a *authService) parseToken(tokenString string) (*auth.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
