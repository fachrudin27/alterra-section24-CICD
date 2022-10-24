package controller

import (
	"net/http"
	"praktikum/dto"
	"praktikum/usecase"

	"github.com/labstack/echo"
)

type UserControllerInterface interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		useCase: userUsecase,
	}
}

func (u *userController) GetAll(c echo.Context) error {
	res, err := u.useCase.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"result":  res,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var payload dto.CreateUserRequest

	if err := c.Bind(&payload); err != nil {
		return err
	}

	res, err := u.useCase.CreateUser(payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"result":  res,
	})
}

func (u *userController) LoginUser(c echo.Context) error {
	// user := model.User{}
	var payloads dto.CreateUserRequest
	_ = c.Bind(&payloads)

	res, err := u.useCase.Login(payloads)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"result":  res,
	})

}
