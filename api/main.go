package main

import (
	"final_project-ftgo-h8/api/router"
	"final_project-ftgo-h8/helper"
)

func main(){
	// load env
	helper.LoadEnv()

	// init app rest api
	app := router.NewEchoInstance()

	// run app
	app.Logger.Fatal(app.Start(":8080"))
}