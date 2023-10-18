package dto

import "github.com/backend/src/pkg/entity"

type CreateStoreInput struct {
	SellerID    entity.ID `json:"sellerID"` // Adicione uma chave estrangeira aqui
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
