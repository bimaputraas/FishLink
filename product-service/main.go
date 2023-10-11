package main

import (
	"final_project-ftgo-h8/config"
	"final_project-ftgo-h8/helper"
	"final_project-ftgo-h8/pb"
	"final_project-ftgo-h8/product-service/repository"
	"final_project-ftgo-h8/product-service/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
    // load env
    helper.LoadEnv()

	// init db connection
    db := config.NewGorm(os.Getenv("DSNGORM"))

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Create a new ProductRepository
    productRepo := repository.NewProductRepository(db)

    // Create a new ProductServer
    productServer := server.NewProductServer(productRepo)

    // Register the ProductServiceServer with the gRPC server
    pb.RegisterProductServiceServer(grpcServer, productServer)

    // Listen on a port
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    log.Println("Server is listening on port 50051")

    // Start the gRPC server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}