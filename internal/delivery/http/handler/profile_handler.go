package handler

import (
	"net/http"
	"strconv"

	"learn-api/internal/entity"
	"learn-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProfileHandler struct {
	uc usecase.ProfileUsecase
}

func NewProfileHandler(uc usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{uc: uc}
}

func (h *ProfileHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/profile", h.CreateProfile)
	r.GET("/profile/:userID", h.GetProfile)
	r.PUT("/profile", h.UpdateProfile)
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var params entity.CreateProfileParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err = h.uc.CreateProfile(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "profile created"})
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	profile, err := h.uc.GetProfileByUserID(int32(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var params entity.UpdateProfileParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err = h.uc.UpdateProfile(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "profile updated"})
}
