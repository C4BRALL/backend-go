package db

import (
	"github.com/backend/src/internal/entity"
	"gorm.io/gorm"
)

type SellerDB struct {
	DB *gorm.DB
}

func NewSeller(db *gorm.DB) *SellerDB {
	return &SellerDB{DB: db}
}

func (s *SellerDB) Create(seller *entity.Seller) error {
	return s.DB.Create(seller).Error
}

func (s *SellerDB) FindByEmail(email string) (*entity.Seller, error) {
	var seller entity.Seller
	if err := s.DB.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (s *SellerDB) Update(seller *entity.Seller) error {
	return s.DB.Where("email = ?", seller.Email).Updates(seller).Error
}

func (s *SellerDB) Delete(email string) error {
	seller, err := s.FindByEmail(email)
	if err != nil {
		return err
	}
	return s.DB.Delete(seller).Error
}

func (*SellerDB) FindAll(page int, limit int, sort string) ([]entity.Seller, error) {
	panic("unimplemented")
}
