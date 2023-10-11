package repository

import "final_project-ftgo-h8/product-service/model"

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetAllProducts() ([]*model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProductByID(id uint) error
}