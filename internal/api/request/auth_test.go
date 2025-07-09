package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"module.resume/internal/domain/user"
)

func TestLoginRequest_ToDomain(t *testing.T) {
	t.Run("LoginRequest를 user.User 도메인 객체로 변환", func(t *testing.T) {
		req := LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		domainUser := req.ToDomain()

		expectedUser := &user.User{
			Email:    "test@example.com",
			Password: "password123",
		}
		assert.Equal(t, expectedUser, domainUser)
	})
}
