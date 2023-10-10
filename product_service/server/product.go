package server

import (
    "context"
    "strconv"

    pb "final_project-ftgo-h8/pb"
    "final_project-ftgo-h8/product_service/model"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "gorm.io/gorm"
)

func (s *ProductServer) ValidateProduct(product *pb.Product) error {
    if product.GetName() == "" || product.GetPrice() <= 0 || product.GetStock() < 0 {
        return status.Error(codes.InvalidArgument, "Invalid product data")
    }
    return nil
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
    newProduct := req.GetProduct()

    if err := s.ValidateProduct(newProduct); err != nil {
        return nil, err
    }

    product := &model.Product{
        Name:        newProduct.GetName(),
        Description: newProduct.GetDescription(),
        Price:       float32(newProduct.GetPrice()),
        Stock:       int32(newProduct.GetStock()),
    }

    if err := s.repo.CreateProduct(product); err != nil {
        return nil, status.Error(codes.Internal, "Failed to create product")
    }

    createdProduct := &pb.Product{
        Id:          strconv.FormatUint(uint64(product.ID), 10),
        Name:        product.Name,
        Description: product.Description,
        Price:       float32(product.Price),
        Stock:       int32(product.Stock),
    }

    return createdProduct, nil
}

func (s *ProductServer) GetAllProduct(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
    products, err := s.repo.GetAllProducts()
    if err != nil {
        return nil, status.Error(codes.Internal, "Failed to retrieve products")
    }

    var productResponses []*pb.Product
    for _, product := range products {
        productResponses = append(productResponses, &pb.Product{
            Id:          strconv.FormatUint(uint64(product.ID), 10),
            Name:        product.Name,
            Description: product.Description,
            Price:       float32(product.Price),
            Stock:       int32(product.Stock),
        })
    }

    response := &pb.GetAllProductResponse{
        Products: productResponses,
    }

    return response, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
    productID := req.GetId()

    id, err := strconv.ParseUint(productID, 10, 64)
    if err != nil {
        return nil, status.Error(codes.InvalidArgument, "Invalid product ID")
    }

    product, err := s.repo.GetProductByID(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, status.Error(codes.NotFound, "Product not found")
        }
        return nil, status.Error(codes.Internal, "Failed to retrieve product")
    }

    productResponse := &pb.Product{
        Id:          strconv.FormatUint(uint64(product.ID), 10),
        Name:        product.Name,
        Description: product.Description,
        Price:       float32(product.Price),
        Stock:       int32(product.Stock),
    }

    return productResponse, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
    updatedProduct := req.GetProduct()

    if err := s.ValidateProduct(updatedProduct); err != nil {
        return nil, err
    }

    productID, err := strconv.ParseUint(updatedProduct.GetId(), 10, 64)
    if err != nil {
        return nil, status.Error(codes.InvalidArgument, "Invalid product ID")
    }

    existingProduct, err := s.repo.GetProductByID(uint(productID))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, status.Error(codes.NotFound, "Product not found")
        }
        return nil, status.Error(codes.Internal, "Failed to retrieve product")
    }

    existingProduct.Name = updatedProduct.GetName()
    existingProduct.Description = updatedProduct.GetDescription()
    existingProduct.Price = float32(updatedProduct.GetPrice())
    existingProduct.Stock = int32(updatedProduct.GetStock())

    if err := s.repo.UpdateProduct(existingProduct); err != nil {
        return nil, status.Error(codes.Internal, "Failed to update product")
    }

    updatedResponse := &pb.Product{
        Id:          strconv.FormatUint(uint64(existingProduct.ID), 10),
        Name:        existingProduct.Name,
        Description: existingProduct.Description,
        Price:       float32(existingProduct.Price),
        Stock:       int32(existingProduct.Stock),
    }

    return updatedResponse, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error) {
    productID := req.GetId()

    id, err := strconv.ParseUint(productID, 10, 64)
    if err != nil {
        return nil, status.Error(codes.InvalidArgument, "Invalid product ID")
    }

    existingProduct, err := s.repo.GetProductByID(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, status.Error(codes.NotFound, "Product not found")
        }
        return nil, status.Error(codes.Internal, "Failed to retrieve product")
    }

    if err := s.repo.DeleteProductByID(uint(id)); err != nil {
        return nil, status.Error(codes.Internal, "Failed to delete product")
    }

    deletedResponse := &pb.Product{
        Id:          strconv.FormatUint(uint64(existingProduct.ID), 10),
        Name:        existingProduct.Name,
        Description: existingProduct.Description,
        Price:       float32(existingProduct.Price),
        Stock:       int32(existingProduct.Stock),
    }

    return deletedResponse, nil
}