package model

import "time"

type User struct {
	Id           uint	`json:"id,omitempty"`
	Name         string `json:"name"`
	Email        string `gorm:"unique" json:"email"`
	Password     string	`json:"password,omitempty"`
	Address      string	`json:"address"`
	Phone        string	`json:"phone"`
	Status		 string `json:"status"`
	Amount		 float64 `json:"amount"`
	Role 		 string `json:"role"`
	RegisteredAt time.Time	`json:"registered_at"`
}

type UserVerification struct {
	Id           		uint	`json:"id"`
	UserID				uint 	`json:"user_id"`
	VerificationCode 	string `json:"verification_code"`
}