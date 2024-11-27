// Code generated by mockery v2.49.1. DO NOT EDIT.

package io_mock

import mock "github.com/stretchr/testify/mock"

// MockCloser is an autogenerated mock type for the Closer type
type MockCloser struct {
	mock.Mock
}

type MockCloser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCloser) EXPECT() *MockCloser_Expecter {
	return &MockCloser_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockCloser) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCloser_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockCloser_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockCloser_Expecter) Close() *MockCloser_Close_Call {
	return &MockCloser_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockCloser_Close_Call) Run(run func()) *MockCloser_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCloser_Close_Call) Return(_a0 error) *MockCloser_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCloser_Close_Call) RunAndReturn(run func() error) *MockCloser_Close_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCloser creates a new instance of MockCloser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCloser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCloser {
	mock := &MockCloser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
