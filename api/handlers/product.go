package handlers

import (
    "context"
    "fmt"
    "strconv"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "gorm.io/gorm"

    "final_project-ftgo-h8/api/model"
    pb "final_project-ftgo-h8/pb"
)

type ProductServiceImpl struct {
    db *gorm.DB
    pb.ProductServiceServer
}

func NewProductHandler(db *gorm.DB) *ProductServiceImpl {
    return &ProductServiceImpl{db: db}
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	newProduct := req.GetProduct()

	if newProduct.GetName() == "" || newProduct.GetPrice() <= 0 || newProduct.GetStock() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product data")
	}

	product := &model.Product{
		Name:        newProduct.GetName(),
		Description: newProduct.GetDescription(),
		Price:       newProduct.GetPrice(),
		Stock:       newProduct.GetStock(),
	}

	if err := s.db.Create(product).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create product: %v", err)
	}

	createdProduct := &pb.Product{
		Id:          fmt.Sprintf("%d", product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	return createdProduct, nil
}

func (s *ProductServiceImpl) GetAllProduct(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
    var products []*model.Product

    if err := s.db.Find(&products).Error; err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to retrieve products: %v", err)
    }

    productResponses := make([]*pb.Product, len(products))
    for i, product := range products {
        productResponses[i] = &pb.Product{
            Id:          fmt.Sprintf("%d", product.ID),
            Name:        product.Name,
            Description: product.Description,
            Price:       product.Price,
            Stock:       product.Stock,
        }
    }

    response := &pb.GetAllProductResponse{
        Products: productResponses,
    }

    return response, nil
}

func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
    productID := req.GetId()

    var product model.Product

    if err := s.db.First(&product, productID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, status.Errorf(codes.NotFound, "model.Product with ID %s not found", productID)
        }
        return nil, status.Errorf(codes.Internal, "Failed to retrieve product: %v", err)
    }

    response := &pb.Product{
        Id:          fmt.Sprintf("%d", product.ID),
        Name:        product.Name,
        Description: product.Description,
        Price:       product.Price,
        Stock:       product.Stock,
    }

    return response, nil
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
    updatedProduct := req.GetProduct()

    if updatedProduct.GetId() == "" || updatedProduct.GetName() == "" || updatedProduct.GetPrice() <= 0 || updatedProduct.GetStock() < 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid product data")
    }

    productID, err := strconv.Atoi(updatedProduct.GetId())
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid product ID")
    }

    var existingProduct model.Product
    if err := s.db.First(&existingProduct, productID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, status.Errorf(codes.NotFound, "model.Product with ID %s not found", updatedProduct.GetId())
        }
        return nil, status.Errorf(codes.Internal, "Failed to retrieve product: %v", err)
    }

    existingProduct.Name = updatedProduct.GetName()
    existingProduct.Description = updatedProduct.GetDescription()
    existingProduct.Price = updatedProduct.GetPrice()
    existingProduct.Stock = updatedProduct.GetStock()

    if err := s.db.Save(&existingProduct).Error; err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to update product: %v", err)
    }

    updatedResponse := &pb.Product{
        Id:          fmt.Sprintf("%d", existingProduct.ID),
        Name:        existingProduct.Name,
        Description: existingProduct.Description,
        Price:       existingProduct.Price,
        Stock:       existingProduct.Stock,
    }

    return updatedResponse, nil
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error) {
    productID := req.GetId()

    productIDInt, err := strconv.Atoi(productID)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid product ID")
    }

    var existingProduct model.Product
    if err := s.db.First(&existingProduct, productIDInt).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, status.Errorf(codes.NotFound, "model.Product with ID %s not found", productID)
        }
        return nil, status.Errorf(codes.Internal, "Failed to retrieve product: %v", err)
    }

    if err := s.db.Delete(&existingProduct).Error; err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to delete product: %v", err)
    }

    deletedResponse := &pb.Product{
        Id:          fmt.Sprintf("%d", existingProduct.ID),
        Name:        existingProduct.Name,
        Description: existingProduct.Description,
        Price:       existingProduct.Price,
        Stock:       existingProduct.Stock,
    }

    return deletedResponse, nil
}