package user

import (
	"time"

	"module.resume/internal/util"
)

type User struct {
	ID           uint
	Email        string
	Name         string
	Password     string
	passwordHash string
	ProfileUrl   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func NewUserForSave(email, name, plainPassword, profileUrl string) (*User, error) {
	hashedPassword, err := util.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:        email,
		Name:         name,
		passwordHash: hashedPassword,
		ProfileUrl:   profileUrl,
	}, nil
}

func NewUserForUpdate(id uint, email, name, profileUrl string) *User {
	return &User{
		ID:         id,
		Email:      email,
		Name:       name,
		ProfileUrl: profileUrl,
	}
}

func NewUserForLogin(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func Hydrate(id uint, email, name, hashedPassword string, createdAt, updatedAt time.Time) *User {
	return &User{
		ID:           id,
		Email:        email,
		Name:         name,
		passwordHash: hashedPassword,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
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
