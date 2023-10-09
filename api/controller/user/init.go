package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller interface{
	Register(echo.Context) error
	Login(echo.Context) error
	RegisterVerification(echo.Context) error
}

type controller struct{
	DB *gorm.DB
}

func NewController(db *gorm.DB) Controller{
	return &controller{DB: db}
}

