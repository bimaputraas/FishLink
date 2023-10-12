package repository

import (
	"final_project-ftgo-h8/api/dto"
	"final_project-ftgo-h8/api/model"
	"time"
)

func (r *userRepository) InsertUser(reqBody dto.ReqBodyRegister) (model.User,error){
	// model
	user := model.User{
		Name: reqBody.Name,
		Email: reqBody.Email,
		Password: reqBody.Password,
		Address: reqBody.Address,
		Phone: reqBody.Phone,
		Status: "Pending Verification",
		Role: reqBody.Role,
		RegisteredAt: time.Now(),
	}

	// create
	result := r.gormDb.Create(&user)
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
	result := r.gormDb.Create(&userVerif)
	if result.Error != nil{
		return result.Error
	}

	return nil
}

func (r *userRepository) UpdateUserStatusByIdAndCode(userId uint, code string) (model.User,error){
	// check existing & validate code user verification
	userVerif := model.UserVerification{}
	result := r.gormDb.Where("user_id = ? and verification_code = ?",userId,code).First(&userVerif)
	if result.Error != nil{
		return model.User{},result.Error
	}

	// find user
	user := model.User{}
	result = r.gormDb.First(&user,userId)
	if result.Error != nil{
		return model.User{},result.Error
	}

	// update user
	user.Status = "Verified"
	result = r.gormDb.Save(&user)
	if result.Error != nil{
		return model.User{},result.Error
	}

	return user,nil
}


func (r *userRepository) FindUserByEmail(email string) (model.User,error){
	// model
	user := model.User{}

	// find
	result := r.gormDb.Where("email = ?",email).First(&user)
	if result.Error != nil{
		return model.User{},result.Error
	}

	return user,nil
}

func (r *userRepository) FindUserById(userId uint) (model.User,error){
	// model
	user := model.User{}

	// find
	result := r.gormDb.First(&user,userId)
	if result.Error != nil{
		return model.User{},result.Error
	}

	return user,nil
}

func (r *userRepository) UpdateAmount(user model.User, amount float64) (model.User, error) {
	user.Amount += amount

	result := r.gormDb.Save(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}