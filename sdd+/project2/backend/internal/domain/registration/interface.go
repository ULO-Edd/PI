package registration

import "github.com/google/uuid"

// Repository defines the interface for registration data storage.
type Repository interface {
	Create(reg *Registration) error
	Delete(eventID, attendeeID uuid.UUID) error
	CountByEventID(eventID uuid.UUID) (int, error)
	FindByEventAndAttendee(eventID, attendeeID uuid.UUID) (*Registration, error)
}
