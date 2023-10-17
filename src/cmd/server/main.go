package main

import (
	"fmt"
	"net/http"

	"github.com/backend/src/configs"
	"github.com/backend/src/internal/entity"
	database "github.com/backend/src/internal/infra/db"
	"github.com/backend/src/internal/infra/webServer/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig("../../../")
	if err != nil {
		panic(err)
	}
	fmt.Println("db name:", config.DbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Seller{}, &entity.Store{})

	sellerDB := database.NewSeller(db)
	sellerHandler := handlers.NewSellerHandler(sellerDB, config.TokenAuth, config.JWTExpiresIn)
	StoreDB := database.NewStore(db)
	StoreHandler := handlers.NewStoreHandler(StoreDB)

	r := chi.NewRouter()

	r.Route("/seller", func(r chi.Router) {
		r.Post("/", sellerHandler.CreateSeller)
		r.Post("/login", sellerHandler.GetJWT)
		r.Get("/{email}", sellerHandler.GetSeller)
		r.Put("/{email}", sellerHandler.UpdateSeller)
		r.Delete("/{document}", sellerHandler.DeleteSeller)
		r.Get("/all", sellerHandler.GetSellers)
	})

	r.Route("/store", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", StoreHandler.NewStore)
		r.Get("/", StoreHandler.GetStore)
		r.Put("/{id}", StoreHandler.UpdateStore)
		r.Delete("/{id}", StoreHandler.DeleteStore)
		r.Get("/all", StoreHandler.GetStores)
	})
	http.ListenAndServe(config.WebServerPort, r)
}
