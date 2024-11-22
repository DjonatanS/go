package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Djonatan", "123456", "djonatan@gmail.com")
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, user.Username, "Djonatan")
	assert.Equal(t, user.Email, "djonatan@gmail.com")
	assert.NotNil(t, user.Username)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidadePassword(t *testing.T) {
	user, err := NewUser("Djonatan", "123456", "djonatan@gmail.com")
	if err != nil {
		panic(err)
	}
	assert.True(t, user.ValidadePassword("123456"))
	assert.False(t, user.ValidadePassword("654321"))
	assert.NotEqual(t, user.Password, "123456")
}
