package mock

import (
	"praktikum/dto"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func (u *MockUser) GetAll() ([]dto.UserResponse, error) {
	args := u.Called()

	return args.Get(0).([]dto.UserResponse), args.Error(1)
}

func (u *MockUser) CreateUser(payloads dto.CreateUserRequest) (dto.CreateUserRequest, error) {
	args := u.Called(payloads)

	return args.Get(0).(dto.CreateUserRequest), args.Error(1)
}

func (u *MockUser) Login(payloads dto.CreateUserRequest) (dto.UserJWT, error) {
	args := u.Called(payloads)

	return args.Get(0).(dto.UserJWT), args.Error(1)

}
