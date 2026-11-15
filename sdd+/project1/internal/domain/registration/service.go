package registration

import (
	"errors"
	"event-api/internal/domain/event"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventFull         = errors.New("event is at full capacity")
	ErrEventNotPublished = errors.New("event is not published")
	ErrAlreadyRegistered = errors.New("user is already registered for this event")
	ErrNotRegistered     = errors.New("user is not registered for this event")
)

type Service struct {
	repo      Repository
	eventRepo event.Repository
}

func NewService(repo Repository, eventRepo event.Repository) *Service {
	return &Service{repo: repo, eventRepo: eventRepo}
}

func (s *Service) RegisterForEvent(eventID, attendeeID uuid.UUID) (*Registration, error) {
	e, err := s.eventRepo.FindByID(eventID)
	if err != nil {
		return nil, err
	}

	if e.Status != "published" {
		return nil, ErrEventNotPublished
	}

	existingReg, err := s.repo.FindByEventAndAttendee(eventID, attendeeID)
	if err != nil {
		return nil, err
	}
	if existingReg != nil {
		return nil, ErrAlreadyRegistered
	}

	currentRegistrations, err := s.repo.CountByEventID(eventID)
	if err != nil {
		return nil, err
	}

	if currentRegistrations >= e.Capacity {
		return nil, ErrEventFull
	}

	reg := &Registration{
		ID:         uuid.New(),
		EventID:    eventID,
		AttendeeID: attendeeID,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Create(reg); err != nil {
		return nil, err
	}

	return reg, nil
}

func (s *Service) UnregisterFromEvent(eventID, attendeeID uuid.UUID) error {
	existingReg, err := s.repo.FindByEventAndAttendee(eventID, attendeeID)
	if err != nil {
		return err
	}
	if existingReg == nil {
		return ErrNotRegistered
	}

	return s.repo.Delete(eventID, attendeeID)
}
