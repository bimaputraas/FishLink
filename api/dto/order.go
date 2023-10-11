package dto

type ReqBodyNewOrder struct {
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type ResDetailNewOrder struct {
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	Ordered_at  string  `json:"ordered_at"`
}

type ResDetailGetOrder struct {
	ProductName        string  `json:"product_name"`
	ProductPrice       float64 `json:"product_price"`
	ProductDescription string  `json:"product_description"`
	Quantity           int     `json:"quantity"`
	TotalPrice         float64 `json:"total_price"`
	Ordered_at         string  `json:"ordered_at"`
}
