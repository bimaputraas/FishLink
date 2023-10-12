package model

import "time"

type Order struct {
	Id        uint      `json:"id,omitempty"`
	UserId    uint      `json:"user_id,omitempty"`
	OrderDate time.Time `json:"order_date"`
}

type OrderDetail struct {
	Id        	uint      `json:"id,omitempty"`
	OrderId    	uint      `json:"order_id"`
	Order		Order	  `json:"order"`
	ProductId   uint 	  `json:"product_id"`
	Product		Product	  `json:"product"`
	Quantity    int   	  `json:"quantity"`
	TotalPrice  int64   `json:"total_price"`
}
