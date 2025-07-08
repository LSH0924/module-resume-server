package mapper

import (
	"time"

	"module.resume/internal/domain/user"
	"module.resume/internal/infrastructure/persistence/gorm/model"
)

func ToDomainUser(dbModel *model.User) *user.User {
	var deletedAt *time.Time
	if dbModel.DeletedAt.Valid {
		deletedAt = &dbModel.DeletedAt.Time
	}

	domainUser := &user.User{
		ID:        dbModel.ID,
		Email:     dbModel.Email,
		Name:      dbModel.Name,
		CreatedAt: dbModel.CreatedAt,
		UpdatedAt: dbModel.UpdatedAt,
		DeletedAt: deletedAt,
	}
	domainUser.SetPasswordHash(dbModel.PasswordHash)
	return domainUser
}

func ToDBUser(user *user.User) *model.User {
	dbUser := &model.User{}
	dbUser.ID = user.Int64ID()
	dbUser.Email = user.Email
	dbUser.Name = user.Name
	dbUser.PasswordHash = user.PasswordHash()
	dbUser.ProfileURL = user.ProfileURL
	dbUser.CreatedAt = user.CreatedAt
	dbUser.UpdatedAt = user.UpdatedAt
	return dbUser
}
