package dto

type ReqBodyCreateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64 `json:"price"`
	Stock       int     `json:"stock"`
}

type ReqBodyUpdateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64 `json:"price"`
	Stock       int     `json:"stock"`
}