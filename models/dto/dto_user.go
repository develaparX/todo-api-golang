package dto

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
