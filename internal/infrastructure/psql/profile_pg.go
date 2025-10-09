package psql

import (
	"context"
	"database/sql"
	"learn-api/internal/db"
	"learn-api/internal/entity"
	"learn-api/internal/repository"
)

type profileRepo struct {
	master *sql.DB
	slave  *sql.DB
}

func NewProfileRepoPG(master, slave *sql.DB) repository.ProfileRepository {
	return &profileRepo{
		master: master,
		slave:  slave,
	}
}

func (r *profileRepo) Create(profile *entity.Profile) error {
	q := db.New(r.master)
	params := db.CreateProfileParams{
		UserID:    profile.UserID,
		Bio:       profile.Bio,
		AvatarUrl: profile.AvatarURL,
	}
	return q.CreateProfile(context.Background(), params)
}

func (r *profileRepo) GetByUserID(userID int32) (*entity.Profile, error) {
	q := db.New(r.slave)
	profile, err := q.GetProfileByUserID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return &entity.Profile{
		ID:        profile.ID,
		UserID:    profile.UserID,
		Bio:       profile.Bio,
		AvatarURL: profile.AvatarUrl,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
	}, nil
}

func (r *profileRepo) Update(userID int32, bio, avatarURL string) error {
	q := db.New(r.master)
	params := db.UpdateProfileParams{
		UserID:    userID,
		Bio:       bio,
		AvatarUrl: avatarURL,
	}
	return q.UpdateProfile(context.Background(), params)
}
