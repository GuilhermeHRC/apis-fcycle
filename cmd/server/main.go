package main

import (
	"net/http"

	"github.com/GuilhermeHRC/apis-fcycle/configs"
	"github.com/GuilhermeHRC/apis-fcycle/internal/entity"
	"github.com/GuilhermeHRC/apis-fcycle/internal/infra/database"
	"github.com/GuilhermeHRC/apis-fcycle/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.GetProducts)

	http.ListenAndServe(":8000", r)
}
