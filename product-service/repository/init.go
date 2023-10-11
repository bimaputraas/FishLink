package repository

import (
	"final_project-ftgo-h8/product-service/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetAllProducts() ([]*model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProductByID(id uint) error
}
type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}