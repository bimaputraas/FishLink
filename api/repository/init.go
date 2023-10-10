package repository

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"

	"gorm.io/gorm"
)

type UserRepository interface{
	InsertUser(reqbody dto.ReqUserRegister) (model.User,error)
	FindUserByEmail(email string) (model.User,error)
	InsertUserVerification(userId uint, code string) error
	UpdateUserStatusByIdAndCode(userId uint, code string) error
}

type userRepository struct{
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository{
	return &userRepository{DB: db}
}
