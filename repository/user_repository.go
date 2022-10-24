package repository

import (
	"praktikum/dto"
	"praktikum/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	CreateUser(data model.User) (dto.CreateUserRequest, error)
	Login(payloads dto.CreateUserRequest) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(data model.User) (dto.CreateUserRequest, error) {
	c := dto.CreateUserRequest{
		Email:    data.Email,
		Password: data.Password,
	}
	if err := u.db.Create(&model.User{
		Email:    c.Email,
		Password: c.Password,
	}).Error; err != nil {
		return c, err
	}

	return c, nil
}

func (u *userRepository) GetAll() ([]model.User, error) {
	users := []model.User{}

	if err := u.db.Model(&model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Login(data dto.CreateUserRequest) (model.User, error) {
	var user model.User

	if err := u.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
