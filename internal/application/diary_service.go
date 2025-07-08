package application

import (
	"context"

	"module.resume/internal/domain/diary"
)

type DiaryService interface {
	Save(ctx context.Context, diary *diary.Diary) (int64, error)
	Update(ctx context.Context, diary *diary.Diary) (int64, error)
	Delete(ctx context.Context, diary *diary.Diary) error
}

type diaryService struct {
	repo diary.Repository
}

func NewDiaryService(repo diary.Repository) DiaryService {
	return &diaryService{repo}
}

func (d diaryService) Save(ctx context.Context, diary *diary.Diary) (int64, error) {
	return d.repo.Save(ctx, diary)
}

func (d diaryService) Update(ctx context.Context, diary *diary.Diary) (int64, error) {
	return d.repo.Update(ctx, diary)
}

func (d diaryService) Delete(ctx context.Context, diary *diary.Diary) error {
	return d.repo.Delete(ctx, diary.ID)
}
