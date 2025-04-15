package product

import (
	"database/sql"
)

// Repository defines the interface for product repository operations.
type Repository interface {
	GetProductByID(id int) (*Product, error)
}

type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

// GetProductByID retrieves a product by its ID.
func (r *ProductRepository) GetProductByID(id int) (*Product, error) {
	return &Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}
