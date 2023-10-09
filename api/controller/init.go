package controller

import (
	"final_project-ftgo-h8/api/repository"

	"github.com/labstack/echo/v4"
)

// user controller
type UserController interface{
	Register(echo.Context) error
	Login(echo.Context) error
	RegisterVerification(echo.Context) error
}

type userController struct{
	repository repository.UserRepository
}

func NewController(r repository.UserRepository) UserController{
	return &userController{repository: r}
}

// product



