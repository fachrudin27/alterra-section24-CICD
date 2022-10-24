package router

import (
	"praktikum/controller"
	"praktikum/repository"
	"praktikum/usecase"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)

	userService := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userService)

	e.POST("/login", userController.LoginUser)
	e.POST("/", userController.CreateUser)
	e.GET("/", userController.GetAll)
}
