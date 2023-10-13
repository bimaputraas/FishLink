package repository

import (
	"errors"
	"fishlink-mainapi/dto"
	"fishlink-mainapi/model"
	"time"
)


func (r *orderRepository) InsertOrderAndDetail(reqBody dto.ReqBodyNewOrder, userId uint) (model.OrderDetail,error) {
	// init user
	var user model.User
	result := r.gormDb.First(&user,userId)
	if result.Error != nil{
		return model.OrderDetail{},result.Error
	}

	// init product
	var product model.Product
	result = r.gormDb.First(&product,reqBody.ProductId)
	if result.Error != nil{
		return model.OrderDetail{},result.Error
	}
	
	// init total price
	totalPrice := product.Price * int64(reqBody.Quantity)

	// init and start gorm transaction
	tx := r.gormDb.Begin()

	// check & update product stock
	product.Stock -= reqBody.Quantity
	if product.Stock < 0 {
		return model.OrderDetail{},errors.New("product stock is unavailable")
	}

	result = tx.Save(&product)
	if result.Error != nil {
		tx.Rollback()
		return model.OrderDetail{},result.Error
	}
	
	// check & update user amount
	user.Amount -= totalPrice
	if user.Amount < 0{
		tx.Rollback()
		return model.OrderDetail{},errors.New("insufficient funds")
	}
	
	result = tx.Save(&user)
	if result.Error != nil {
		tx.Rollback()
		return model.OrderDetail{},result.Error
	}

	// create order
	order := model.Order{
		UserId: user.Id,
		OrderDate: time.Now(),
	}

	result = tx.Create(&order)
	if result.Error != nil {
		tx.Rollback()
		return model.OrderDetail{},result.Error
	}

	// create order detail
	orderDetail := model.OrderDetail{
		OrderId: order.Id,
		Order: order,
		ProductId: product.ID,
		Product: product,
		TotalPrice: totalPrice,
		Quantity: reqBody.Quantity,
	}

	result = tx.Create(&orderDetail)
	if result.Error != nil {
		tx.Rollback()
		return model.OrderDetail{},result.Error
	}

	tx.Commit()
	
	return orderDetail,nil
}

func (r *orderRepository) FindOrderDetails(userId uint) ([]model.OrderDetail,error) {
	// model
	orderDetails := []model.OrderDetail{}
	orderDetailsResult := []model.OrderDetail{}

	// find
	result := r.gormDb.Preload("Product").Preload("Order").Find(&orderDetails)
	if result.Error != nil {
		return nil,result.Error
	}

	// result
	for _,v := range orderDetails{
		if v.Order.UserId != userId{
			continue
		}
		orderDetailsResult = append(orderDetailsResult, v)
	}
	
	return orderDetailsResult,nil
}

