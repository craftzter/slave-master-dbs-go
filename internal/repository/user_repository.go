package repository

import "learn-api/internal/entity"

type UserRepository interface {
	All() ([]*entity.User, error)
	GetByEmail(string) (*entity.User, error)
	GetByID(string) (*entity.User, error)
	Create(*entity.User) error
}
