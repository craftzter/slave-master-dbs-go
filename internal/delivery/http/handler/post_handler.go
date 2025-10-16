package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"learn-api/internal/delivery/http/middleware"
	"learn-api/internal/entity"
	"learn-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	uc usecase.PostUsecase
}

func NewPostHandler(uc usecase.PostUsecase) *PostHandler {
	return &PostHandler{uc: uc}
}

func (h *PostHandler) RegisterRoutes(r *gin.Engine) {
	posts := r.Group("/posts")
	posts.Use(middleware.AuthMiddleware())
	{
		posts.POST("", h.CreatePost)
		posts.PUT("/:id", h.UpdatePost)
		posts.DELETE("/:id", h.DeletePost)
	}
	r.GET("/posts", h.GetAllPosts)     // Public
	r.GET("/posts/:id", h.GetPostByID) // Public for now
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID := userIDVal.(uint)

	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form"})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and content are required"})
		return
	}

	var imagePath string
	files := form.File["image"]
	if len(files) > 0 {
		file := files[0]
		// Validate file type
		fileHeader, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
			return
		}
		defer fileHeader.Close()
		buff := make([]byte, 512)
		_, err = fileHeader.Read(buff)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read content"})
			return
		}

		mimeType := http.DetectContentType(buff)
		if mimeType != "image/jpeg" && mimeType != "image/png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "only JPEG, PNG and JPG are allowed"})
			return
		}
		if file.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file exceeds 5MB"})
			return
		}
		// Save file
		filename := fmt.Sprintf("%d_%d_%s", userID, time.Now().Unix(), filepath.Base(file.Filename))
		imagePath = "/uploads/" + filename
		if err = c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
			return
		}
	}

	params := entity.CreatePostParams{
		UserID:  int32(userID),
		Title:   title,
		Content: content,
		Image:   imagePath,
	}

	err = h.uc.CreatePost(int32(userID), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "post created"})
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.uc.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	post, err := h.uc.GetPostByID(int32(userID), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form"})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and content are required"})
		return
	}

	var imagePath string
	files := form.File["image"]
	if len(files) > 0 {
		file := files[0]
		// Validate file type
		fileHeader, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
			return
		}
		defer fileHeader.Close()
		buff := make([]byte, 512)
		_, err = fileHeader.Read(buff)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failes to read content"})
			return
		}

		mimeType := http.DetectContentType(buff)
		if mimeType != "image/jpeg" && mimeType != "image/png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "only JPEG, PNG and JPG are allowed"})
			return
		}
		if file.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file exceeds 5MB"})
			return
		}
		// Save file
		filename := fmt.Sprintf("%d_%d_%s", userID, time.Now().Unix(), filepath.Base(file.Filename))
		imagePath = "/uploads/" + filename
		if err = c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
			return
		}

	}

	params := entity.UpdatePostParams{
		ID:      int32(id),
		Title:   title,
		Content: content,
		Image:   imagePath,
	}

	err = h.uc.UpdatePost(int32(userID), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post updated"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	err = h.uc.DeletePost(int32(userID), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}
