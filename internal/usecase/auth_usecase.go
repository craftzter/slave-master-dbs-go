package usecase

import (
	"errors"

	"learn-api/internal/entity"
	"learn-api/internal/repository"
	"learn-api/pkg/hash"
	"learn-api/pkg/jwt"
)

type AuthUsecase interface {
	Login(string, string) (string, error)
	Register(string, string, string) error
}


type authUsecase struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(repo repository.UserRepository, jwtSecret string) AuthUsecase {
	return &authUsecase{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (uc *authUsecase) Login(email, password string) (string, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}
	if user == nil {
		return "", errors.New("email not found")
	}
	err = hash.CheckPassword(user.Password, password)
	if err != nil {
		return "", errors.New("invalid password")
	}
	token, err := jwt.GenerateToken(uint(user.ID), uc.jwtSecret)
	if err != nil {
		return "", errors.New("failed to signing token")
	}
	return token, nil
}

func (uc *authUsecase) Register(username, email, password string) error {
	findUser, _ := uc.repo.GetByEmail(email)
	if findUser != nil {
		return errors.New("email already exists")
	}
	hashPaswd, err := hash.HashPassword(password)
	if err != nil {
		return errors.New("failed to hash a password")
	}

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: hashPaswd,
	}
	err = uc.repo.Create(user)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
