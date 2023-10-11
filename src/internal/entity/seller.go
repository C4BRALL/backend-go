package entity

import (
	"time"

	enums "github.com/backend/src/internal/enums"

	"github.com/backend/src/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	ID        entity.ID      `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Document  string         `json:"document"`
	Password  string         `json:"-"`
	Phone     string         `json:"phone"`
	Type      enums.UserType `json:"type"`
	Status    enums.Status   `json:"status"`
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt *time.Time     `json:"-"`
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
		Type:     enums.UserType(enums.Salesperson),
		Status:   enums.Status(enums.Active),
	}, nil
}

func (s *Seller) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password))
	return err == nil
}
