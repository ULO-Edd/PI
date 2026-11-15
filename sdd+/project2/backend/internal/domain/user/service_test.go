package user_test

import (
	"errors"
	"event-api/internal/domain/user"
	"event-api/internal/domain/user/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestService_Register(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := user.NewService(mockRepo)

	t.Run("successful registration", func(t *testing.T) {
		username := "testuser"
		password := "password123"
		role := "attendee"

		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(nil).Once()

		u, err := service.Register(username, password, role)

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, username, u.Username)
		assert.Equal(t, role, u.Role)
		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)))
	})

	t.Run("username too short", func(t *testing.T) {
		_, err := service.Register("us", "password123", "attendee")
		assert.Error(t, err)
		assert.Equal(t, user.ErrInvalidUsername, err)
	})

	t.Run("username non-alphanumeric", func(t *testing.T) {
		_, err := service.Register("user_123", "password123", "attendee")
		assert.Error(t, err)
		assert.Equal(t, user.ErrInvalidUsername, err)
	})

	t.Run("password too short", func(t *testing.T) {
		_, err := service.Register("user123", "pass", "attendee")
		assert.Error(t, err)
		assert.Equal(t, user.ErrInvalidPassword, err)
	})

	t.Run("invalid role", func(t *testing.T) {
		_, err := service.Register("user123", "password123", "unknown")
		assert.Error(t, err)
		assert.Equal(t, user.ErrInvalidRole, err)
	})

	t.Run("normalization of roles", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(nil).Once()

		u, err := service.Register("user123", "password123", "Attendees")
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, "attendee", u.Role)
	})

	t.Run("username taken", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(errors.New("pq: duplicate key value violates unique constraint \"users_username_key\"")).Once()

		_, err := service.Register("takenuser", "password123", "attendee")
		assert.Error(t, err)
		assert.Equal(t, user.ErrUsernameTaken, err)
	})
}

func TestService_Login(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := user.NewService(mockRepo)

	username := "testuser"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockUser := &user.User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         "attendee",
		Status:       "Approved",
		CreatedAt:    time.Now(),
	}

	t.Run("successful login", func(t *testing.T) {
		mockRepo.On("FindByUsername", username).Return(mockUser, nil).Once()

		u, err := service.Login(username, password)

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, username, u.Username)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("FindByUsername", "unknownuser").Return(nil, errors.New("not found")).Once()

		_, err := service.Login("unknownuser", password)
		assert.Error(t, err)
		assert.Equal(t, user.ErrUserNotFound, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		mockRepo.On("FindByUsername", username).Return(mockUser, nil).Once()

		_, err := service.Login(username, "wrongpassword")
		assert.Error(t, err)
		assert.Equal(t, user.ErrInvalidCredentials, err)
	})
}
