package container

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"module.resume/internal/api"
	"module.resume/internal/api/handler"
	"module.resume/internal/api/middleware"
	"module.resume/internal/application"
	"module.resume/internal/infrastructure/cache"
	"module.resume/internal/infrastructure/persistence/gorm"
)

type Container struct {
	Router *gin.Engine
}

func NewContainer() (*Container, error) {
	db, err := gorm.NewDB()
	if err != nil {
		return nil, err
	}

	userRepo := gorm.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	redis, err := cache.NewRedisClient()
	if err != nil {
		return nil, err
	}
	cache := cache.NewRedisCache(redis)
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return nil, errors.New("JWT secret key not set")
	}
	authService := application.NewAuthService(userRepo, cache, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleWare := middleware.AuthMiddleware(authService)

	h := &handler.Handlers{
		User: userHandler,
		Auth: authHandler,
	}

	r := api.MakeRouter(h, authMiddleWare)

	return &Container{
		Router: r,
	}, nil
}
