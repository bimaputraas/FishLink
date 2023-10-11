package main

import (
	"final_project-ftgo-h8/api/router"

	_ "github.com/joho/godotenv/autoload"
)

func main(){
	// init app rest api
	app := router.NewEchoInstance()

	// run app
	app.Logger.Fatal(app.Start(":8080"))
}