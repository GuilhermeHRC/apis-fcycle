package database

import "github.com/GuilhermeHRC/apis-fcycle/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

// type ProductInterface interface {
// 	Create(product *entity.Product) error
// 	FindByID(id entity.ID) (*entity.Product, error)
// 	FindAll() ([]*entity.Product, error)
// }
