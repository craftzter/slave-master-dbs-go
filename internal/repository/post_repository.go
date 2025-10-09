package repository

import "learn-api/internal/entity"

type PostRepository interface {
	Create(post *entity.Post) error
	GetAll() ([]*entity.Post, error)
	GetByID(id int32) (*entity.Post, error)
	GetByUserID(userID int32) ([]*entity.Post, error)
	Update(id int32, title, content, image string) error
	Delete(id int32) error
}
