package controller

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/helper"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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

	// create user
	user,err := c.repository.InsertUser(reqBody)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to create",err.Error())
	}

	// create user verification
	codeStr := helper.GenerateRandomString(20)
	err = c.repository.InsertUserVerification(user.Id,codeStr)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to create user verification code",err.Error())
	}

	// send code with email notification
	
	return dto.WriteResponseWithDetail(ctx, 201, "success register", user)
}

func (c *userController) RegisterVerification(ctx echo.Context) error{
	// get param path
	idStr := ctx.Param("id")
	id,err := strconv.ParseUint(idStr,36,32)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to parse string to uinteger",err.Error())
	}
	codeStr := ctx.Param("code")

	// update
	err = c.repository.UpdateUserStatusByIdAndCode(uint(id),codeStr)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to verification user account",err.Error())
	}

	return dto.WriteResponseWithDetail(ctx, 200, "success register verification", "detail")
}

func (c *userController) Login(ctx echo.Context) error{
	// bind
	var reqBody dto.ReqUserLogin
	err := ctx.Bind(&reqBody)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to bind",err.Error())
	}

	// find by email
	user,err := c.repository.FindUserByEmail(reqBody.Email)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"wrong email or password",err.Error())
	}

	// compare hash
	if !helper.CheckPasswordHash(reqBody.Password,user.Password){
		return dto.WriteResponseWithDetail(ctx,400,"wrong email or password",err.Error())
	}

	// check status
	if user.Status != "Verified" {
		return dto.WriteResponseWithDetail(ctx,400,"your account has not been verified. please check your email for the verification process",err.Error())
	}

	// generate jwt
	secretsign := []byte(os.Getenv("SECRETSIGN"))
	tokenString,err := helper.GenerateJWT(jwt.MapClaims{"id":user.Id},secretsign)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to generate jwt",err.Error())
	}


	return dto.WriteResponseWithDetail(ctx, 200, "success login", echo.Map{
		"jwt":tokenString,
	})
}