package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/backend/src/internal/dto"
	"github.com/backend/src/internal/entity"
	"github.com/backend/src/internal/infra/db/usecases"
)

type StoreHandler struct {
	StoreDB usecases.StoreDBnterface
}

func NewStoreHandler(storeDB usecases.StoreDBnterface) *StoreHandler {
	return &StoreHandler{StoreDB: storeDB}
}

func (h *StoreHandler) NewStore(w http.ResponseWriter, r *http.Request) {
	var store dto.CreateStoreInput
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	s, err := entity.NewStore(store.SellerID, store.Name, store.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = h.StoreDB.Create(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("store created")
}
