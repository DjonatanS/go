package database

import (
	"database/sql"
)

func InitDB(db *sql.DB) error {
	// Criar tabela de categorias
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT
		)
	`)
	if err != nil {
		return err
	}

	// Criar tabela de cursos
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS courses (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			category_id TEXT,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		)
	`)
	return err
}
