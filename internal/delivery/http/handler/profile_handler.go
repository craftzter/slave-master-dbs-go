package handler

import (
	"net/http"
	"strconv"

	"learn-api/internal/delivery/http/middleware"
	"learn-api/internal/entity"
	"learn-api/internal/usecase"
	"learn-api/pkg/validation"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	uc usecase.ProfileUsecase
}

func NewProfileHandler(uc usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{uc: uc}
}

func (h *ProfileHandler) RegisterRoutes(r *gin.Engine) {
	profile := r.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	{
		profile.POST("", h.CreateProfile)
		profile.PUT("", h.UpdateProfile)
	}

	// public endpoint
	r.GET("profile/:userID", h.GetProfile)
	// r.POST("/profile", h.CreateProfile)
	// r.GET("/profile/:userID", h.GetProfile)
	// r.PUT("/profile", h.UpdateProfile)
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	userIDVal, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}
	var params entity.CreateProfileParams
	params.UserID = int32(userID)
	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err = validation.ValidateStruct(params)
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
	userIDVal, exist := c.Get("UserID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid userID"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	var params entity.UpdateProfileParams
	params.UserID = int32(userID)
	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	err = validation.ValidateStruct(params)
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
