package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DATABASE_HOST = os.Getenv("DATABASE_HOST")
var DATABASE_PORT = os.Getenv("DATABASE_PORT")
var DATABASE_USER = os.Getenv("DATABASE_USER")
var DATABASE_PASS = os.Getenv("DATABASE_PASS")
var DATABASE_NAME = os.Getenv("DATABASE_NAME")

func NewGorm() *gorm.DB {
	gormDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", DATABASE_USER, DATABASE_PASS, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME)
	db, err := gorm.Open(postgres.Open(gormDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}