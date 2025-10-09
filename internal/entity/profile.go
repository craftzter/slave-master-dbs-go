package entity

import "time"

type Profile struct {
	ID        int32     `json:"id" db:"id"`
	UserID    int32     `json:"user_id" db:"user_id"`
	Bio       string    `json:"bio" db:"bio"`
	AvatarURL string    `json:"avatar_url" db:"avatar_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateProfileParams struct {
	UserID    int32  `json:"user_id" validate:"required"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}

type UpdateProfileParams struct {
	UserID    int32  `json:"user_id" validate:"required"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}
