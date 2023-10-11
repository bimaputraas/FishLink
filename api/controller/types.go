package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Register(echo.Context) error
	Login(echo.Context) error
	RegisterVerification(echo.Context) error
	TopUp(ctx echo.Context) error
	GetInfo(ctx echo.Context) error
}
type ProductController interface {
	CreateProduct(echo.Context) error
	GetAllProducts(echo.Context) error
	GetProduct(echo.Context) error
	UpdateProduct(echo.Context) error
	DeleteProduct(ctx echo.Context) error
}

type OrderController interface{
	NewOrder(ctx echo.Context) error
	GetOrders(ctx echo.Context) error
}