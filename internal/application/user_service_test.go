package application

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"module.resume/internal/domain/user"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserRepository) Save(ctx context.Context, u *user.User) (uint, error) {
	args := m.Called(ctx, u)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, u *user.User) (uint, error) {
	args := m.Called(ctx, u)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserService_Save(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)
	ctx := context.Background()
	testUser := &user.User{Email: "test@example.com", Password: "password"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Save", ctx, testUser).Return(1, nil).Once()

		id, err := userService.Save(ctx, testUser)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("Save", ctx, testUser).Return(0, errors.New("db error")).Once()

		id, err := userService.Save(ctx, testUser)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		assert.Equal(t, uint(0), id)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_Update(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)
	ctx := context.Background()
	testUser := &user.User{ID: 1, Email: "update@example.com"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Update", ctx, testUser).Return(1, nil).Once()

		id, err := userService.Update(ctx, testUser)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("Update", ctx, testUser).Return(0, errors.New("update failed")).Once()

		id, err := userService.Update(ctx, testUser)

		assert.Error(t, err)
		assert.Equal(t, "update failed", err.Error())
		assert.Equal(t, uint(0), id)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_Delete(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)
	ctx := context.Background()
	testUser := &user.User{ID: 1}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", ctx, testUser.ID).Return(nil).Once()

		err := userService.Delete(ctx, testUser)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("Delete", ctx, testUser.ID).Return(errors.New("delete failed")).Once()

		err := userService.Delete(ctx, testUser)

		assert.Error(t, err)
		assert.Equal(t, "delete failed", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
