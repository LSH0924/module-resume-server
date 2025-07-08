package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"module.resume/internal/api/request"
	"module.resume/internal/api/response"
	"module.resume/internal/application"
)

type AuthHandler struct {
	service application.AuthService
}

func NewAuthHandler(service application.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

func (a *AuthHandler) Login(c *gin.Context) {
	loginRequest := &request.LoginRequest{}
	if err := c.ShouldBindJSON(loginRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	token, err := a.service.Login(c.Request.Context(), loginRequest.ToDomain())
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	result := response.LoginResponse{AccessToken: token}
	c.JSON(http.StatusOK, result)
}

func (a *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	err := a.service.Logout(c.Request.Context(), tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, nil)
}
