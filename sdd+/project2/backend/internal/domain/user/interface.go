package user

import (
	"github.com/google/uuid"
)

// Repository defines the interface for user data storage.
type Repository interface {
	Create(u *User) error
	FindByUsername(username string) (*User, error)
	FindAllByRole(role string) ([]*User, error)
	UpdateStatus(id uuid.UUID, status string) error
}
