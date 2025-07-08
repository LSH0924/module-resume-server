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
	ProfileURL   string
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
		ProfileURL:   profileUrl,
	}, nil
}

func NewUserForUpdate(id int64, email, name, profileUrl string) *User {
	return &User{
		ID:         id,
		Email:      email,
		Name:       name,
		ProfileURL: profileUrl,
	}
}

func NewUserForDelete(id int64) *User {
	return &User{
		ID: id,
	}
}

func NewUserForLogin(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func Hydrate(id int64, email, name, hashedPassword string, createdAt, updatedAt time.Time, deletedAt *time.Time) *User {
	return &User{
		ID:           id,
		Email:        email,
		Name:         name,
		passwordHash: hashedPassword,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
	}
}

func (u *User) CheckPassword(plainPassword string) bool {
	return util.CheckPasswordHash(plainPassword, u.passwordHash)
}

func (u *User) Int64ID() int64 {
	return int64(u.ID)
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}
