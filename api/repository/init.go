package repository

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"

	"gorm.io/gorm"
)

type UserRepository interface{
	CreateUser(dto.ReqUserRegister) (model.User,error)
}

type userRepository struct{
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository{
	return &userRepository{DB: db}
}
