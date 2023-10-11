package repository

import (
	"gorm.io/gorm"
)

// user
type userRepository struct{
	gormDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{gormDb: db}
}
