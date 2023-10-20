package dto

type GetJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtOutput struct {
	AccessToken string `json:"access_token"`
}
