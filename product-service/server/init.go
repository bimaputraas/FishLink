package server

import (
	pb "final_project-ftgo-h8/pb"
	"final_project-ftgo-h8/product-service/repository"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	repo repository.ProductRepository
}

func NewProductServer(repo repository.ProductRepository) pb.ProductServiceServer {
	return &ProductServer{repo: repo}
}