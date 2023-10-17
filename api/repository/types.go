package repository

import (
	"fishlink-mainapi/dto"
	"fishlink-mainapi/model"
)

// user
type UserRepository interface {
	InsertUser(reqbody dto.ReqBodyRegister) (model.User, error)
	FindUserByEmail(email string) (model.User, error)
	InsertUserVerification(userId uint, code string) error
	UpdateUserStatusByIdAndCode(userId uint, code string) (model.User, error)
	FindUserById(userId uint) (model.User, error)
	UpdateAmount(user model.User, amount int64) (model.User, error)
}

type OrderRepository interface{
	InsertOrderAndDetail(reqBody dto.ReqBodyNewOrder, userId uint) (model.OrderDetail,error)
	FindOrderDetails(userId uint) ([]model.OrderDetail,error)
}