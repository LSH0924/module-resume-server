package user

import (
	"time"

	"module.resume/internal/util"
)

type User struct {
	ID           int64
	Email        string
	Name         string
	Password     string
	passwordHash string
	ProfileURL   *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func NewUserForSave(email, name, plainPassword string, profileUrl *string) (*User, error) {
	hashedPassword, err := util.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:        email,
		Name:         name,
		passwordHash: hashedPassword,
		ProfileURL:   profileUrl,
	}, nil
}

func (u *User) CheckPassword(plainPassword string) bool {
	return util.CheckPasswordHash(plainPassword, u.passwordHash)
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}
