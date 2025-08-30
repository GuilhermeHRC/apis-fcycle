package main

import (
	"fmt"
	"net/http"

	"github.com/GuilhermeHRC/apis-fcycle/configs"
	_ "github.com/GuilhermeHRC/apis-fcycle/docs"
	"github.com/GuilhermeHRC/apis-fcycle/internal/entity"
	"github.com/GuilhermeHRC/apis-fcycle/internal/infra/database"
	"github.com/GuilhermeHRC/apis-fcycle/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Guilherme Coelho
// @contact.email  coelho.ghr@gmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Postgres connection
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			configs.DBHost, configs.DBUser, configs.DBPassword, configs.DBName, configs.DBPort)), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/generate-token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
