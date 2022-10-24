package usecase

import (
	"praktikum/dto"
	m "praktikum/middleware"
	"praktikum/model"
	"praktikum/repository"
)

type UserUsecase interface {
	GetAll() ([]dto.UserResponse, error)
	CreateUser(payloads dto.CreateUserRequest) (dto.CreateUserRequest, error)
	Login(payloads dto.CreateUserRequest) (dto.UserJWT, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) CreateUser(payloads dto.CreateUserRequest) (dto.CreateUserRequest, error) {

	userData := model.User{
		Email:    payloads.Email,
		Password: payloads.Password,
	}

	res, err := s.userRepository.CreateUser(userData)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *userUsecase) GetAll() ([]dto.UserResponse, error) {
	res, err := s.userRepository.GetAll()

	var a []dto.UserResponse

	for _, v := range res {

		a = append(a, dto.UserResponse{
			Email: v.Email,
		})

	}

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *userUsecase) Login(payloads dto.CreateUserRequest) (dto.UserJWT, error) {

	var res dto.UserJWT

	user, err := s.userRepository.Login(payloads)

	token, errt := m.CreateToken(uint(user.ID), user.Email)

	res = dto.UserJWT{
		Email: user.Email,
		Token: token,
	}

	if err != nil {
		return res, err
	}

	if errt != nil {
		return res, errt
	}

	return res, nil
}
