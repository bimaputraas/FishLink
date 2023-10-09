package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(gormDSN string) *gorm.DB{
	db, err := gorm.Open(postgres.Open(gormDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
