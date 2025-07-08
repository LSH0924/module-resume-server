package request

import (
	"module.resume/internal/domain/diary"
)

type SaveDiary struct {
	UsersID           int64   `json:"usersId" binding:"required,min=1"`
	TypesID           int64   `json:"typesId" binding:"required,min=1"`
	WorkExperiencesID *int64  `json:"workExperiencesId"`
	ProjectsID        *int64  `json:"projectsId"`
	Content           string  `json:"content" binding:"required"`
	Mood              *string `json:"mood" binding:"max=255"`
}

func (s SaveDiary) ToDomain() *diary.Diary {
	return &diary.Diary{
		UsersID:           s.UsersID,
		TypesID:           s.TypesID,
		WorkExperiencesID: s.WorkExperiencesID,
		ProjectsID:        s.ProjectsID,
		Content:           s.Content,
		Mood:              s.Mood,
	}
}

type UpdateDiary struct {
	ID                int64   `json:"id" binding:"required,min=1"`
	UsersID           int64   `json:"usersId" binding:"required,min=1"`
	TypesID           int64   `json:"typesId" binding:"required,min=1"`
	WorkExperiencesID *int64  `json:"workExperiencesId"`
	ProjectsID        *int64  `json:"projectsId"`
	Content           string  `json:"content" binding:"required"`
	Mood              *string `json:"mood" binding:"max=255"`
}

func (u UpdateDiary) ToDomain() *diary.Diary {
	return &diary.Diary{
		ID:                u.ID,
		UsersID:           u.UsersID,
		TypesID:           u.TypesID,
		WorkExperiencesID: u.WorkExperiencesID,
		ProjectsID:        u.ProjectsID,
		Content:           u.Content,
		Mood:              u.Mood,
	}
}

type DeleteDiary struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (u DeleteDiary) ToDomain() *diary.Diary {
	return &diary.Diary{
		ID: u.ID,
	}
}
