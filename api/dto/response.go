package dto

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Response struct {
	Message string
	Code    int
	Detail  interface{}
}

func WriteResponse(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, Response{
		Message: message,
		Code:    code,
		Detail:  "-",
	})
}

func WriteResponseWithDetail(ctx echo.Context, code int, message string, detail interface{}) error {
	return ctx.JSON(code, Response{
		Message: message,
		Code:    code,
		Detail:  detail,
	})
}

func ErrorResponse(ctx echo.Context, err error) error {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.NotFound:
			return echo.NewHTTPError(404, echo.Map{"message": "not found", "detail": st.Message()})
		case codes.InvalidArgument:
			return echo.NewHTTPError(400, echo.Map{"message": "invalid argument", "detail": st.Message()})
		default:
			return echo.NewHTTPError(500, echo.Map{"message": "internal server error", "detail": st.Message()})
		}
	}
	return err
}