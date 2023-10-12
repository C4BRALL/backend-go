package entity

import (
	"errors"
	"time"

	enums "github.com/backend/src/internal/enums"

	"github.com/backend/src/pkg/entity"
)

var (
	ErrDescriptionIsRequired = errors.New("description required")
	ErrIDSellerIsRequired    = errors.New("id_seller required")
)

type Store struct {
	ID          entity.ID    `json:"id" gorm:"primaryKey"`
	SellerID    entity.ID    `json:"sellerID" gorm:"index"` // Adicione uma chave estrangeira aqui
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      enums.Status `json:"status"`
	CreatedAt   *time.Time   `json:"createdAt"`
	UpdatedAt   *time.Time   `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"-"`
}

func NewStore(id_seller entity.ID, name string, description string) (*Store, error) {
	store := &Store{
		ID:          entity.NewID(),
		SellerID:    id_seller,
		Name:        name,
		Description: description,
		Status:      enums.Status(enums.Active),
	}

	storeValid := store.Validate()
	if storeValid != nil {
		return nil, storeValid
	}
	return store, nil
}

func (s *Store) Validate() error {
	if s.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(s.ID.String()); err != nil {
		return ErrInvalidId
	}
	if _, err := entity.ParseID(string(s.SellerID.String())); err != nil {
		return ErrInvalidId
	}
	if s.Name == "" {
		return ErrNameIsRequired
	}
	if s.Description == "" {
		return ErrDescriptionIsRequired
	}
	return nil
}
