package dto

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}
