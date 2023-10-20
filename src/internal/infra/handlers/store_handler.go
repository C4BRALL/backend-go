package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/backend/src/internal/dto"
// 	"github.com/backend/src/internal/entity"
// 	"github.com/go-chi/chi"
// )

// type StoreHandler struct {
// 	StoreDB usecases.StoreDBnterface
// }

// func NewStoreHandler(storeDB usecases.StoreDBnterface) *StoreHandler {
// 	return &StoreHandler{StoreDB: storeDB}
// }

// func (h *StoreHandler) NewStore(w http.ResponseWriter, r *http.Request) {
// 	var store dto.CreateStoreInput
// 	err := json.NewDecoder(r.Body).Decode(&store)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	s, err := entity.NewStore(store.SellerID, store.Name, store.Description)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	err = h.StoreDB.Create(s)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("store created")
// }

// func (h *StoreHandler) GetStore(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("id not provided")
// 	}
// 	store, err := h.StoreDB.FindById(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}
// 	w.Header().Set("content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(store)
// }

// func (h *StoreHandler) UpdateStore(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("id not provided")
// 	}
// 	var store entity.Store
// 	err := json.NewDecoder(r.Body).Decode(&store)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// 	_, err = h.StoreDB.FindById(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("store not found")
// 	}
// 	err = h.StoreDB.Update(&store)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("store updated")
// }

// func (h *StoreHandler) DeleteStore(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("id not provided")
// 	}
// 	store, err := h.StoreDB.FindById(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("store not found")
// 	}
// 	err = h.StoreDB.Delete(store.ID.String())
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("store deleted")
// }

// func (h *StoreHandler) GetStores(w http.ResponseWriter, r *http.Request) {
// 	page := r.URL.Query().Get("page")
// 	limit := r.URL.Query().Get("limit")
// 	sort := r.URL.Query().Get("sort")
// 	pageInt, err := strconv.Atoi(page)
// 	if err != nil {
// 		pageInt = 0
// 	}
// 	limitInt, err := strconv.Atoi(limit)
// 	if err != nil {
// 		limitInt = 0
// 	}
// 	stores, err := h.StoreDB.FindAll(pageInt, limitInt, sort)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode("Internal server error")
// 	}
// 	w.Header().Set("content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(stores)
// }
