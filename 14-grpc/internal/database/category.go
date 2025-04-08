package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		DB: db,
	}
}

func (c *Category) Create(name string, description string) (Category, error) {
	// Add nil check to prevent nil pointer dereference
	if c.DB == nil {
		return Category{}, errors.New("database connection is nil")
	}

	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)",
		id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.DB.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.DB.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = ?", courseID).
		Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) Find(id string) (Category, error) {
	var name, description string
	err := c.DB.QueryRow("SELECT name, description FROM categories WHERE id = ?", id).
		Scan(&name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
