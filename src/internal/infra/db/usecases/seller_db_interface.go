package usecases

import "github.com/backend/src/internal/entity"

type SellerDBInterface interface {
	CreateSeller(seller *entity.Seller) error
	FindByEmail(email string) (*entity.Seller, error)
}
