package models

import "time"

type User struct {
	UserId   int64     `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	Address  string    `json:"address"`
	Role     string    `json:"role" binding:"required"`
	CreateAt time.Time `json:"created_at"`
}
