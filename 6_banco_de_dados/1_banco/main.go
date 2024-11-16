package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=database sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}

func main() {
	db, err := OpenConn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
