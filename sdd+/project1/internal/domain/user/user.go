package user

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // Do not expose password hash
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
