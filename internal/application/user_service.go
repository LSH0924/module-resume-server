package application

import (
	"context"

	"module.resume/internal/domain/user"
)

type UserService interface {
	Save(context context.Context, user *user.User) (int64, error)
	Update(context context.Context, user *user.User) (int64, error)
	Delete(context context.Context, user *user.User) error
}

type userService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) UserService {
	return &userService{
		repo,
	}
}

func (service *userService) Save(context context.Context, user *user.User) (int64, error) {
	return service.repo.Save(context, user)
}

func (service *userService) Update(context context.Context, user *user.User) (int64, error) {
	return service.repo.Update(context, user)
}

func (service *userService) Delete(context context.Context, user *user.User) error {
	return service.repo.Delete(context, user.ID)
}
