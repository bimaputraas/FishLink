package controller

import (
	"context"
	"encoding/json"
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"
	"final_project-ftgo-h8/helper"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (c *userController) Register(ctx echo.Context) error{
	// bind
	var reqBody dto.ReqBodyRegister
	err := ctx.Bind(&reqBody)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to bind",err.Error())
	}
	
	// validate 

	// validate and set role
	switch reqBody.Role{
	case "Admin":
		reqBody.Role = "Admin"
	case "User":
		reqBody.Role = "User"
	case "admin":
		reqBody.Role = "Admin"
	case "user":
		reqBody.Role = "User"
	case "":
		reqBody.Role = "User"
	default:
		return dto.WriteResponse(ctx,400,"invalid role")
	}

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

	// omitempty
	user.Id = 0
	user.Password = ""
	
	return dto.WriteResponseWithDetail(ctx, 201, "register success, please check your email for verification", user)
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
	_,err = c.repository.UpdateUserStatusByIdAndCode(uint(id),codeStr)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to verification user account",err.Error())
	}

	return dto.WriteResponse(ctx, 200, "your account has been successfully verified")
}

func (c *userController) Login(ctx echo.Context) error{
	// bind
	var reqBody dto.ReqBodyLogin
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
	tokenString,err := helper.GenerateJWT(jwt.MapClaims{"user_id":user.Id},secretsign)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,500,"failed to generate jwt",err.Error())
	}
	
	return dto.WriteResponseWithDetail(ctx, 200, "login success", echo.Map{
		"jwt":tokenString,
	})
}

func (c *userController) GetInfo(ctx echo.Context) error{
	user := ctx.Get("user").(model.User)
	// omitempty
	user.Password = ""
	user.Id = 0

	return dto.WriteResponseWithDetail(ctx, 200, "user account information", user)
}

func (c *userController) TopUp(ctx echo.Context) error{
	userId := ctx.Get("user").(model.User).Id

	var reqBody dto.TopUpReqBody
	err := ctx.Bind(&reqBody); if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"invalid request body",err.Error())
	}

	if reqBody.Amount < 1 {
		return dto.WriteResponse(ctx,400,"invalid amount")
	}

	user, err := c.repository.FindUserById(userId)
	if err != nil {
		return dto.WriteResponse(ctx,400,"user not found")
	}

	if _, err := c.repository.UpdateAmount(user, reqBody.Amount); err != nil {
		return dto.WriteResponse(ctx,400,"failed to top-up")
	}

	return dto.WriteResponse(ctx, 200, "top-up successfull")
}