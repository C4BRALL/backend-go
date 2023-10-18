package dto

type CreateSellerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type GetJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtOutput struct {
	AccessToken string `json:"access_token"`
}
