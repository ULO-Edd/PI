package mocks

import (
	"event-api/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: u
func (_m *MockRepository) Create(u *user.User) error {
	ret := _m.Called(u)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user.User) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByUsername provides a mock function with given fields: username
func (_m *MockRepository) FindByUsername(username string) (*user.User, error) {
	ret := _m.Called(username)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
