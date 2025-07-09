package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserForSave(t *testing.T) {
	email := "test@example.com"
	name := "Test User"
	password := "password123"
	profileURL := "http://example.com/profile.jpg"

	t.Run("성공", func(t *testing.T) {
		user, err := NewUserForSave(email, name, password, &profileURL)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, &profileURL, user.ProfileURL)
		assert.NotEmpty(t, user.PasswordHash())
		assert.NotEqual(t, password, user.PasswordHash())

		assert.True(t, user.CheckPassword(password))
	})
}

func TestUser_CheckPassword(t *testing.T) {
	password := "my-secure-password"
	profileUrl := "http://example.com/profile.jpg"
	user, err := NewUserForSave("user@test.com", "User", password, &profileUrl)
	assert.NoError(t, err)

	t.Run("올바른 비밀번호", func(t *testing.T) {
		assert.True(t, user.CheckPassword(password))
	})

	t.Run("잘못된 비밀번호", func(t *testing.T) {
		assert.False(t, user.CheckPassword("wrong-password"))
	})

	t.Run("빈 비밀번호", func(t *testing.T) {
		assert.False(t, user.CheckPassword(""))
	})
}

func TestUser_SetPasswordHash(t *testing.T) {
	user := &User{}
	hash := "a-new-hash"
	user.SetPasswordHash(hash)
	assert.Equal(t, hash, user.PasswordHash())
}
