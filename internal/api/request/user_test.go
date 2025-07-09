package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"module.resume/internal/domain/user"
)

func TestSaveUser_ToDomain(t *testing.T) {
	t.Run("SaveUser를 user.User 도메인 객체로 성공적으로 변환", func(t *testing.T) {
		profileURL := "http://example.com/profile.jpg"
		req := SaveUser{
			Email:      "test@example.com",
			Name:       "Test User",
			Password:   "a-strong-password-123",
			ProfileUrl: &profileURL,
		}

		domainUser, err := req.ToDomain()

		assert.NoError(t, err)
		assert.NotNil(t, domainUser)
		assert.Equal(t, req.Email, domainUser.Email)
		assert.Equal(t, req.Name, domainUser.Name)
		assert.NotEqual(t, req.Password, domainUser.Password)
		assert.Equal(t, req.ProfileUrl, domainUser.ProfileURL)
	})

	t.Run("잘못된 사용자 데이터에 대해 오류를 반환", func(t *testing.T) {
		req := SaveUser{
			Email:    "test@example.com",
			Name:     "Test User",
			Password: "veryLongPassword-veryLongPassword-12345!veryLongPassword-veryLongPassword-12345!",
		}

		domainUser, err := req.ToDomain()

		assert.Error(t, err)
		assert.Nil(t, domainUser)
	})
}

func TestUpdateUser_ToDomain(t *testing.T) {
	t.Run("UpdateUser를 user.User 도메인 객체로 변환", func(t *testing.T) {
		profileURL := "http://example.com/new-profile.jpg"
		req := UpdateUser{
			ID:         1,
			Email:      "new.email@example.com",
			Name:       "New Name",
			ProfileURL: &profileURL,
		}

		domainUser := req.ToDomain()

		expectedUser := &user.User{
			ID:         1,
			Email:      "new.email@example.com",
			Name:       "New Name",
			ProfileURL: &profileURL,
		}
		assert.Equal(t, expectedUser, domainUser)
	})
}

func TestDeleteUser_ToDomain(t *testing.T) {
	t.Run("DeleteUser를 user.User 도메인 객체로 변환", func(t *testing.T) {
		req := DeleteUser{
			ID: 1,
		}

		domainUser := req.ToDomain()

		expectedUser := &user.User{
			ID: 1,
		}
		assert.Equal(t, expectedUser, domainUser)
	})
}
