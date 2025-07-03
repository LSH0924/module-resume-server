package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenExtractorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.Next()
			return
		}
		c.Set("token", token)
		c.Next()
	}
}
