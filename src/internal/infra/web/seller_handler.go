package web

import (
	"encoding/json"
	"net/http"

	"github.com/backend/src/internal/entity"
	"github.com/backend/src/internal/usecase"
)

type WebSellerHandler struct {
	SellerRepository entity.SellerRepositoryInterface
}

func NewWebSellerHandler(
	SellerRepository entity.SellerRepositoryInterface,
) *WebSellerHandler {
	return &WebSellerHandler{
		SellerRepository: SellerRepository,
	}
}

// Create Seller godoc
// @Sumary Create new seller
// @Description Create a new seller
// @Tags seller
// @Accept json
// @Produce json
// @Param request body usecase.SellerInputDTO true "seller request"
// @Success 201 "seller created"
// @Failure 500 {object} Error
// @Router /seller [post]
func (h *WebSellerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var seller usecase.SellerInputDTO
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createSeller := usecase.NewCreateSellerUsecase(h.SellerRepository)
	output, err := createSeller.Execute(seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
