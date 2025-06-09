package models

import "time"

type Product struct {
	ProductId   int64     `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int32     `json:"stock"`
	CategoryId  int64     `json:"category_id"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
}
