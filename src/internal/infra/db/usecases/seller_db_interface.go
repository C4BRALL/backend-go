package usecases

import "github.com/backend/src/internal/entity"

type SellerDBInterface interface {
	Create(seller *entity.Seller) error
	FindByEmail(email string) (*entity.Seller, error)
	FindAll(page, limit int, sort string) ([]entity.Seller, error)
	Delete(id string) error
	Update(seller *entity.Seller) error
}
