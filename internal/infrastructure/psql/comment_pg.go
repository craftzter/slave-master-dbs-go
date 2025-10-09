package psql

import (
	"context"
	"database/sql"

	"learn-api/internal/db"
	"learn-api/internal/entity"
	"learn-api/internal/repository"
)

type commentRepo struct {
	master *sql.DB
	slave  *sql.DB
}

func NewCommentRepoPG(master, slave *sql.DB) repository.CommentRepository {
	return &commentRepo{
		master: master,
		slave:  slave,
	}
}

func (r *commentRepo) Create(comment *entity.Comment) error {
	q := db.New(r.master)
	params := db.CreateCommentParams{
		PostID:  comment.PostID,
		UserID:  comment.UserID,
		Content: comment.Content,
	}
	return q.CreateComment(context.Background(), params)
}

func (r *commentRepo) GetByPostID(postID int32) ([]*entity.Comment, error) {
	q := db.New(r.slave)
	comments, err := q.GetCommentsByPostID(context.Background(), postID)
	if err != nil {
		return nil, err
	}
	var result []*entity.Comment
	for _, c := range comments {
		result = append(result, &entity.Comment{
			ID:        c.ID,
			PostID:    c.PostID,
			UserID:    c.UserID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		})
	}
	return result, nil
}

func (r *commentRepo) Delete(id int32) error {
	q := db.New(r.master)
	return q.DeleteComment(context.Background(), id)
}
