package test

import (
	"testing"

	entity "github.com/backend/src/internal/entity"
	uuid "github.com/backend/src/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	idSeller := uuid.NewID()
	store, err := entity.NewStore(idSeller, "testeName", "descrptionStore")
	assert.Nil(t, err)
	assert.NotNil(t, store)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.ID_seller)
	assert.NotEmpty(t, store.Name)
	assert.Equal(t, idSeller.String(), store.ID_seller.ID.String())
	assert.Equal(t, "testeName", store.Name)
	assert.Equal(t, "descrptionStore", store.Description)
}
