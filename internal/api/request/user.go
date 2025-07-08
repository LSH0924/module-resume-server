package request

import (
	"module.resume/internal/domain/user"
)

type SaveUser struct {
	Email      string `json:"email" binding:"required,email"`
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required,min=12"`
	ProfileUrl string `json:"profileUrl" binding:"url"`
}

func (s SaveUser) ToDomain() (*user.User, error) {
	domain, err := user.NewUserForSave(s.Email, s.Name, s.Password, s.ProfileUrl)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

type UpdateUser struct {
	ID         int64  `json:"id" binding:"required,min=1"`
	Email      string `json:"email" binding:"email"`
	Name       string `json:"name"`
	ProfileUrl string `json:"profileUrl" binding:"url"`
}

func (u UpdateUser) ToDomain() *user.User {
	return user.NewUserForUpdate(u.ID, u.Email, u.Name, u.ProfileUrl)
}

type UpdateUserPassword struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=newPassword"`
}

type DeleteUser struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (d DeleteUser) ToDomain() *user.User {
	return user.NewUserForDelete(d.ID)
}
