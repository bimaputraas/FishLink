package router

import (
	"fishlink-mainapi/config"
	"fishlink-mainapi/controller"
	"fishlink-mainapi/middleware"
	"fishlink-mainapi/pb"
	"fishlink-mainapi/publisher"
	"fishlink-mainapi/repository"
	"log"
	"net/http"

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
	gormDb := config.NewGorm()

	// init repository
	userRepository := repository.NewUserRepository(gormDb)
	orderRepository := repository.NewOrderRepository(gormDb)

	// init chan
	channel := config.NewChannel()

	// add queue for email notification
	_ = config.AddQueue(channel, "fishlink-email_notification")

	// init publisher
	emailNotification := publisher.NewPublisher(channel)

	// init controller
	userController := controller.NewUserController(userRepository, emailNotification)
	orderController := controller.NewOrderController(orderRepository)

	// init authentication middleware
	authMiddleware := middleware.NewAuthenticationMiddleware(userRepository)

	// init gRPC connection
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	// defer grpcConn.Close()

	// init gRPC client
	grpcClient := pb.NewProductServiceClient(grpcConn)

	// init ProductController with the gRPC client
	productController := controller.NewProductController(grpcClient)

	// user gateaway (before login)
	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)
	e.GET("user-verification-register/:id/:code", userController.RegisterVerification)
	e.GET("/product/:id", productController.GetProduct)
	e.GET("/product", productController.GetAllProducts)

	// user route (after login)
	user := e.Group("/user", authMiddleware.Authentication)
	{
		user.GET("/info", userController.GetInfo)
		user.PUT("/top-up", userController.TopUp)
		user.POST("/order",orderController.NewOrder)
		user.GET("/order",orderController.GetOrders)
	}

	// product route - admin
	product := e.Group("/product")
	{
		product.POST("", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
	}

	return e
}
