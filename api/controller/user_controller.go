package controller

import (
	"final_project-ftgo-h8/api/dto"

	"github.com/labstack/echo/v4"
)

func (c *userController) Register(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 201, "Success register", "detail")
}

func (c *userController) RegisterVerification(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 200, "Success register verification", "detail")
}

func (c *userController) Login(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 200, "Success login", "detail")
}