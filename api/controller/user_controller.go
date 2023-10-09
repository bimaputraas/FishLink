package controller

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/helper"

	"github.com/labstack/echo/v4"
)

func (c *userController) Register(ctx echo.Context) error{
	// bind
	var reqBody dto.ReqUserRegister
	err := ctx.Bind(&reqBody)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to bind",err.Error())
	}

	// validate

	// hash
	hash,err := helper.HashPassword(reqBody.Password)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to hash",err.Error())
	}
	reqBody.Password = hash

	// create
	user,err := c.repository.CreateUser(reqBody)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to create",err.Error())
	}
	
	return dto.WriteResponseWithDetail(ctx, 201, "Success register", user)
}

func (c *userController) RegisterVerification(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 200, "Success register verification", "detail")
}

func (c *userController) Login(ctx echo.Context) error{
	return dto.WriteResponseWithDetail(ctx, 200, "Success login", "detail")
}