package router

import (
	"final_project-ftgo-h8/api/controller"
	"final_project-ftgo-h8/api/publisher"
	"final_project-ftgo-h8/api/repository"
	"final_project-ftgo-h8/config"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func StartEcho(){
	// init echo
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// init db
	dsn := os.Getenv("DSNGORM")
	gormDb := config.NewGorm(dsn)

	// init repository
	userRepository := repository.NewRepository(gormDb)

	// init chan
	channel := config.NewChannel()

	// add queue for email notification
	_ = config.AddQueue(channel,"fishlink-email_notification")

	// init publisher
	emailNotification := publisher.NewPublisher(channel)

	// init controller
	userController := controller.NewController(userRepository,emailNotification)

	// user route
	user := e.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.GET("/verification-register/:id/:code", userController.RegisterVerification)
	}
	
	
	e.Logger.Fatal(e.Start(":8080"))
}