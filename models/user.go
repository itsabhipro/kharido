package models

import "time"

type User struct {
	UserId   int64     `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `binding:required`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Role     string    `json:"role"`
	CreateAt time.Time `json:"created_at"`
}
