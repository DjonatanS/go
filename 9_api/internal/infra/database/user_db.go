package database

import (
	"api/internal/entity"
	"database/sql"
	"fmt"
)

type User struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	stmt, err := u.DB.Prepare("INSERT INTO users(id, username, password, email) VALUES(?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Username, user.Password, user.Email)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	stmt, err := u.DB.Prepare("SELECT id, username, password, email FROM users WHERE email = ?")
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	var user entity.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil user if not found
		}
		return nil, fmt.Errorf("error scanning result: %w", err)
	}

	return &user, nil
}
