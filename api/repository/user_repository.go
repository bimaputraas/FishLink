package repository

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"
	"time"
)

func (r *userRepository) InsertUser(reqBody dto.ReqUserRegister) (model.User,error){
	// model
	user := model.User{
		Name: reqBody.Name,
		Email: reqBody.Email,
		Password: reqBody.Password,
		Address: reqBody.Address,
		Phone: reqBody.Phone,
		Status: "Pending Verification",
		RegisteredAt: time.Now(),
	}

	// create
	result := r.DB.Create(&user)
	if result.Error != nil{
		return model.User{},result.Error
	}

	return user,nil
}

func (r *userRepository) InsertUserVerification(userId uint, code string) error{
	// model
	userVerif := model.UserVerification{
		UserID: userId,
		VerificationCode: code,
	}

	// create
	result := r.DB.Create(&userVerif)
	if result.Error != nil{
		return result.Error
	}

	return nil
}

func (r *userRepository) UpdateUserStatusByIdAndCode(userId uint, code string) error{
	// find user verification
	userVerif := model.UserVerification{
		UserID: userId,
		VerificationCode: code,
	}
	result := r.DB.First(&userVerif)
	if result.Error != nil{
		return result.Error
	}

	// find user
	user := model.User{
		Id: userVerif.UserID,
	}
	result = r.DB.First(&user)
	if result.Error != nil{
		return result.Error
	}

	// update user
	user.Status = "Verified"
	result = r.DB.Save(&user)
	if result.Error != nil{
		return result.Error
	}

	return nil
}


func (r *userRepository) FindUserByEmail(email string) (model.User,error){
	// model
	user := model.User{}

	// create
	result := r.DB.Where("email = ?",email).First(&user)
	if result.Error != nil{
		return model.User{},result.Error
	}

	return user,nil
}

