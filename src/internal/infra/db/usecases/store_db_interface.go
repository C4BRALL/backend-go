package usecases

import "github.com/backend/src/internal/entity"

type StoreDBnterface interface {
	Create(store *entity.Store) error
	FindAll(page, limit int, sort string) ([]entity.Store, error)
	FindById(id string) (*entity.Store, error)
	Update(store *entity.Store) error
	Delete(id string) error
}
