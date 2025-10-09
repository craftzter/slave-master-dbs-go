package entity

import "time"

type Post struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostParams struct {
	UserID  int32  `json:"user_id"`
	Title   string `json:"title" validate:"required,min=1,max=255"`
	Content string `json:"content" validate:"required"`
	Image   string `json:"image"`
}

type UpdatePostParams struct {
	ID      int32  `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required,min=1,max=255"`
	Content string `json:"content" validate:"required"`
	Image   string `json:"image"`
}
