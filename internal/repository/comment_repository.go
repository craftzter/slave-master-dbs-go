package repository

import (
	"learn-api/internal/entity"
)

type CommentRepository interface {
	Create(comment *entity.Comment) error
	GetByPostID(postID int32) ([]*entity.Comment, error)
	Delete(id int32) error
}
