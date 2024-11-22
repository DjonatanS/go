package entity

import (
	"api/pkg/entity"
	"errors"
	"time"
)

var (
	ErrorIDIsRequire     = errors.New("ID is required")
	ErrorInvalidID       = errors.New("ID is invalid")
	ErrorNameIsRequired  = errors.New("Name is required")
	ErrorPriceIsRequired = errors.New("Price is required")
	ErrorPriceIsInvalid  = errors.New("Price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt string    `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().Format(time.DateTime),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequire
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidID
	}
	if p.Name == "" {
		return ErrorNameIsRequired
	}
	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorPriceIsInvalid
	}
	return nil
}
