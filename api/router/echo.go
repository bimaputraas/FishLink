package router

import (
	"final_project-ftgo-h8/api/controller"
	"final_project-ftgo-h8/api/repository"
	"final_project-ftgo-h8/config"
	"final_project-ftgo-h8/pb"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func StartEcho(){
	// init echo
	e := echo.New()

	// init db
	dsn := os.Getenv("DSNGORM")
	gormDb := config.NewGorm(dsn)

	// init repository
	userRepository := repository.NewRepository(gormDb)

	// init controller
	userController := controller.NewController(userRepository)

	// user route
	user := e.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.GET("/verification-register/:id/:code", userController.RegisterVerification)
	}

	// init gRPC connection
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	// init gRPC client
	grpcClient := pb.NewProductServiceClient(grpcConn)

	// init ProductController with the gRPC client
	productController := controller.NewProductController(grpcClient)
	
	product := e.Group("/product")
	{
		product.POST("", productController.CreateProduct)
		product.GET("/:id", productController.GetProduct)
		product.GET("", productController.GetAllProducts)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
	}
	
	e.Logger.Fatal(e.Start(":8080"))
}