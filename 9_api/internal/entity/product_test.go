package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product", 100)
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Product")
	assert.Equal(t, product.Price, 100)
	assert.NotEmpty(t, product.ID)
}

func TestWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.Nil(t, product)
	assert.Equal(t, ErrorNameIsRequired, err)
}

func TestWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("product", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrorPriceIsRequired, err)

}
