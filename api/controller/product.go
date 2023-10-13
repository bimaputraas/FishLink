package controller

import (
	"fishlink-mainapi/dto"
	"fishlink-mainapi/pb"

	"github.com/labstack/echo/v4"
)

func (c *productController) CreateProduct(ctx echo.Context) error {
	reqBody := dto.ReqBodyCreateProduct{}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "error to bind",
			"detail":  err.Error(),
		})
	}

	req := pb.CreateProductRequest{
		Product: &pb.Product{
			Name:        reqBody.Name,
			Description: reqBody.Description,
			Price:       reqBody.Price,
			Stock:       int32(reqBody.Stock),
		},
	}

	newProduct, err := c.Service.CreateProduct(ctx.Request().Context(), &req)
	if err != nil {
		return dto.ErrorResponse(ctx, err)
	}

	return ctx.JSON(201, echo.Map{
		"message": "success create",
		"detail":  newProduct,
	})
}

func (c *productController) GetAllProducts(ctx echo.Context) error {
    allProducts, err := c.Service.GetAllProduct(ctx.Request().Context(), &pb.GetAllProductRequest{})
    if err != nil {
        return dto.ErrorResponse(ctx, err)
    }

    return ctx.JSON(200, allProducts)
}

func (c *productController) GetProduct(ctx echo.Context) error {
	productID := ctx.Param("id")

	product, err := c.Service.GetProduct(ctx.Request().Context(), &pb.GetProductRequest{Id: productID})
	if err != nil {
		return dto.ErrorResponse(ctx, err)
	}

	return ctx.JSON(200, product)
}

func (c *productController) UpdateProduct(ctx echo.Context) error {
	productID := ctx.Param("id")

	reqBody := dto.ReqBodyUpdateProduct{}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "error to bind",
			"detail":  err.Error(),
		})
	}

	req := pb.UpdateProductRequest{
		Product: &pb.Product{
			Id:          productID,
			Name:        reqBody.Name,
			Description: reqBody.Description,
			Price:       reqBody.Price,
			Stock:       int32(reqBody.Stock),
		},
	}

	updatedProduct, err := c.Service.UpdateProduct(ctx.Request().Context(), &req)
	if err != nil {
		return dto.ErrorResponse(ctx, err)
	}

	return ctx.JSON(200, updatedProduct)
}

func (c *productController) DeleteProduct(ctx echo.Context) error {
	productID := ctx.Param("id")

	_, err := c.Service.DeleteProduct(ctx.Request().Context(), &pb.DeleteProductRequest{Id: productID})
	if err != nil {
		return dto.ErrorResponse(ctx, err)
	}

	return dto.WriteResponse(ctx,200,"delete success")
}