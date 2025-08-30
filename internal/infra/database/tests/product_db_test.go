package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/GuilhermeHRC/apis-fcycle/internal/entity"
	"github.com/GuilhermeHRC/apis-fcycle/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Test Product", 100)
	assert.NoError(t, err)

	productDB := database.NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	ProductDB := database.NewProduct(db)
	products, err := ProductDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = ProductDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test Product", 100)
	assert.NoError(t, err)

	productDB := database.NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	foundProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID.String(), foundProduct.ID.String())
	assert.Equal(t, product.Name, foundProduct.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test Product", 100)
	assert.NoError(t, err)

	productDB := database.NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Updated Product"
	err = productDB.Update(product)
	assert.NoError(t, err)

	updatedProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Updated Product", updatedProduct.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Test Product", 100)
	assert.NoError(t, err)

	productDB := database.NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
