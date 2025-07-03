package application

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"module.resume/internal/domain/user"
	"module.resume/internal/util"
)

const testSecret = "test-secret"

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func generateTestToken(t *testing.T, email string, secret string, expiresAt time.Time) string {
	claims := jwt.MapClaims{
		"sub": email,
		"iat": time.Now().Unix(),
		"exp": expiresAt.Unix(),
		"iss": "module-resume-server",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	assert.NoError(t, err)
	return signedToken
}

func TestAuthService_Login(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockCache := new(MockCache)
	authService := NewAuthService(mockUserRepo, mockCache, testSecret)
	ctx := context.Background()

	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := util.HashPassword(password)

	loginAttemptUser := &user.User{Email: email, Password: password}

	storedUser := &user.User{ID: 1, Email: email}
	storedUser.SetPasswordHash(hashedPassword)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", ctx, email).Return(storedUser, nil).Once()

		token, err := authService.Login(ctx, loginAttemptUser)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", ctx, email).Return(nil, errors.New("user not found")).Once()

		token, err := authService.Login(ctx, loginAttemptUser)

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, "user not found", err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		wrongPasswordUser := &user.User{Email: email, Password: "wrong-password"}
		mockUserRepo.On("FindByEmail", ctx, email).Return(storedUser, nil).Once()

		token, err := authService.Login(ctx, wrongPasswordUser)

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, "invalid password", err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestAuthService_Logout(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockCache := new(MockCache)
	authService := NewAuthService(mockUserRepo, mockCache, testSecret)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		expiresAt := time.Now().Add(time.Hour)
		token := generateTestToken(t, "test@example.com", testSecret, expiresAt)
		remainingTime := time.Until(expiresAt)

		mockCache.On("Set", ctx, "blocklist:"+token, "true", mock.AnythingOfType("time.Duration")).Return(nil).Once()

		err := authService.Logout(ctx, token)

		assert.NoError(t, err)

		args := mockCache.Calls[0].Arguments
		duration := args.Get(3).(time.Duration)
		assert.InDelta(t, remainingTime, duration, float64(time.Second))
		mockCache.AssertExpectations(t)
	})
}

func TestAuthService_Authenticate(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockCache := new(MockCache)
	authService := NewAuthService(mockUserRepo, mockCache, testSecret)
	ctx := context.Background()
	email := "user@example.com"

	t.Run("success", func(t *testing.T) {
		token := generateTestToken(t, email, testSecret, time.Now().Add(time.Hour))
		mockCache.On("Get", ctx, "blocklist:"+token).Return("", nil).Once()

		claims, err := authService.Authenticate(ctx, token)

		assert.NoError(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, email, claims.Subject)
		mockCache.AssertExpectations(t)
	})

	t.Run("token is blocklisted", func(t *testing.T) {
		token := generateTestToken(t, email, testSecret, time.Now().Add(time.Hour))
		mockCache.On("Get", ctx, "blocklist:"+token).Return("true", nil).Once()

		claims, err := authService.Authenticate(ctx, token)

		assert.Error(t, err)
		assert.Nil(t, claims)
		assert.Equal(t, "token is blocklisted", err.Error())
		mockCache.AssertExpectations(t)
	})

	t.Run("invalid token - bad signature", func(t *testing.T) {
		token := generateTestToken(t, email, "wrong-secret", time.Now().Add(time.Hour))
		mockCache.On("Get", ctx, "blocklist:"+token).Return("", nil).Once()

		claims, err := authService.Authenticate(ctx, token)

		assert.Error(t, err)
		assert.Nil(t, claims)
		mockCache.AssertExpectations(t)
	})

	t.Run("invalid token - expired", func(t *testing.T) {
		token := generateTestToken(t, email, testSecret, time.Now().Add(-time.Hour))
		mockCache.On("Get", ctx, "blocklist:"+token).Return("", nil).Once()

		claims, err := authService.Authenticate(ctx, token)

		assert.Error(t, err)
		assert.Nil(t, claims)
		assert.Contains(t, err.Error(), "token is expired")
		mockCache.AssertExpectations(t)
	})
}
