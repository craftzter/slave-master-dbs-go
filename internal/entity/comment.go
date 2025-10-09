package entity

import "time"

type Comment struct {
	ID        int32     `json:"id"`
	PostID    int32     `json:"post_id" validate:"required"`
	UserID    int32     `json:"user_id" validate:"required"`
	Content   string    `json:"content" validate:"required,min=1,max=500"`
	CreatedAt time.Time `json:"created_at"`
}
