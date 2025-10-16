package handler

import (
	"net/http"

	"learn-api/internal/entity"
	"learn-api/internal/usecase"
	"learn-api/pkg/validation"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

// Register user
func (h *AuthHandler) Register(c *gin.Context) {
	var params entity.RegisterParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input check all requirement field"})
		return
	}

	err = validation.ValidateStruct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input check all requirement field"})
		return
	}

	err = h.uc.Register(params.Username, params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

// login user
func (h *AuthHandler) Login(c *gin.Context) {
	var params entity.LoginParams

	err := c.ShouldBindBodyWithJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input check all requirement input logins"})
		return
	}

	err = validation.ValidateStruct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input check all requirement input logins"})
		return
	}

	token, err := h.uc.Login(params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}
