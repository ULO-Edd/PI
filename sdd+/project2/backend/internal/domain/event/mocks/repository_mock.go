package mocks

import (
	"event-api/internal/domain/event"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (_m *MockRepository) Create(e *event.Event) error {
	ret := _m.Called(e)
	return ret.Error(0)
}

func (_m *MockRepository) FindAll(status, location, organiserID string) ([]*event.Event, error) {
	ret := _m.Called(status, location, organiserID)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]*event.Event), ret.Error(1)
}

func (_m *MockRepository) FindByID(id uuid.UUID) (*event.Event, error) {
	ret := _m.Called(id)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*event.Event), ret.Error(1)
}

func (_m *MockRepository) Update(e *event.Event) error {
	ret := _m.Called(e)
	return ret.Error(0)
}

func (_m *MockRepository) Delete(id uuid.UUID) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
