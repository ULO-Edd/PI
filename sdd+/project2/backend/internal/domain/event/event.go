package event

import (
	"github.com/google/uuid"
	"time"
)

// Event represents an event in the system.
type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Location    string    `json:"location,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	OrganiserID uuid.UUID `json:"organiser_id"`
	CategoryID  uuid.UUID `json:"category_id,omitempty"`
	Capacity    int       `json:"capacity"`
	Status      string    `json:"status"`
	ContactEmail string    `json:"contact_email,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
