package models

import "time"

type Order struct {
	OrderId        int64     `json:"order_id"`
	UserId         int64     `json:"user_id"`
	TotalAmount    float64   `json:"total_amount"`
	PaymentStatus  string    `json:"payment_status"`
	ShippingStatus string    `json:"shipping_status"`
	CreatedAt      time.Time `json:"created_at"`
}
