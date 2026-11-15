package registration

import (
	"github.com/google/uuid"
	"time"
)

// Registration represents a user's registration for an event.
type Registration struct {
	ID         uuid.UUID `json:"id"`
	EventID    uuid.UUID `json:"event_id"`
	AttendeeID uuid.UUID `json:"attendee_id"`
	CreatedAt  time.Time `json:"created_at"`
}
