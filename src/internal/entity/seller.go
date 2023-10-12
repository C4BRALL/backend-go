package entity

import (
	"errors"
	"time"

	enums "github.com/backend/src/internal/enums"

	"github.com/backend/src/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrIDIsRequired       = errors.New("id required")
	ErrInvalidId          = errors.New("invalid id")
	ErrNameIsRequired     = errors.New("name required")
	ErrEmailIsRequired    = errors.New("email required")
	ErrDocumentIsRequired = errors.New("document required")
	ErrPasswordIsRequired = errors.New("password required")
	ErrPhoneIsRequired    = errors.New("phone required")
)

type Seller struct {
	ID        entity.ID      `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Document  string         `json:"document"`
	Password  string         `json:"-"`
	Phone     string         `json:"phone"`
	Type      enums.UserType `json:"type"`
	Status    enums.Status   `json:"status"`
	Stores    []Store        `json:"stores"` // Não é necessário gorm:"foreignKey:SellerID"
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt *time.Time     `json:"-"`
}

func NewSeller(name, email, document, password string, phone string) (*Seller, error) {
	if password == "" {
		return nil, ErrPasswordIsRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return nil, err
	}

	seller := &Seller{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Document: document,
		Password: string(hash),
		Phone:    phone,
		Type:     enums.UserType(enums.Salesperson),
		Status:   enums.Status(enums.Active),
	}

	validSeller := seller.Validate()
	if validSeller != nil {
		return nil, validSeller
	}
	return seller, nil
}

func (s *Seller) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password))
	return err == nil
}

func (s *Seller) Validate() error {
	if s.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(s.ID.String()); err != nil {
		return ErrInvalidId
	}
	if s.Name == "" {
		return ErrNameIsRequired
	}
	if s.Email == "" {
		return ErrEmailIsRequired
	}
	if s.Document == "" {
		return ErrDocumentIsRequired
	}

	if s.Phone == "" {
		return ErrPhoneIsRequired
	}
	return nil
}
