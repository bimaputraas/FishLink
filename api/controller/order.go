package controller

import (
	"fishlink-mainapi/dto"
	"fishlink-mainapi/model"
	"time"

	"github.com/labstack/echo/v4"
)

func (c *orderController) NewOrder(ctx echo.Context) error{
	// get user id
	userId := ctx.Get("user").(model.User).Id

	// bind
	var reqBody dto.ReqBodyNewOrder
	err := ctx.Bind(&reqBody)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to bind",err.Error())
	}
	
	// validate

	// validate quantity
	if reqBody.Quantity < 1 {
		return dto.WriteResponse(ctx,400,"invalid to bind field quantity")
	}


	// create order
	orderDetail,err := c.repository.InsertOrderAndDetail(reqBody,userId)
	if err != nil {
		return dto.WriteResponseWithDetail(ctx,400,"failed to order",err.Error())
	}

	// detail
	resDetail := dto.ResDetailNewOrder{
		ProductName: orderDetail.Product.Name,
		Quantity: orderDetail.Quantity,
		Ordered_at: orderDetail.Order.OrderDate.Format(time.DateTime),
		TotalPrice: orderDetail.TotalPrice,
	}

	
	return dto.WriteResponseWithDetail(ctx, 201, "success order", resDetail)
}

func (c *orderController) GetOrders(ctx echo.Context) error{
	// get user id
	userId := ctx.Get("user").(model.User).Id

	// find
	orderDetails,err := c.repository.FindOrderDetails(userId)
	if err != nil {
		dto.WriteResponseWithDetail(ctx,400,"failed to get order details",orderDetails)
	}

	// detail
	var resDetails []dto.ResDetailGetOrder
	for _,v := range orderDetails{
		resDetail := dto.ResDetailGetOrder{
			ProductName: v.Product.Name,
			ProductPrice: v.Product.Price,
			ProductDescription: v.Product.Description,
			Quantity: v.Quantity,
			TotalPrice: v.TotalPrice,
			Ordered_at: v.Order.OrderDate.Format(time.DateTime),
		}
		resDetails = append(resDetails, resDetail)
	}

	return dto.WriteResponseWithDetail(ctx, 200, "order detail informations", resDetails)
}