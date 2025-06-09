package models

import "time"

type Notification struct {
	NotificationId int64     `json:"notification_id"`
	UserId         int64     `json:"user_id"`
	Message        string    `json:"message"`
	IsRead         bool      `json:"is_read"`
	CreatedAt      time.Time `json:"created_at"`
}
