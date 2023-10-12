package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewGorm() *gorm.DB {
	// dsn data
	DATABASE_HOST := os.Getenv("DATABASE_HOST")
 	DATABASE_PORT := os.Getenv("DATABASE_PORT")
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASS := os.Getenv("DATABASE_PASS")
 	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	// load env
	gormDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", DATABASE_USER, DATABASE_PASS, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME)
	
	// gormDSN := "postgresql://postgres:SHd6S6PqrLIJzNi1YhOA@containers-us-west-198.railway.app:5646/railway"
	db, err := gorm.Open(postgres.Open(gormDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}