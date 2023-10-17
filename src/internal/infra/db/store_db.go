package db

import (
	"github.com/backend/src/internal/entity"
	"gorm.io/gorm"
)

type StoreDB struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *StoreDB {
	return &StoreDB{DB: db}
}

func (s *StoreDB) Create(store *entity.Store) error {
	return s.DB.Create(store).Error
}

func (s *StoreDB) FindById(id string) (*entity.Store, error) {
	var store entity.Store
	err := s.DB.First(&store, "id = ?", id).Error
	return &store, err
}

func (s *StoreDB) Update(store *entity.Store) error {
	_, err := s.FindById(string(store.ID.String()))
	if err != nil {
		return err
	}
	return s.DB.Save(store).Error
}

func (s *StoreDB) Delete(id string) error {
	store, err := s.FindById(id)
	if err != nil {
		return err
	}
	return s.DB.Delete(store).Error
}

func (s *StoreDB) FindAll(page, limit int, sort string) ([]entity.Store, error) {
	var stores []entity.Store
	var err error
	if sort != "" && sort != "desc" && sort != "asc" {
		sort = "desc"
	}
	if page != 0 && limit != 0 {
		err = s.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + sort).Find(&stores).Error
	} else {
		err = s.DB.Order("created_at" + sort).Find(&stores).Error
	}
	return stores, err
}
