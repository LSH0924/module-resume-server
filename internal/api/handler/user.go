package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"module.resume/internal/api/request"
	"module.resume/internal/application"
)

type UserHandler struct {
	service application.UserService
}

func NewUserHandler(service application.UserService) *UserHandler {
	return &UserHandler{
		service,
	}
}

func (h *UserHandler) Save(c *gin.Context) {
	requestUser := request.SaveUser{}
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	domainUser, err := requestUser.ToDomain()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	userId, err := h.service.Save(c.Request.Context(), domainUser)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Database operation timed out",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userId)
}

func (h *UserHandler) Update(c *gin.Context) {
	requestUser := request.UpdateUser{}
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	domainUser := requestUser.ToDomain()
	userId, err := h.service.Update(c.Request.Context(), domainUser)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Database operation timed out",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, userId)
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
}

func (h *UserHandler) Delete(c *gin.Context) {
}
