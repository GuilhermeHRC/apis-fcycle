package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "test@test.com", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "test@test.com", user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "test@test.com", "password123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, user.Password, "password123")
}
