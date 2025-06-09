package models

type OrderItem struct {
	OrderItemId int64   `json:"order_item_id"`
	OrderId     int64   `json:"order_id"`
	ProductId   int64   `json:"product_id"`
	Quantity    int32   `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}
