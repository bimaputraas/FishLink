package model

import "time"

type Order struct {
	Id        uint      `json:"id,omitempty"`
	UserId    uint      `json:"user_id,omitempty"`
	OrderDate time.Time `json:"order_date"`
}

type OrderDetail struct {
	Id        	uint      `json:"id,omitempty"`
	UserId    	uint      `json:"user_id"`
	ProductId    uint   `json:"product_id"`
	Quantity    int   `json:"quantity"`
	TotalPrice    float64   `json:"total_price"`
}

// CREATE TABLE orders (
//     id SERIAL PRIMARY KEY,
//     user_id INT NOT NULL,
//     order_date TIMESTAMP,
//     FOREIGN KEY (user_id) REFERENCES users(id)
// );

// CREATE TABLE order_details (
//     id SERIAL PRIMARY KEY,
//     order_id INT NOT NULL,
//     product_id INT NOT NULL,
//     quantity INT,
//     total_price FLOAT,
//     FOREIGN KEY (order_id) REFERENCES orders(id),
//     FOREIGN KEY (product_id) REFERENCES products(id)
// );