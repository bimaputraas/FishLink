package user

import (
	"final_project-ftgo-h8/api/dto"

	"github.com/labstack/echo/v4"
)

func (c *controller) Register(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 201, "Success register", "detail")
}

func (c *controller) RegisterVerification(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 200, "Success register verification", "detail")
}

func (c *controller) Login(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 200, "Success login", "detail")
}