package dto

import "github.com/labstack/echo/v4"

type Response struct {
	Message string
    Code int
	Detail  interface{}
}

func WriteResponse(ctx echo.Context,code int, message string) error {
	return ctx.JSON(code,Response{
        Message: message,
        Code: code,
        Detail: "-",
    })
}

func WriteResponseWithDetail(ctx echo.Context,code int, message string, detail interface{}) error {
	return ctx.JSON(code,Response{
        Message: message,
        Code: code,
        Detail: detail,
    })
}