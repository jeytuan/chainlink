// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Clock is an autogenerated mock type for the Clock type
type Clock struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *Clock) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

type mockConstructorTestingTNewClock interface {
	mock.TestingT
	Cleanup(func())
}

// NewClock creates a new instance of Clock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClock(t mockConstructorTestingTNewClock) *Clock {
	mock := &Clock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
