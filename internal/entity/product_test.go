package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Test Product", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "Test Product", p.Name)
	assert.Equal(t, 100, p.Price)
	assert.NotEmpty(t, p.ID)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Test Product", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Test Product", -100)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
	assert.Nil(t, p)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Valid Product", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)

	err = p.Validate()
	assert.Nil(t, err)
}
