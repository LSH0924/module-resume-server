package request

import (
	"module.resume/internal/domain/user"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (l LoginRequest) ToDomain() *user.User {
	return user.NewUserForLogin(l.Email, l.Password)
}
