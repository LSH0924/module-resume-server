package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"module.resume/internal/domain/diary"
)

func TestSaveDiary_ToDomain(t *testing.T) {
	t.Run("SaveDiary를 diary.Diary 도메인 객체로 변환", func(t *testing.T) {
		workExperiencesID := int64(1)
		projectsID := int64(1)
		mood := "happy"
		req := SaveDiary{
			UsersID:           1,
			TypesID:           1,
			WorkExperiencesID: &workExperiencesID,
			ProjectsID:        &projectsID,
			Content:           "This is a test diary entry.",
			Mood:              &mood,
		}

		domainDiary := req.ToDomain()

		expectedDiary := &diary.Diary{
			UsersID:           1,
			TypesID:           1,
			WorkExperiencesID: &workExperiencesID,
			ProjectsID:        &projectsID,
			Content:           "This is a test diary entry.",
			Mood:              &mood,
		}
		assert.Equal(t, expectedDiary, domainDiary)
	})
}

func TestUpdateDiary_ToDomain(t *testing.T) {
	t.Run("UpdateDiary를 diary.Diary 도메인 객체로 변환", func(t *testing.T) {
		workExperiencesID := int64(2)
		projectsID := int64(2)
		mood := "excited"
		req := UpdateDiary{
			ID:                1,
			UsersID:           1,
			TypesID:           2,
			WorkExperiencesID: &workExperiencesID,
			ProjectsID:        &projectsID,
			Content:           "This is an updated diary entry.",
			Mood:              &mood,
		}

		domainDiary := req.ToDomain()

		expectedDiary := &diary.Diary{
			ID:                1,
			UsersID:           1,
			TypesID:           2,
			WorkExperiencesID: &workExperiencesID,
			ProjectsID:        &projectsID,
			Content:           "This is an updated diary entry.",
			Mood:              &mood,
		}
		assert.Equal(t, expectedDiary, domainDiary)
	})
}

func TestDeleteDiary_ToDomain(t *testing.T) {
	t.Run("DeleteDiary를 diary.Diary 도메인 객체로 변환", func(t *testing.T) {
		req := DeleteDiary{
			ID: 1,
		}

		domainDiary := req.ToDomain()

		expectedDiary := &diary.Diary{
			ID: 1,
		}
		assert.Equal(t, expectedDiary, domainDiary)
	})
}
