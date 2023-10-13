package middleware

import (
	"fishlink-mainapi/repository"

	"github.com/labstack/echo/v4"
)

type AuthenticationMiddleware interface {
	Authentication(next echo.HandlerFunc) echo.HandlerFunc
	AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc
}

type authenticationMiddleware struct {
	repository repository.UserRepository
}

func NewAuthenticationMiddleware(r repository.UserRepository) AuthenticationMiddleware {
	return &authenticationMiddleware{repository: r}
}