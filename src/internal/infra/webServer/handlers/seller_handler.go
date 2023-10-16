package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/backend/src/internal/dto"
	"github.com/backend/src/internal/entity"
	"github.com/backend/src/internal/infra/db/usecases"
)

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
