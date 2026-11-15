package registration_test

import (
	"event-api/internal/domain/event"
	eventMocks "event-api/internal/domain/event/mocks"
	"event-api/internal/domain/registration"
	regMocks "event-api/internal/domain/registration/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_RegisterForEvent(t *testing.T) {
	mockRepo := new(regMocks.MockRepository)
	mockEventRepo := new(eventMocks.MockRepository)
	service := registration.NewService(mockRepo, mockEventRepo)

	eventID := uuid.New()
	attendeeID := uuid.New()
	publishedEvent := &event.Event{
		ID:       eventID,
		Status:   "published",
		Capacity: 50,
	}

	draftEvent := &event.Event{
		ID:       eventID,
		Status:   "draft",
		Capacity: 50,
	}

	t.Run("successful registration", func(t *testing.T) {
		mockEventRepo.On("FindByID", eventID).Return(publishedEvent, nil).Once()
		mockRepo.On("FindByEventAndAttendee", eventID, attendeeID).Return(nil, nil).Once()
		mockRepo.On("CountByEventID", eventID).Return(49, nil).Once()
		mockRepo.On("Create", mock.AnythingOfType("*registration.Registration")).Return(nil).Once()

		reg, err := service.RegisterForEvent(eventID, attendeeID)

		assert.NoError(t, err)
		assert.NotNil(t, reg)
	})

	t.Run("event is full", func(t *testing.T) {
		mockEventRepo.On("FindByID", eventID).Return(publishedEvent, nil).Once()
		mockRepo.On("FindByEventAndAttendee", eventID, attendeeID).Return(nil, nil).Once()
		mockRepo.On("CountByEventID", eventID).Return(50, nil).Once()

		_, err := service.RegisterForEvent(eventID, attendeeID)

		assert.Error(t, err)
		assert.Equal(t, registration.ErrEventFull, err)
	})

	t.Run("event not published", func(t *testing.T) {
		mockEventRepo.On("FindByID", eventID).Return(draftEvent, nil).Once()

		_, err := service.RegisterForEvent(eventID, attendeeID)

		assert.Error(t, err)
		assert.Equal(t, registration.ErrEventNotPublished, err)
	})

	t.Run("already registered", func(t *testing.T) {
		mockEventRepo.On("FindByID", eventID).Return(publishedEvent, nil).Once()
		mockRepo.On("FindByEventAndAttendee", eventID, attendeeID).Return(&registration.Registration{}, nil).Once()

		_, err := service.RegisterForEvent(eventID, attendeeID)

		assert.Error(t, err)
		assert.Equal(t, registration.ErrAlreadyRegistered, err)
	})
}

func TestService_UnregisterFromEvent(t *testing.T) {
	mockRepo := new(regMocks.MockRepository)
	mockEventRepo := new(eventMocks.MockRepository)
	service := registration.NewService(mockRepo, mockEventRepo)

	eventID := uuid.New()
	attendeeID := uuid.New()

	t.Run("successful unregistration", func(t *testing.T) {
		mockRepo.On("FindByEventAndAttendee", eventID, attendeeID).Return(&registration.Registration{}, nil).Once()
		mockRepo.On("Delete", eventID, attendeeID).Return(nil).Once()

		err := service.UnregisterFromEvent(eventID, attendeeID)
		assert.NoError(t, err)
	})

	t.Run("not registered", func(t *testing.T) {
		mockRepo.On("FindByEventAndAttendee", eventID, attendeeID).Return(nil, nil).Once()

		err := service.UnregisterFromEvent(eventID, attendeeID)
		assert.Error(t, err)
		assert.Equal(t, registration.ErrNotRegistered, err)
	})
}
