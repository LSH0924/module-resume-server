package mapper

import (
	"time"

	"module.resume/internal/domain/diary"
	"module.resume/internal/infrastructure/persistence/gorm/model"
)

func ToDomainDiary(dbModel *model.Diary) *diary.Diary {
	var deletedAt *time.Time
	if dbModel.DeletedAt.Valid {
		deletedAt = &dbModel.DeletedAt.Time
	}

	domainDiary := &diary.Diary{
		ID:                dbModel.ID,
		UsersID:           dbModel.UsersID,
		WorkExperiencesID: &dbModel.WorkExperiencesID,
		TypesID:           dbModel.TypesID,
		ProjectsID:        &dbModel.ProjectsID,
		Content:           dbModel.Content,
		Mood:              &dbModel.Mood,
		CreatedAt:         dbModel.CreatedAt,
		UpdatedAt:         dbModel.UpdatedAt,
		DeletedAt:         deletedAt,
	}
	return domainDiary
}

func ToDBDiary(diary *diary.Diary) *model.Diary {
	dbDiary := &model.Diary{}
	dbDiary.ID = diary.ID
	dbDiary.UsersID = diary.UsersID
	dbDiary.TypesID = diary.TypesID
	if diary.WorkExperiencesID != nil {
		dbDiary.WorkExperiencesID = *diary.WorkExperiencesID
	}
	if diary.ProjectsID != nil {
		dbDiary.ProjectsID = *diary.ProjectsID
	}
	dbDiary.Content = diary.Content
	if diary.Mood != nil {
		dbDiary.Mood = *diary.Mood
	}
	dbDiary.CreatedAt = diary.CreatedAt
	dbDiary.UpdatedAt = diary.UpdatedAt
	return dbDiary
}
