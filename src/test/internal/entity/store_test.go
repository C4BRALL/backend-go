package test

import (
	"testing"

	entity "github.com/backend/src/internal/entity"
	uuid "github.com/backend/src/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	idSeller := uuid.NewID().String()
	store, err := entity.NewStore(idSeller, "testeName", "descrptionStore")
	assert.Nil(t, err)
	assert.NotNil(t, store)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.ID_seller)
	assert.NotEmpty(t, store.Name)
	assert.Equal(t, idSeller, store.ID_seller)
	assert.Equal(t, "testeName", store.Name)
	assert.Equal(t, "descrptionStore", store.Description)
}

func TestStoreWhenNameIsRequired(t *testing.T) {
	idSeller := uuid.NewID().String()
	s, err := entity.NewStore(idSeller, "", "descrptionStore")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrNameIsRequired, err)
}

func TestStoreWhenDescriptionIsRequired(t *testing.T) {
	idSeller := uuid.NewID().String()
	s, err := entity.NewStore(idSeller, "testeName", "")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrDescriptionIsRequired, err)
}

func TestStoreValidate(t *testing.T) {
	idSeller := uuid.NewID().String()
	s, err := entity.NewStore(idSeller, "testeName", "descrptionStore")
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Nil(t, s.Validate())
}
