package gorm

import (
	"time"

	"gorm.io/gorm"
	"module.resume/internal/domain/user"
)

type User struct {
	gorm.Model
	Email        string `gorm:"column:email;not null"`
	Name         string `gorm:"column:name;not null"`
	PasswordHash string `gorm:"column:password_hash;not null"`
	ProfileUrl   string `gorm:"column:profile_url"`
}

func (User) TableName() string {
	return "user"
}

func (m User) toDomain() *user.User {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	domainUser := &user.User{
		ID:        m.ID,
		Email:     m.Email,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAt,
	}
	domainUser.SetPasswordHash(m.PasswordHash)
	return domainUser
}

func fromDomain(u *user.User) *User {
	return &User{
		Email:        u.Email,
		Name:         u.Name,
		PasswordHash: u.PasswordHash(),
		ProfileUrl:   u.ProfileUrl,
	}
}
