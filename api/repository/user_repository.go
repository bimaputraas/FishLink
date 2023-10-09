package repository

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"
	"time"
)

func (r *userRepository) CreateUser(reqBody dto.ReqUserRegister) (model.User,error){
	// model
	user := model.User{
		Name: reqBody.Name,
		Email: reqBody.Email,
		Password: reqBody.Password,
		Address: reqBody.Address,
		Phone: reqBody.Phone,
		RegisteredAt: time.Now(),
	}

	// create
	result := r.DB.Create(&user)
	if result.Error != nil{
		return model.User{},result.Error
	}

	return user,nil
}