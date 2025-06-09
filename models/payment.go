package models

import "time"

type Payment struct {
	PaymentId     int64     `json:"payment_id"`
	OrderId       int64     `json:"order_id"`
	UserId        int64     `json:"user_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	CreatedAt     time.Time `json:"created_at"`
}
