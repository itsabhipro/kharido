package models

import "time"

type Review struct {
	ReviewId   int64     `json:"review_id"`
	UserId     int64     `json:"user_id"`
	ProductId  int64     `json:"product_id"`
	Rating     int       `json:"rating"`
	ReviewText string    `json:"review_text"`
	CreatedAt  time.Time `json:"create_at"`
}
