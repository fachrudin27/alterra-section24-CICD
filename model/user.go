package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersLoginResponse struct {
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
