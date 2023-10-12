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

// order
type orderRepository struct{
	gormDb *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository{
	return &orderRepository{gormDb: db}
}
