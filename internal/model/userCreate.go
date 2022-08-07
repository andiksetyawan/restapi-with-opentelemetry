package model

type UserCreateRequest struct {
	Name string
}

type UserCreateResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}
