package config

import (
	"final_project-ftgo-h8/product-service/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=timothy dbname=final port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    db.AutoMigrate(&model.Product{})

    return db, nil
}