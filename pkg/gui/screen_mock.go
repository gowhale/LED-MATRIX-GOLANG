// Code generated by mockery v1.0.0. DO NOT EDIT.

package gui

import mock "github.com/stretchr/testify/mock"

// MockScreen is an autogenerated mock type for the Screen type
type MockScreen struct {
	mock.Mock
}

// AllVapesOff provides a mock function with given fields:
func (_m *MockScreen) AllVapesOff() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisplayMatrix provides a mock function with given fields: _a0
func (_m *MockScreen) DisplayMatrix(_a0 [][]int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([][]int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRow provides a mock function with given fields:
func (_m *MockScreen) NewRow() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShowAndRun provides a mock function with given fields:
func (_m *MockScreen) ShowAndRun() {
	_m.Called()
}

// VapeLightOff provides a mock function with given fields: _a0
func (_m *MockScreen) VapeLightOff(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VapeLightOn provides a mock function with given fields: _a0
func (_m *MockScreen) VapeLightOn(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}