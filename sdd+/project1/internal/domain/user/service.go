package user

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUsernameTaken      = errors.New("username is already taken")
	ErrInvalidUsername     = errors.New("username must be at least 3 characters and alphanumeric")
	ErrInvalidPassword     = errors.New("password must be at least 6 characters")
	ErrInvalidRole         = errors.New("role must be attendee, organiser, or administrator")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

func (s *Service) Register(username, password, role string) (*User, error) {
	username = strings.TrimSpace(username)
	if len(username) < 3 || !isAlphanumeric(username) {
		return nil, ErrInvalidUsername
	}
	if len(password) < 6 {
		return nil, ErrInvalidPassword
	}
	role = strings.ToLower(strings.TrimSpace(role))
	if role == "attendees" {
		role = "attendee"
	} else if role == "organisers" {
		role = "organiser"
	} else if role == "administrators" || role == "admin" {
		role = "administrator"
	}

	if role != "attendee" && role != "organiser" && role != "administrator" {
		return nil, ErrInvalidRole
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		CreatedAt:    time.Now(),
	}

	if err := s.repo.Create(u); err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "23505") {
			return nil, ErrUsernameTaken
		}
		return nil, err
	}

	return u, nil
}

func (s *Service) Login(username, password string) (*User, error) {
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}
