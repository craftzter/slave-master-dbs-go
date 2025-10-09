package usecase

import (
	"learn-api/internal/entity"
	"learn-api/internal/repository"
)

type ProfileUsecase interface {
	CreateProfile(params entity.CreateProfileParams) error
	GetProfileByUserID(userID int32) (*entity.Profile, error)
	UpdateProfile(params entity.UpdateProfileParams) error
}

type profileUsecase struct {
	repo repository.ProfileRepository
}

func NewProfileUsecase(repo repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{repo: repo}
}

func (uc *profileUsecase) CreateProfile(params entity.CreateProfileParams) error {
	profile := &entity.Profile{
		UserID:    params.UserID,
		Bio:       params.Bio,
		AvatarURL: params.AvatarURL,
	}
	return uc.repo.Create(profile)
}

func (uc *profileUsecase) GetProfileByUserID(userID int32) (*entity.Profile, error) {
	return uc.repo.GetByUserID(userID)
}

func (uc *profileUsecase) UpdateProfile(params entity.UpdateProfileParams) error {
	return uc.repo.Update(params.UserID, params.Bio, params.AvatarURL)
}
