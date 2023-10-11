package entity

import (
	"time"

	"github.com/backend/src/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Status int

const (
	Active   Status = 1
	Inactive Status = 0
)

type SellerType string

const (
	Customer    SellerType = "customers"
	Salesperson SellerType = "seller"
)

type Seller struct {
	ID        entity.ID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Document  string     `json:"document"`
	Password  string     `json:"-"`
	Phone     string     `json:"phone"`
	Type      SellerType `json:"type"`
	Status    Status     `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

func NewSeller(name, email, document, password string, phone string) (*Seller, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Seller{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Document: document,
		Password: string(hash),
		Phone:    phone,
		Type:     SellerType(Customer),
		Status:   Status(Active),
	}, nil
}

func (s *Seller) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password))
	return err == nil
}
