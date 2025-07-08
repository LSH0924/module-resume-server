package gorm

import (
	"context"

	"module.resume/internal/domain/diary"
	"module.resume/internal/infrastructure/persistence/gorm/mapper"
	"module.resume/internal/infrastructure/persistence/gorm/query"
)

type DiaryRepository struct {
	query *query.Query
}

func NewDiaryRepository(query *query.Query) diary.Repository {
	return &DiaryRepository{query}
}

func (r *DiaryRepository) Save(ctx context.Context, diary *diary.Diary) (int64, error) {
	dbDiary := mapper.ToDBDiary(diary)
	if err := r.query.Diary.WithContext(ctx).Create(dbDiary); err != nil {
		return 0, err
	}
	return dbDiary.ID, nil
}

func (r *DiaryRepository) Update(ctx context.Context, diary *diary.Diary) (int64, error) {
	dbDiary := mapper.ToDBDiary(diary)
	d := r.query.Diary
	_, err := d.WithContext(ctx).Where(d.ID.Eq(diary.ID)).Updates(dbDiary)
	if err != nil {
		return 0, err
	}
	return diary.ID, nil
}

func (r *DiaryRepository) Delete(ctx context.Context, id int64) error {
	d := r.query.Diary
	_, err := d.WithContext(ctx).Where(d.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}
