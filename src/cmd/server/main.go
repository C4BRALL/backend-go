package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/backend/src/configs"
	"github.com/backend/src/internal/dto"
	"github.com/backend/src/internal/entity"
	database "github.com/backend/src/internal/infra/db"
	"github.com/backend/src/internal/infra/db/usecases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig("../../../")
	if err != nil {
		panic(err)
	}
	fmt.Println(config.DbDriver)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Seller{}, &entity.Store{})

	sellerDB := database.NewSeller(db)
	sellerHandler := NewSellerHandler(sellerDB)
	http.HandleFunc("/seller", sellerHandler.CreateSeller)

	http.ListenAndServe(config.WebServerPort, nil)
}

type SellerHandler struct {
	SellerDB usecases.SellerDBInterface
}

func NewSellerHandler(db usecases.SellerDBInterface) *SellerHandler {
	return &SellerHandler{SellerDB: db}
}

func (h *SellerHandler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	var seller dto.CreateSellerInput
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s, err := entity.NewSeller(seller.Name, seller.Email, seller.Document, seller.Password, seller.Phone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.SellerDB.Create(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
