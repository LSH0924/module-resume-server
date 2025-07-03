package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUserForSave(t *testing.T) {
	email := "test@example.com"
	name := "Test User"
	password := "password123"
	profileURL := "http://example.com/profile.jpg"

	t.Run("success", func(t *testing.T) {
		user, err := NewUserForSave(email, name, password, profileURL)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, profileURL, user.ProfileUrl)
		assert.NotEmpty(t, user.PasswordHash())
		assert.NotEqual(t, password, user.PasswordHash())

		assert.True(t, user.CheckPassword(password))
	})
}

func TestUser_CheckPassword(t *testing.T) {
	password := "my-secure-password"
	user, err := NewUserForSave("user@test.com", "User", password, "")
	assert.NoError(t, err)

	t.Run("correct password", func(t *testing.T) {
		assert.True(t, user.CheckPassword(password))
	})

	t.Run("incorrect password", func(t *testing.T) {
		assert.False(t, user.CheckPassword("wrong-password"))
	})

	t.Run("empty password", func(t *testing.T) {
		assert.False(t, user.CheckPassword(""))
	})
}

func TestNewUserForUpdate(t *testing.T) {
	id := uint(1)
	email := "update@example.com"
	name := "Updated User"
	profileURL := "http://example.com/new.jpg"

	user := NewUserForUpdate(id, email, name, profileURL)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, profileURL, user.ProfileUrl)
}

func TestNewUserForLogin(t *testing.T) {
	email := "login@example.com"
	password := "password123"

	user := NewUserForLogin(email, password)

	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
}

func TestHydrate(t *testing.T) {
	id := uint(1)
	email := "hydrated@example.com"
	name := "Hydrated User"
	hashedPassword := "hashed_password_string"
	now := time.Now()

	user := Hydrate(id, email, name, hashedPassword, now, now)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, hashedPassword, user.PasswordHash())
	assert.Equal(t, now, user.CreatedAt)
	assert.Equal(t, now, user.UpdatedAt)
}

func TestUser_SetPasswordHash(t *testing.T) {
	user := &User{}
	hash := "a-new-hash"
	user.SetPasswordHash(hash)
	assert.Equal(t, hash, user.PasswordHash())
}
