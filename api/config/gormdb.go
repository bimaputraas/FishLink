package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewGorm() *gorm.DB {
	// load env
	gormDSN := os.Getenv("DBURL_POSTGRES")
	
	// gormDSN := "postgresql://postgres:SHd6S6PqrLIJzNi1YhOA@containers-us-west-198.railway.app:5646/railway"
	db, err := gorm.Open(postgres.Open(gormDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}