package server

import (
	pb "fishlink-product-service/pb"
	"fishlink-product-service/repository"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	repo repository.ProductRepository
}

func NewProductServer(repo repository.ProductRepository) pb.ProductServiceServer {
	return &ProductServer{repo: repo}
}