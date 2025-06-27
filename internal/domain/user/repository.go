package user

import "context"

type Repository interface {
	Save(ctx context.Context, user *User) (uint, error)
	Update(ctx context.Context, user *User) (uint, error)
	Delete(ctx context.Context, id uint) error
}
