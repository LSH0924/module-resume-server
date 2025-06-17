package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from api/router.go!",
		})
	})

	return r
}
