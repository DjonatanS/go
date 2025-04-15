package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database connection
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usecase := NewUseCase(db)

	// Use the product use case to get a product by ID
	product, err := usecase.GetProductByID(1)
	if err != nil {
		panic(err)
	}

	// Print the product details
	println("Product ID:", product.ID)
	println("Product Name:", product.Name)
}
