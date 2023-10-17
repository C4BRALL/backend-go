package main

import (
	"fmt"
	"net/http"

	"github.com/backend/src/configs"
	"github.com/backend/src/internal/entity"
	database "github.com/backend/src/internal/infra/db"
	"github.com/backend/src/internal/infra/webServer/handlers"
	"github.com/go-chi/chi"
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
	r.Post("/seller", sellerHandler.CreateSeller)
	r.Post("/seller/token", sellerHandler.GetJWT)
	r.Get("/seller/{email}", sellerHandler.GetSeller)
	r.Put("/seller/{email}", sellerHandler.UpdateSeller)
	r.Delete("/seller/{document}", sellerHandler.DeleteSeller)
	r.Get("/sellers", sellerHandler.GetSellers)

	r.Post("/store", StoreHandler.NewStore)
	http.ListenAndServe(config.WebServerPort, r)
}
