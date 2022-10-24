package controller

import (
	// "github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	// "github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/dto"

	"github.com/stretchr/testify/suite"

	// "problem1/model"
	mocks "praktikum/usecase/mock"
	"testing"
)

type suiteUsers struct {
	suite.Suite
	controller *userController
	mockServ   *mocks.MockUser
}

func (s *suiteUsers) SetupSuite() {
	mocks := &mocks.MockUser{}
	s.mockServ = mocks

	s.controller = &userController{
		useCase: mocks,
	}
}
func (s *suiteUsers) TestGetAll() {
	s.mockServ.On("GetAll").Return([]dto.UserResponse{
		{
			Email: "fachrudin@alterra.com",
		},
	}, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpextedBody       dto.UserResponse
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			dto.UserResponse{
				Email: "fachrudin@alterra.com",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.controller.GetAll(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)
		})
	}
}

func (s *suiteUsers) TestCreateUser() {
	data := dto.CreateUserRequest{
		Email:    "fachru@alterra.com",
		Password: "12345",
	}
	s.mockServ.On("CreateUser", data).Return(data, nil)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               dto.CreateUserRequest
		HasReturnBody      bool
		ExpextedBody       dto.CreateUserRequest
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			dto.CreateUserRequest{
				Email:    "fachru@alterra.com",
				Password: "12345",
			},
			true,
			dto.CreateUserRequest{
				Email:    "fachru@alterra.com",
				Password: "12345",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			err := s.controller.CreateUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]dto.CreateUserRequest
				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				s.Equal(v.ExpextedBody.Email, resp["result"].Email)
			}
		})
	}
}

func (s *suiteUsers) TestLoginUser() {
	data := dto.CreateUserRequest{
		Email:    "fachru@alterra.com",
		Password: "12345",
	}
	s.mockServ.On("Login", data).Return(dto.UserJWT{
		Email: "fachru@alterra.com",
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJvbmFsZG9AYWx0ZXJyYS5jb20iLCJleHAiOjE2NjYwODQxMDYsInVzZXJJZCI6NH0.8o2TLUIQhs-Ics06IcqQU0N1Imo2ZN-rjwOyk8fFGtM",
	}, nil)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               dto.CreateUserRequest
		HasReturnBody      bool
		ExpextedBody       dto.CreateUserRequest
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			dto.CreateUserRequest{
				Email:    "fachru@alterra.com",
				Password: "12345",
			},
			true,
			dto.CreateUserRequest{
				Email: "fachru@alterra.com",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/login", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			err := s.controller.LoginUser(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]dto.UserJWT
				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				s.Equal(v.ExpextedBody.Email, resp["result"].Email)
			}
		})
	}
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
