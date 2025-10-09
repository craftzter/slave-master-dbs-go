package entity

import "time"

type Notifications struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Type      string    `json:"type" validate:"required"`
	Message   string    `json:"message" validate:"required"`
	RelatedID int32     `json:"related_id" validate:"required"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
