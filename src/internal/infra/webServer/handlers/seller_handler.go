package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/backend/src/internal/dto"
	"github.com/backend/src/internal/entity"
	"github.com/backend/src/internal/infra/db/usecases"
	"github.com/go-chi/chi"
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
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	s, err := entity.NewSeller(seller.Name, seller.Email, seller.Document, seller.Password, seller.Phone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = h.SellerDB.Create(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Seller created!")
}

func (h *SellerHandler) GetSeller(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("email not provided")
	}

	seller, err := h.SellerDB.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(seller)
}

func (h *SellerHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("email not provided")
	}

	var seller entity.Seller
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	_, err = h.SellerDB.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("seller not found")
	}

	err = h.SellerDB.Update(&seller)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("seller updated")
}

func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	document := chi.URLParam(r, "document")
	if document == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("document not provided")
	}

	seller, err := h.SellerDB.FindByDocument(document)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("seller not found")
	}

	err = h.SellerDB.Delete(seller.ID.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("seller deleted")
}

func (h *SellerHandler) GetSellers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sellers, err := h.SellerDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal server error")
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellers)
}
