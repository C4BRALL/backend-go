package test

import (
	"testing"

	entity "github.com/backend/src/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	store, err := entity.NewStore("sellerId", "testeName", "descrptionStore")
	assert.Nil(t, err)
	assert.NotNil(t, store)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.ID_seller)
	assert.NotEmpty(t, store.Name)
	assert.Equal(t, "sellerId", store.ID_seller)
	assert.Equal(t, "testeName", store.Name)
	assert.Equal(t, "descrptionStore", store.Description)
}
