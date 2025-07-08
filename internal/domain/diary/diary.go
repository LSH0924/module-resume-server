package diary

import (
	"time"
)

type Diary struct {
	ID                int64
	TypesID           int64
	UsersID           int64
	WorkExperiencesID *int64
	ProjectsID        *int64
	Content           string
	Mood              *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
