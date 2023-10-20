package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/backend/src/internal/dto"
// 	"github.com/backend/src/internal/entity"
// 	"github.com/backend/src/internal/usecase"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/jwtauth"
// )

// type Error struct {
// 	Message string `json:"message"`
// }

// type SellerHandler struct {
// 	SellerDB usecases.SellerDBInterface
// }

// func NewSellerHandler(db usecases.SellerDBInterface) *SellerHandler {
// 	return &SellerHandler{
// 		SellerDB: db,
// 	}
// }

// // Get jwt token generate godoc
// // @Sumary Generate token jwt
// // @Description Generate token jwt
// // @Tags seller
// // @Accept json
// // @Produce json
// // @Param request body dto.GetJwtInput true "seller request"
// // @Success 200 {object} dto.GetJwtOutput
// // @Failure 500 {object} Error
// // @Router /seller/signin [post]
// func (h *SellerHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
// 	var seller dto.GetJwtInput
// 	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
// 	jwtExpiresIn := r.Context().Value("tokenExpiresIn").(int)
// 	err := json.NewDecoder(r.Body).Decode(&seller)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: "bad Request"}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	s, err := h.SellerDB.FindByEmail(seller.Email)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		error := Error{Message: "seller not found"}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	if !s.ValidatePassword(seller.Password) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		error := Error{Message: "credentials invalid"}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	_, tokenString, err := jwt.Encode(map[string]interface{}{
// 		"sub": s.ID.String(),
// 		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
// 	})
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadGateway)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	accessToken := dto.GetJwtOutput{AccessToken: tokenString}
// 	w.Header().Set("content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(accessToken)
// }

// // Create Seller godoc
// // @Sumary Create new seller
// // @Description Create a new seller
// // @Tags seller
// // @Accept json
// // @Produce json
// // @Param request body usecase.SellerInputDTO true "seller request"
// // @Success 201 "seller created"
// // @Failure 500 {object} Error
// // @Router /seller [post]
// func (h *SellerHandler) CreateSeller(w http.ResponseWriter, r *http.Request) {
// 	var seller usecase.SellerInputDTO
// 	err := json.NewDecoder(r.Body).Decode(&seller)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	s, err := entity.NewSeller(seller.Name, seller.Email, seller.Document, seller.Password, seller.Phone)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	err = h.SellerDB.Create(s)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode("Seller created!")
// }

// // Get Seller godoc
// // @Sumary Get a seller by email
// // @Description Get a seller by email
// // @Tags seller
// // @Accept json
// // @Produce json
// // @Param email path string true "E-mail of the seller"
// // @Success 200 {object} entity.Seller
// // @Failure 404 {object} Error
// // @Failure 500 {object} Error
// // @Router /seller/{email} [get]
// func (h *SellerHandler) GetSeller(w http.ResponseWriter, r *http.Request) {
// 	email := chi.URLParam(r, "email")
// 	if email == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: "email not provided"}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	seller, err := h.SellerDB.FindByEmail(email)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 		return
// 	}
// 	w.Header().Set("content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(seller)
// }

// func (h *SellerHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
// 	email := chi.URLParam(r, "email")
// 	if email == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: "email not provided"}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	var seller entity.Seller
// 	err := json.NewDecoder(r.Body).Decode(&seller)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	_, err = h.SellerDB.FindByEmail(email)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		error := Error{Message: "seller not found"}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	err = h.SellerDB.Update(&seller)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("seller updated")
// }

// func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
// 	document := chi.URLParam(r, "document")
// 	if document == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		error := Error{Message: "document not provided"}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	seller, err := h.SellerDB.FindByDocument(document)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		error := Error{Message: "seller not found"}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	err = h.SellerDB.Delete(seller.ID.String())
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("seller deleted")
// }

// // Get All Sellers godoc
// // @Sumary Get all sellers
// // @Description Get all sellers
// // @Tags seller
// // @Accept json
// // @Produce json
// // @Param page query string false "page_number"
// // @Param limit query string false "limit"
// // @Success 200 {array} entity.Seller
// // @Failure 404 {object} Error
// // @Failure 500 {object} Error
// // @Router /sellers/all [get]
// func (h *SellerHandler) GetSellers(w http.ResponseWriter, r *http.Request) {
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
// 	sellers, err := h.SellerDB.FindAll(pageInt, limitInt, sort)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		error := Error{Message: err.Error()}
// 		json.NewEncoder(w).Encode(error)
// 	}
// 	w.Header().Set("content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(sellers)
// }
