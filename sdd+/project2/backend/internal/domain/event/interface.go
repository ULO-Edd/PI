package event

import "github.com/google/uuid"

// Repository defines the interface for event data storage.
type Repository interface {
	Create(e *Event) error
	FindAll(status, location, organiserID string) ([]*Event, error)
	FindByID(id uuid.UUID) (*Event, error)
	Update(e *Event) error
	Delete(id uuid.UUID) error
}
