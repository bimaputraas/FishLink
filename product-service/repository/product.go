package repository

import (
	"fishlink-product-service/model"
)

func (r *ProductRepositoryImpl) CreateProduct(product *model.Product) error {
    // insert product into database
    result := r.db.Create(product)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]*model.Product, error) {
    // retrieve products from database
    var products []*model.Product
    if err := r.db.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id uint) (*model.Product, error) {
    // retrieve product by ID from database
    var product model.Product
    if err := r.db.First(&product, id).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(product *model.Product) error {
    // update product in the database
    result := r.db.Save(product)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *ProductRepositoryImpl) DeleteProductByID(id uint) error {
    // delete product from the database
    result := r.db.Delete(&model.Product{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}