package models

import "time"

type Shipping struct {
	ShippingId        int64     `json:"shipping_id"`
	OrderId           int64     `json:"order_id"`
	TrackingNumber    string    `json:"tracking_number"`
	ShippingAddress   string    `json:"shipping_address"`
	EstimatedDelivery time.Time `json:"estimated_delivery"`
	CurrentStatus     string    `json:"current_status"`
}
