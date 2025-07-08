package user

import "context"

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Delete(ctx context.Context, id int64) error
}
