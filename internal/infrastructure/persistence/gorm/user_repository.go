package gorm

import (
	"context"

	"module.resume/internal/domain/user"
	"module.resume/internal/infrastructure/persistence/gorm/mapper"
	"module.resume/internal/infrastructure/persistence/gorm/query"
)

type UserRepository struct {
	query *query.Query
}

func NewUserRepository(query *query.Query) user.Repository {
	return &UserRepository{query}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	u := r.query.User
	result, err := u.WithContext(ctx).Where(u.Email.Eq(email)).First()
	if err != nil {
		return nil, err
	}
	return mapper.ToDomainUser(result), nil
}

func (r *UserRepository) Save(ctx context.Context, user *user.User) (int64, error) {
	dbUser := mapper.ToDBUser(user)
	if err := r.query.User.WithContext(ctx).Create(dbUser); err != nil {
		return 0, err
	}
	return dbUser.ID, nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) (int64, error) {
	dbUser := mapper.ToDBUser(user)
	u := r.query.User
	_, err := u.WithContext(ctx).Where(u.ID.Eq(user.ID)).Updates(dbUser)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	u := r.query.User
	_, err := u.WithContext(ctx).Where(u.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}
