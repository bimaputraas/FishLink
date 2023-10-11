package repository

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"
)

// user
type UserRepository interface {
	InsertUser(reqbody dto.ReqBodyRegister) (model.User, error)
	FindUserByEmail(email string) (model.User, error)
	InsertUserVerification(userId uint, code string) error
	UpdateUserStatusByIdAndCode(userId uint, code string) (model.User, error)
	FindUserById(userId uint) (model.User,error)
}