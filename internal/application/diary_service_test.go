package application

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"module.resume/internal/domain/diary"
)

type MockDiaryRepository struct {
	mock.Mock
}

func (m *MockDiaryRepository) Save(ctx context.Context, d *diary.Diary) (int64, error) {
	args := m.Called(ctx, d)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockDiaryRepository) Update(ctx context.Context, d *diary.Diary) (int64, error) {
	args := m.Called(ctx, d)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockDiaryRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestDiaryService_Save(t *testing.T) {
	mockRepo := new(MockDiaryRepository)
	diaryService := NewDiaryService(mockRepo)
	ctx := context.Background()
	testDiary := &diary.Diary{
		UsersID: 1,
		TypesID: 1,
		Content: "Diary!",
	}

	t.Run("성공", func(t *testing.T) {
		mockRepo.On("Save", ctx, testDiary).Return(1, nil).Once()

		id, err := diaryService.Save(ctx, testDiary)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("실패", func(t *testing.T) {
		mockRepo.On("Save", ctx, testDiary).Return(0, errors.New("db error")).Once()

		id, err := diaryService.Save(ctx, testDiary)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		assert.Equal(t, int64(0), id)
		mockRepo.AssertExpectations(t)
	})
}

func TestDiaryService_Update(t *testing.T) {
	mockRepo := new(MockDiaryRepository)
	diaryService := NewDiaryService(mockRepo)
	ctx := context.Background()
	testDiary := &diary.Diary{
		ID:      1,
		UsersID: 1,
		TypesID: 1,
		Content: "Diary!",
	}

	t.Run("성공", func(t *testing.T) {
		mockRepo.On("Update", ctx, testDiary).Return(1, nil).Once()

		id, err := diaryService.Update(ctx, testDiary)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("실패", func(t *testing.T) {
		mockRepo.On("Update", ctx, testDiary).Return(0, errors.New("update failed")).Once()

		id, err := diaryService.Update(ctx, testDiary)

		assert.Error(t, err)
		assert.Equal(t, "update failed", err.Error())
		assert.Equal(t, int64(0), id)
		mockRepo.AssertExpectations(t)
	})
}

func TestDiaryService_Delete(t *testing.T) {
	mockRepo := new(MockDiaryRepository)
	diaryService := NewDiaryService(mockRepo)
	ctx := context.Background()
	testDiary := &diary.Diary{
		ID:      1,
		UsersID: 1,
		TypesID: 1,
		Content: "Diary!",
	}

	t.Run("성공", func(t *testing.T) {
		mockRepo.On("Delete", ctx, testDiary.ID).Return(nil).Once()

		err := diaryService.Delete(ctx, testDiary)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("실패", func(t *testing.T) {
		mockRepo.On("Delete", ctx, testDiary.ID).Return(errors.New("delete failed")).Once()

		err := diaryService.Delete(ctx, testDiary)

		assert.Error(t, err)
		assert.Equal(t, "delete failed", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
