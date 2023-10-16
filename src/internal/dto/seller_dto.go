package dto

type CreateSellerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
