package router

import (
	"final_project-ftgo-h8/api/controller"
	"final_project-ftgo-h8/api/middleware"
	"final_project-ftgo-h8/api/publisher"
	"final_project-ftgo-h8/api/repository"
	"final_project-ftgo-h8/config"
	"final_project-ftgo-h8/pb"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func NewEchoInstance() *echo.Echo{
	// init echo
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// init db
	dsn := os.Getenv("DSNGORM")
	gormDb := config.NewGorm(dsn)

	// init repository
	userRepository := repository.NewUserRepository(gormDb)

	// init chan
	channel := config.NewChannel()

	// add queue for email notification
	_ = config.AddQueue(channel, "fishlink-email_notification")

	// init publisher
	emailNotification := publisher.NewPublisher(channel)

	// init controller
	userController := controller.NewController(userRepository, emailNotification)

	// init authentication middleware
	authMiddleware := middleware.NewAuthenticationMiddleware(userRepository)

	// user gateaway (before login)
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
	e.GET("user-verification-register/:id/:code", userController.RegisterVerification)
	// user route (after login)
	user := e.Group("/user",authMiddleware.Authentication)
	{
		user.GET("/info", userController.GetInfo)
		user.PUT("/top-up", userController.TopUp)
		user.POST("/order",func(c echo.Context) error{return nil})
		user.GET("/order",func(c echo.Context) error{return nil})
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

	// product route
	product := e.Group("/product")
	{
		product.POST("", productController.CreateProduct)
		product.GET("/:id", productController.GetProduct)
		product.GET("", productController.GetAllProducts)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
	}

	return e
}
