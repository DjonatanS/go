package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=database sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

func main() {
	db, err := OpenConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	product := NewProduct("Notebook", 2000.0)
	err = InsertProduct(db, product)
	if err != nil {
		fmt.Println(err)
		return
	}
	product.Price = 200
	err = UpdateProduct(db, product)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		p, err := SelectProduct(db, product.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Product: %+v\n price: %f\n", p.Name, p.Price)
	*/

	products, err := SelectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Name: %v, Price: %.2f \n", p.Name, p.Price)
	}
}

func InsertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = $1, price = $2 WHERE id = $3")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
	return nil
}

func CreateProductTable(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE products (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price NUMERIC(10, 2) NOT NULL
	)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

func SelectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &product, nil
}

func SelectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name , price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func DeleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}
