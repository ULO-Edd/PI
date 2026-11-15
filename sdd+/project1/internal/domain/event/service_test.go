package event_test

import (
	"event-api/internal/domain/event"
	"event-api/internal/domain/event/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateEvent(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := event.NewService(mockRepo)

	title := "Test Event"
	description := "A great event"
	location := "Online"
	startTime := time.Now().Add(24 * time.Hour)
	endTime := time.Now().Add(26 * time.Hour)
	organiserID := uuid.New()
	capacity := 100

	t.Run("successful event creation", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*event.Event")).Return(nil).Once()

		e, err := service.CreateEvent(title, description, location, startTime, endTime, organiserID, capacity)

		assert.NoError(t, err)
		assert.NotNil(t, e)
		assert.Equal(t, title, e.Title)
		assert.Equal(t, "draft", e.Status)
		assert.Equal(t, capacity, e.Capacity)
	})

	t.Run("empty title", func(t *testing.T) {
		_, err := service.CreateEvent("", description, location, startTime, endTime, organiserID, capacity)
		assert.Error(t, err)
		assert.Equal(t, event.ErrInvalidTitle, err)
	})

	t.Run("start time in the past", func(t *testing.T) {
		pastStart := time.Now().Add(-1 * time.Hour)
		_, err := service.CreateEvent(title, description, location, pastStart, endTime, organiserID, capacity)
		assert.Error(t, err)
		assert.Equal(t, event.ErrStartTimePast, err)
	})

	t.Run("end time before start time", func(t *testing.T) {
		invalidEnd := startTime.Add(-1 * time.Hour)
		_, err := service.CreateEvent(title, description, location, startTime, invalidEnd, organiserID, capacity)
		assert.Error(t, err)
		assert.Equal(t, event.ErrInvalidTimes, err)
	})

	t.Run("negative capacity", func(t *testing.T) {
		_, err := service.CreateEvent(title, description, location, startTime, endTime, organiserID, -5)
		assert.Error(t, err)
		assert.Equal(t, event.ErrInvalidCapacity, err)
	})
}

func TestService_UpdateEvent(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := event.NewService(mockRepo)

	t.Run("successful update", func(t *testing.T) {
		e := &event.Event{
			ID:          uuid.New(),
			Title:       "Updated Title",
			Description: "Some desc",
			Location:    "Onsite",
			StartTime:   time.Now().Add(2 * time.Hour),
			EndTime:     time.Now().Add(4 * time.Hour),
			OrganiserID: uuid.New(),
			Capacity:    50,
			Status:      "Published",
		}

		mockRepo.On("Update", e).Return(nil).Once()

		err := service.UpdateEvent(e)
		assert.NoError(t, err)
		assert.Equal(t, "published", e.Status) // normalized to lowercase
	})

	t.Run("invalid capacity", func(t *testing.T) {
		e := &event.Event{
			ID:        uuid.New(),
			Title:     "Title",
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(2 * time.Hour),
			Capacity:  0,
		}
		err := service.UpdateEvent(e)
		assert.Error(t, err)
		assert.Equal(t, event.ErrInvalidCapacity, err)
	})

	t.Run("invalid status", func(t *testing.T) {
		e := &event.Event{
			ID:        uuid.New(),
			Title:     "Title",
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(2 * time.Hour),
			Capacity:  10,
			Status:    "unknown",
		}
		err := service.UpdateEvent(e)
		assert.Error(t, err)
		assert.Equal(t, event.ErrInvalidStatus, err)
	})
}

func TestService_GetAllEvents(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := event.NewService(mockRepo)

	mockRepo.On("FindAll", "published", "online", "").Return([]*event.Event{}, nil).Once()

	events, err := service.GetAllEvents("published", "online", "")
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

func TestService_GetEventByID(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := event.NewService(mockRepo)

	eventID := uuid.New()
	mockRepo.On("FindByID", eventID).Return(&event.Event{ID: eventID}, nil).Once()

	e, err := service.GetEventByID(eventID)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, eventID, e.ID)
}

func TestService_DeleteEvent(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := event.NewService(mockRepo)

	eventID := uuid.New()
	mockRepo.On("Delete", eventID).Return(nil).Once()

	err := service.DeleteEvent(eventID)
	assert.NoError(t, err)
}
