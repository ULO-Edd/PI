package mocks

import (
	"event-api/internal/domain/registration"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (_m *MockRepository) Create(reg *registration.Registration) error {
	ret := _m.Called(reg)
	return ret.Error(0)
}

func (_m *MockRepository) Delete(eventID, attendeeID uuid.UUID) error {
	ret := _m.Called(eventID, attendeeID)
	return ret.Error(0)
}

func (_m *MockRepository) CountByEventID(eventID uuid.UUID) (int, error) {
	ret := _m.Called(eventID)
	return ret.Int(0), ret.Error(1)
}

func (_m *MockRepository) FindByEventAndAttendee(eventID, attendeeID uuid.UUID) (*registration.Registration, error) {
	ret := _m.Called(eventID, attendeeID)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*registration.Registration), ret.Error(1)
}

