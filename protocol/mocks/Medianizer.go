// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Medianizer is an autogenerated mock type for the Medianizer type
type Medianizer struct {
	mock.Mock
}

// MedianUint64 provides a mock function with given fields: input
func (_m *Medianizer) MedianUint64(input []uint64) (uint64, error) {
	ret := _m.Called(input)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func([]uint64) (uint64, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func([]uint64) uint64); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func([]uint64) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMedianizer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMedianizer creates a new instance of Medianizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMedianizer(t mockConstructorTestingTNewMedianizer) *Medianizer {
	mock := &Medianizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
