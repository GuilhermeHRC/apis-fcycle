package entity

import (
	"errors"
	"time"

	"github.com/GuilhermeHRC/apis-fcycle/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrIDsRequired     = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id format")
	ErrNameIsRequired  = errors.New("name is required")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrPriceIsRequired = errors.New("price is required")
)

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	return nil
}
