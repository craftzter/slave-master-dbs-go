package usecase

import (
	"log"

	"learn-api/internal/entity"
	"learn-api/internal/repository"
	bizerrors "learn-api/pkg/errors"
	"learn-api/pkg/validation"
)

type PostUsecase interface {
	CreatePost(userID int32, params entity.CreatePostParams) error
	GetAllPosts() ([]*entity.Post, error)
	GetPostByID(userID, postID int32) (*entity.Post, error)
	GetPostsByUserID(userID int32) ([]*entity.Post, error)
	UpdatePost(userID int32, params entity.UpdatePostParams) error
	DeletePost(userID, postID int32) error
}

type postUsecase struct {
	repo repository.PostRepository
}

func NewPostUsecase(repo repository.PostRepository) PostUsecase {
	return &postUsecase{repo: repo}
}

func (uc *postUsecase) CreatePost(userID int32, params entity.CreatePostParams) error {
	log.Printf("Creating post for user %d", userID)

	// Validasi user
	if err := validation.ValidateUserID(userID); err != nil {
		return bizerrors.WrapError("validation failed", err)
	}

	// Validasi title
	if err := validation.ValidatePostTitle(params.Title); err != nil {
		return bizerrors.WrapError("title validation failed", err)
	}

	// Validasi content
	if err := validation.ValidatePostContent(params.Content); err != nil {
		return bizerrors.WrapError("content validation failed", err)
	}

	// Business Logic: Set userID dari context auth
	params.UserID = userID

	post := &entity.Post{
		UserID:  params.UserID,
		Title:   params.Title,
		Content: params.Content,
		Image:   params.Image,
	}

	err := uc.repo.Create(post)
	if err != nil {
		log.Printf("Failed to create post: %v", err)
		return bizerrors.NewBusinessError("CREATE_FAILED", "failed to create post", err)
	}

	log.Printf("Post created successfully with ID %d", post.ID)
	return nil
}

func (uc *postUsecase) GetAllPosts() ([]*entity.Post, error) {
	log.Println("Fetching all posts")

	posts, err := uc.repo.GetAll()
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		return nil, bizerrors.WrapError("failed to retrieve posts", err)
	}

	// Business Logic: Filter posts (e.g., only public)
	// filtered := filterPublicPosts(posts)

	log.Printf("Retrieved %d posts", len(posts))
	return posts, nil
}

func (uc *postUsecase) GetPostByID(userID, postID int32) (*entity.Post, error) {
	log.Printf("Fetching post %d for user %d", postID, userID)

	if postID <= 0 {
		return nil, validation.ErrPostNotFound
	}

	post, err := uc.repo.GetByID(postID)
	if err != nil {
		log.Printf("Post %d not found: %v", postID, err)
		return nil, bizerrors.WrapError("post not found", err)
	}

	// Business Logic: Check ownership
	if post.UserID != int32(userID) {
		return nil, bizerrors.NewBusinessError("UNAUTHORIZED", "you can only view your own posts", nil)
	}

	log.Printf("Post %d retrieved", postID)
	return post, nil
}

func (uc *postUsecase) GetPostsByUserID(userID int32) ([]*entity.Post, error) {
	log.Printf("Fetching posts for user %d", userID)

	if err := validation.ValidateUserID(userID); err != nil {
		return nil, bizerrors.WrapError("invalid user", err)
	}

	posts, err := uc.repo.GetByUserID(userID)
	if err != nil {
		log.Printf("Failed to get posts for user %d: %v", userID, err)
		return nil, bizerrors.WrapError("failed to retrieve user posts", err)
	}

	log.Printf("Retrieved %d posts for user %d", len(posts), userID)
	return posts, nil
}

func (uc *postUsecase) UpdatePost(userID int32, params entity.UpdatePostParams) error {
	log.Printf("Updating post %d for user %d", params.ID, userID)

	// Validasi
	if err := validation.ValidatePostTitle(params.Title); err != nil {
		return bizerrors.WrapError("title validation failed", err)
	}
	if err := validation.ValidatePostContent(params.Content); err != nil {
		return bizerrors.WrapError("content validation failed", err)
	}

	// Check ownership
	existing, err := uc.repo.GetByID(params.ID)
	if err != nil {
		return bizerrors.WrapError("post not found", err)
	}
	if existing.UserID != userID {
		return bizerrors.NewBusinessError("UNAUTHORIZED", "you can only update your own posts", nil)
	}

	err = uc.repo.Update(params.ID, params.Title, params.Content, params.Image)
	if err != nil {
		log.Printf("Failed to update post %d: %v", params.ID, err)
		return bizerrors.NewBusinessError("UPDATE_FAILED", "failed to update post", err)
	}

	log.Printf("Post %d updated", params.ID)
	return nil
}

func (uc *postUsecase) DeletePost(userID, postID int32) error {
	log.Printf("Deleting post %d for user %d", postID, userID)

	if postID <= 0 {
		return validation.ErrPostNotFound
	}

	// Check ownership
	existing, err := uc.repo.GetByID(postID)
	if err != nil {
		return bizerrors.WrapError("post not found", err)
	}
	if existing.UserID != userID {
		return bizerrors.NewBusinessError("UNAUTHORIZED", "you can only delete your own posts", nil)
	}

	err = uc.repo.Delete(postID)
	if err != nil {
		log.Printf("Failed to delete post %d: %v", postID, err)
		return bizerrors.WrapError("failed to delete post", err)
	}

	log.Printf("Post %d deleted", postID)
	return nil
}
