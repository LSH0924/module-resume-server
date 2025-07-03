package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"module.resume/internal/api/handler"
	"module.resume/internal/api/middleware"
)

func MakeRouter(handlers *handler.Handlers, authMiddleware gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.TokenExtractorMiddleware())
	r.Use(middleware.TimeoutMiddleware(10 * time.Second))

	user := r.Group("/user")
	{
		user.POST("/", handlers.User.Save)
		user.DELETE("/:id", handlers.User.Delete)
		me := user.Group("/me")
		{
			me.Use(authMiddleware)
			me.PUT("/", handlers.User.Update)
			me.PUT("/password", handlers.User.UpdatePassword)
		}
	}

	{
		r.POST("/login", handlers.Auth.Login)
		r.POST("/logout", handlers.Auth.Logout)
	}

	return r
}
