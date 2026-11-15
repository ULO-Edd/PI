package event

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidTitle    = errors.New("event title must be non-empty and less than 255 characters")
	ErrInvalidTimes    = errors.New("event end time must be strictly after start time")
	ErrStartTimePast   = errors.New("event start time must be in the future")
	ErrInvalidCapacity = errors.New("event capacity must be greater than zero")
	ErrInvalidStatus   = errors.New("event status must be draft, published, or cancelled")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateEvent(title, description, location string, startTime, endTime time.Time, organiserID uuid.UUID, capacity int) (*Event, error) {
	title = strings.TrimSpace(title)
	if title == "" || len(title) > 255 {
		return nil, ErrInvalidTitle
	}
	if startTime.Before(time.Now()) {
		return nil, ErrStartTimePast
	}
	if !endTime.After(startTime) {
		return nil, ErrInvalidTimes
	}
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	e := &Event{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Location:    location,
		StartTime:   startTime,
		EndTime:     endTime,
		OrganiserID: organiserID,
		Capacity:    capacity,
		Status:      "draft",
		CreatedAt:   time.Now(),
	}
	if err := s.repo.Create(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *Service) GetAllEvents(status, location, organiserID string) ([]*Event, error) {
	return s.repo.FindAll(status, location, organiserID)
}

func (s *Service) GetEventByID(id uuid.UUID) (*Event, error) {
	return s.repo.FindByID(id)
}

func (s *Service) UpdateEvent(e *Event) error {
	e.Title = strings.TrimSpace(e.Title)
	if e.Title == "" || len(e.Title) > 255 {
		return ErrInvalidTitle
	}
	if !e.EndTime.After(e.StartTime) {
		return ErrInvalidTimes
	}
	if e.Capacity <= 0 {
		return ErrInvalidCapacity
	}
	status := strings.ToLower(e.Status)
	if status != "draft" && status != "published" && status != "cancelled" {
		return ErrInvalidStatus
	}
	e.Status = status
	return s.repo.Update(e)
}

func (s *Service) DeleteEvent(id uuid.UUID) error {
	return s.repo.Delete(id)
}
