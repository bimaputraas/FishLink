package repository

import "gorm.io/gorm"

type UserRepository interface{}

type userRepository struct{
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository{
	return &userRepository{DB: db}
}
