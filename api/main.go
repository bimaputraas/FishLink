package main

import (
	"final_project-ftgo-h8/api/router"

	_ "github.com/joho/godotenv/autoload"
)

func main(){


	// run rest api
	router.StartEcho()
}