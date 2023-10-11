package entity

import (
	"time"

	enums "github.com/backend/src/internal/enums"

	"github.com/backend/src/pkg/entity"
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
	return &Store{
		ID:          entity.NewID(),
		ID_seller:   id_seller,
		Name:        name,
		Description: description,
		Status:      enums.Status(enums.Active),
	}, nil
}
