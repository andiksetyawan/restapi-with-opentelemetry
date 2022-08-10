package model

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginSignupResponse struct {
	Token string `json:"token"`
}
