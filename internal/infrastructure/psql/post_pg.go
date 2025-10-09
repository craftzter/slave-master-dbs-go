package psql

import (
	"context"
	"database/sql"
	"learn-api/internal/db"
	"learn-api/internal/entity"
	"learn-api/internal/repository"
)

type postRepo struct {
	master *sql.DB
	slave  *sql.DB
}

func NewPostRepoPG(master, slave *sql.DB) repository.PostRepository {
	return &postRepo{
		master: master,
		slave:  slave,
	}
}

func (r *postRepo) Create(post *entity.Post) error {
	q := db.New(r.master)
	params := db.CreatePostParams{
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
		Image:   sql.NullString{String: post.Image, Valid: post.Image != ""},
	}
	return q.CreatePost(context.Background(), params)
}

func (r *postRepo) GetAll() ([]*entity.Post, error) {
	q := db.New(r.slave)
	posts, err := q.GetPosts(context.Background())
	if err != nil {
		return nil, err
	}
	var result []*entity.Post
	for _, p := range posts {
		result = append(result, &entity.Post{
			ID:        p.ID,
			UserID:    p.UserID,
			Title:     p.Title,
			Content:   p.Content,
			Image:     p.Image.String,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return result, nil
}

func (r *postRepo) GetByID(id int32) (*entity.Post, error) {
	q := db.New(r.slave)
	post, err := q.GetPostByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &entity.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Content:   post.Content,
		Image:     post.Image.String,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (r *postRepo) GetByUserID(userID int32) ([]*entity.Post, error) {
	q := db.New(r.slave)
	posts, err := q.GetPostsByUserID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	var result []*entity.Post
	for _, p := range posts {
		result = append(result, &entity.Post{
			ID:        p.ID,
			UserID:    p.UserID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return result, nil
}

func (r *postRepo) Update(id int32, title, content, image string) error {
	q := db.New(r.master)
	params := db.UpdatePostParams{
		ID:      id,
		Title:   title,
		Content: content,
		Image:   sql.NullString{String: image, Valid: image != ""},
	}
	return q.UpdatePost(context.Background(), params)
}

func (r *postRepo) Delete(id int32) error {
	q := db.New(r.master)
	return q.DeletePost(context.Background(), id)
}
