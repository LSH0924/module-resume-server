package container

import (
	"github.com/gin-gonic/gin"
	"module.resume/internal/api"
	"module.resume/internal/api/handler"
	"module.resume/internal/application"
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

	h := &handler.Handlers{
		User: userHandler,
	}

	r := api.MakeRouter(h)

	return &Container{
		Router: r,
	}, nil
}
