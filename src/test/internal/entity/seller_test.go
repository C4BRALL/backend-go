package test

import (
	"testing"

	entity "github.com/backend/src/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewSeller(t *testing.T) {
	seller, err := entity.NewSeller("teste", "teste@mail.com", "12312312332", "passphrase", "85912341234")
	assert.Nil(t, err)
	assert.NotNil(t, seller)
	assert.NotEmpty(t, seller.ID)
	assert.NotEmpty(t, seller.Password)
	assert.Equal(t, "teste", seller.Name)
	assert.Equal(t, "teste@mail.com", seller.Email)
}

func TestSeller_ValidatePassword(t *testing.T) {
	seller, err := entity.NewSeller("teste", "teste@mail.com", "12312312332", "passphrase", "85912341234")
	assert.Nil(t, err)
	assert.True(t, seller.ValidatePassword("passphrase"))
	assert.False(t, seller.ValidatePassword("passphrase3"))
	assert.NotEqual(t, "passphrase", seller.Password)
}

func TestSellerWhenNameIsRequired(t *testing.T) {
	s, err := entity.NewSeller("", "teste@mail.com", "12312312332", "passphrase", "85912341234")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrNameIsRequired, err)
}

func TestSellerWhenEmailIsRequired(t *testing.T) {
	s, err := entity.NewSeller("teste", "", "12312312332", "passphrase", "85912341234")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrEmailIsRequired, err)
}

func TestSellerWhenDocumentIsRequired(t *testing.T) {
	s, err := entity.NewSeller("teste", "teste@mail.com", "", "passphrase", "85912341234")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrDocumentIsRequired, err)
}

func TestSellerWhenPasswordIsRequired(t *testing.T) {
	s, err := entity.NewSeller("teste", "teste@mail.com", "12312312332", "", "85912341234")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrPasswordIsRequired, err)
}

func TestSellerWhenPhoneIsRequired(t *testing.T) {
	s, err := entity.NewSeller("teste", "teste@mail.com", "12312312332", "passphrase", "")
	assert.Nil(t, s)
	assert.Equal(t, entity.ErrPhoneIsRequired, err)
}

func TestSellerValidate(t *testing.T) {
	s, err := entity.NewSeller("teste", "teste@mail.com", "12312312332", "passphrase", "85912341234")
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Nil(t, s.Validate())
}
