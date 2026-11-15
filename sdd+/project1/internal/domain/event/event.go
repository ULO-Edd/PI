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
	Capacity    int       `json:"capacity"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
