package database

import (
	"github.com/backend/src/internal/entity"
	"gorm.io/gorm"
)

type SellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{DB: db}
}

func (s *SellerRepository) Create(seller *entity.Seller) error {
	return s.DB.Create(seller).Error
}

func (s *SellerRepository) FindByDocument(document string) (*entity.Seller, error) {
	var seller entity.Seller
	if err := s.DB.Where("document = ?", document).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (s *SellerRepository) FindById(id string) (*entity.Seller, error) {
	var seller entity.Seller
	if err := s.DB.Where("id = ?", id).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (s *SellerRepository) FindByEmail(email string) (*entity.Seller, error) {
	var seller entity.Seller
	if err := s.DB.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (s *SellerRepository) Update(seller *entity.Seller) error {
	return s.DB.Where("email = ?", seller.Email).Updates(seller).Error
}

func (s *SellerRepository) Delete(id string) error {
	seller, err := s.FindById(id)
	if err != nil {
		return err
	}
	return s.DB.Delete(seller).Error
}

func (s *SellerRepository) FindAll(page int, limit int, sort string) ([]entity.Seller, error) {
	var sellers []entity.Seller
	var err error
	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "desc"
	}
	if page != 0 && limit != 0 {
		err = s.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + sort).Find(&sellers).Error
	} else {
		err = s.DB.Order("created_at" + sort).Find(&sellers).Error
	}
	return sellers, err
}
