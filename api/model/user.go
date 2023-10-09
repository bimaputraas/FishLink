package model

import "time"

type User struct {
	Id           uint	`json:"id"`
	Name         string `json:"name"`
	Email        string `gorm:"unique" json:"email"`
	Password     string	`json:"password"`
	Address      string	`json:"address"`
	Phone        string	`json:"phone"`
	RegisteredAt time.Time	`json:"registered_at"`
}