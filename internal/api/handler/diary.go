package handler

import (
	"github.com/gin-gonic/gin"
	"module.resume/internal/application"
)

type DiaryHandler struct {
	service application.AuthService
}

func NewDiaryHandler(service application.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

func (d *DiaryHandler) Save(c *gin.Context) {
}

func (d *DiaryHandler) Update(c *gin.Context) {
}

func (d *DiaryHandler) Delete(c *gin.Context) {
}
