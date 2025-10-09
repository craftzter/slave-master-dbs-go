package usecase

import (
	"testing"

	"learn-api/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) All() ([]*entity.User, error) {
	args := m.Called()
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestAuthUsecase_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	uc := NewAuthUsecase(mockRepo, "secret")

	t.Run("successful login", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := &entity.User{
			ID:       1,
			Username: "testuser",
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}
		mockRepo.On("GetByEmail", "test@example.com").Return(user, nil).Once()

		token, err := uc.Login("test@example.com", "password123")

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("login with non-existent email", func(t *testing.T) {
		mockRepo.On("GetByEmail", "nonexistent@example.com").Return((*entity.User)(nil), nil).Once()

		token, err := uc.Login("nonexistent@example.com", "password123")

		assert.Error(t, err)
		assert.Equal(t, "email not found", err.Error())
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("login with wrong password", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := &entity.User{
			ID:       1,
			Username: "testuser",
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}
		mockRepo.On("GetByEmail", "test@example.com").Return(user, nil).Once()

		token, err := uc.Login("test@example.com", "wrongpassword")

		assert.Error(t, err)
		assert.Equal(t, "invalid password", err.Error())
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthUsecase_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	uc := NewAuthUsecase(mockRepo, "secret")

	t.Run("successful registration", func(t *testing.T) {
		mockRepo.On("GetByEmail", "new@example.com").Return((*entity.User)(nil), nil).Once()
		mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil).Once()

		err := uc.Register("newuser", "new@example.com", "password123")

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("registration with existing email", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := &entity.User{
			ID:       1,
			Username: "existinguser",
			Email:    "existing@example.com",
			Password: string(hashedPassword),
		}
		mockRepo.On("GetByEmail", "existing@example.com").Return(user, nil).Once()

		err := uc.Register("newuser", "existing@example.com", "password123")

		assert.Error(t, err)
		assert.Equal(t, "email already exists", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("registration failure on create", func(t *testing.T) {
		mockRepo.On("GetByEmail", "fail@example.com").Return((*entity.User)(nil), nil).Once()
		mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(assert.AnError).Once()

		err := uc.Register("failuser", "fail@example.com", "password123")

		assert.Error(t, err)
		assert.Equal(t, assert.AnError.Error(), err.Error())
		mockRepo.AssertExpectations(t)
	})
}
