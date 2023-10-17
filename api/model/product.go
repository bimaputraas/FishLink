package model

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64 `json:"price"`
	Stock       int     `json:"stock"`
}