package controller

import (
	"context"
	"encoding/json"
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

	// marshal model
	msgByte,err := json.Marshal(dto.UserEmailVerification{
		Email: user.Email,
		UserId: user.Id,
		VerificationCode: codeStr,
	})
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to marshal object",err.Error())
	}

	// send code with email notification
	err = c.publisher.PublishMessage(context.Background(),"fishlink-email_notification",msgByte)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to send email notification",err.Error())
	}

	
	return dto.WriteResponseWithDetail(ctx, 201, "success register", user)
}

func (c *userController) RegisterVerification(ctx echo.Context) error{
	// get param path
	idStr := ctx.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to parse string to uinteger",err.Error())
	}
	codeStr := ctx.Param("code")

	// update
	user,err := c.repository.UpdateUserStatusByIdAndCode(uint(id),codeStr)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to verification user account",err.Error())
	}

	return dto.WriteResponseWithDetail(ctx, 200, "success register verification", user)
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
		return dto.WriteResponse(ctx,400,"wrong email or password")
	}

	// compare hash
	if !helper.CheckPasswordHash(reqBody.Password,user.Password){
		return dto.WriteResponse(ctx,400,"wrong email or password")
	}
	
	// check status
	if user.Status != "Verified" {
		return dto.WriteResponse(ctx,401,"your account has not been verified. please check your email for the verification process")
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