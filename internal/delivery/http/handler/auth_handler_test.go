package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthUsecase is a mock implementation of AuthUsecase
type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockAuthUsecase) Register(username, email, password string) error {
	args := m.Called(username, email, password)
	return args.Error(0)
}

func TestAuthHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUC := new(MockAuthUsecase)
	handler := NewAuthHandler(mockUC)

	r := gin.Default()
	handler.RegisterRoutes(r)

	t.Run("successful registration", func(t *testing.T) {
		mockUC.On("Register", "testuser", "test@example.com", "password123").Return(nil).Once()

		body := map[string]string{
			"username": "testuser",
			"email":    "test@example.com",
			"password": "password123",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "registered")
		mockUC.AssertExpectations(t)
	})

	t.Run("registration with invalid input", func(t *testing.T) {
		body := map[string]string{
			"username": "",
			"email":    "invalid-email",
			"password": "123",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid input")
	})

	t.Run("registration failure", func(t *testing.T) {
		mockUC.On("Register", "testuser", "test@example.com", "password123").Return(assert.AnError).Once()

		body := map[string]string{
			"username": "testuser",
			"email":    "test@example.com",
			"password": "password123",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
		mockUC.AssertExpectations(t)
	})
}

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUC := new(MockAuthUsecase)
	handler := NewAuthHandler(mockUC)

	r := gin.Default()
	handler.RegisterRoutes(r)

	t.Run("successful login", func(t *testing.T) {
		mockUC.On("Login", "test@example.com", "password123").Return("jwt-token", nil).Once()

		body := map[string]string{
			"email":    "test@example.com",
			"password": "password123",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "jwt-token")
		mockUC.AssertExpectations(t)
	})

	t.Run("login with invalid input", func(t *testing.T) {
		body := map[string]string{
			"email":    "",
			"password": "",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid input")
	})

	t.Run("login failure", func(t *testing.T) {
		mockUC.On("Login", "test@example.com", "wrongpassword").Return("", assert.AnError).Once()

		body := map[string]string{
			"email":    "test@example.com",
			"password": "wrongpassword",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
		mockUC.AssertExpectations(t)
	})
}
