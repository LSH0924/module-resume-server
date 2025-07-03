package gorm

import (
	"context"

	"gorm.io/gorm"
	"module.resume/internal/domain/user"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	user := &User{}
	result := r.db.WithContext(ctx).Where("email = ?", email).First(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user.toDomain(), nil
}

func (r *UserRepository) Save(ctx context.Context, user *user.User) (uint, error) {
	gormUser := fromDomain(user)
	err := r.db.WithContext(ctx).Create(gormUser).Error
	if err != nil {
		return 0, err
	}
	return gormUser.ID, nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) (uint, error) {
	gormUser := fromDomain(user)
	err := r.db.WithContext(ctx).Where("id = ?", user.ID).Updates(gormUser).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&User{}, id).Error
}
