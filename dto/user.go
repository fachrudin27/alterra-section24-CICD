package dto

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Email string `json:"email"`
}

type UserJWT struct {
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
