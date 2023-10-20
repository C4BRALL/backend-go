package usecase

import (
	"fmt"

	"github.com/backend/src/internal/entity"
)

type SellerInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type SellerOutpuDTO struct {
	Message string `json:"message"`
}

type CreateSellerUseCase struct {
	SellerRepository entity.SellerRepositoryInterface
}

func NewCreateSellerUsecase(
	SellerRepository entity.SellerRepositoryInterface,
) *CreateSellerUseCase {
	return &CreateSellerUseCase{
		SellerRepository: SellerRepository,
	}
}

func (c *CreateSellerUseCase) Execute(input SellerInputDTO) (SellerOutpuDTO, error) {
	seller := entity.Seller{
		Name:     input.Name,
		Email:    input.Email,
		Document: input.Document,
		Password: input.Password,
		Phone:    input.Phone,
	}

	if err := c.SellerRepository.Create(&seller); err != nil {
		return SellerOutpuDTO{Message: err.Error()}, err
	}

	dto := SellerOutpuDTO{
		Message: fmt.Sprintf("seller %s created successfully", seller.Name),
	}

	return dto, nil
}
