package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"module.resume/internal/application"
)

func AuthMiddleware(authService application.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, exists := c.Get("token")
		if !exists || token.(string) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			return
		}

		claims, err := authService.Authenticate(c, utils.ToString(token))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set("email", claims.Subject)
		c.Next()
	}
}
