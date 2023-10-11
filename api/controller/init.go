package controller

import (
	"final_project-ftgo-h8/api/publisher"
	"final_project-ftgo-h8/api/repository"
	"final_project-ftgo-h8/pb"
)

// user controller

type userController struct{
	repository repository.UserRepository
	publisher publisher.Publisher
}

func NewController(r repository.UserRepository, p publisher.Publisher) UserController{
	return &userController{
		repository: r,
		publisher: p,
	}
}

// product controller

type productController struct {
	Service pb.ProductServiceClient
}

func NewProductController(pb pb.ProductServiceClient) ProductController {
	return &productController{Service: pb}
}