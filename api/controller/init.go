package controller

import (
	"final_project-ftgo-h8/api/repository"
	"final_project-ftgo-h8/pb"

	"github.com/labstack/echo/v4"
)

// user controller
type UserController interface {
	Register(echo.Context) error
	Login(echo.Context) error
	RegisterVerification(echo.Context) error
}

type userController struct {
	repository repository.UserRepository
}

func NewController(r repository.UserRepository) UserController {
	return &userController{repository: r}
}

// product controller
type ProductController interface {
	CreateProduct(echo.Context) error
	GetAllProducts(echo.Context) error
	GetProduct(echo.Context) error
	UpdateProduct(echo.Context) error
	DeleteProduct(ctx echo.Context) error
}

type productController struct {
	Service pb.ProductServiceClient
}

func NewProductController(pb pb.ProductServiceClient) ProductController {
	return &productController{Service: pb}
}