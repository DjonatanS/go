package entity

import (
	"api/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

func NewUser(username, password, email string) (*User, error) {
	// Gerenate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}, nil
}

func (u *User) ValidadePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
