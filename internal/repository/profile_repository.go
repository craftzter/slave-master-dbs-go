package repository

import "learn-api/internal/entity"

type ProfileRepository interface {
	Create(profile *entity.Profile) error
	GetByUserID(userID int32) (*entity.Profile, error)
	Update(userID int32, bio, avatarURL string) error
}
