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
	ID          entity.ID    `json:"id"`
	ID_seller   string       `json:"id_seller"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      enums.Status `json:"status"`
	CreatedAt   *time.Time   `json:"createdAt"`
	UpdatedAt   *time.Time   `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"-"`
}

func NewStore(id_seller string, name string, description string) (*Store, error) {
	store := &Store{
		ID:          entity.NewID(),
		ID_seller:   id_seller,
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
	if s.ID_seller == "" {
		return ErrIDSellerIsRequired
	}
	if _, err := entity.ParseID(s.ID_seller); err != nil {
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
