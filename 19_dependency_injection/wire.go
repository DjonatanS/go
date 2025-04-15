//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"dependency/product"

	"github.com/google/wire"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
