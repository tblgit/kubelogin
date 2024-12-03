// Code generated by mockery v2.49.2. DO NOT EDIT.

package writer_mock

import (
	credentialplugin "github.com/int128/kubelogin/pkg/credentialplugin"
	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// Write provides a mock function with given fields: out
func (_m *MockInterface) Write(out credentialplugin.Output) error {
	ret := _m.Called(out)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(credentialplugin.Output) error); ok {
		r0 = rf(out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockInterface_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - out credentialplugin.Output
func (_e *MockInterface_Expecter) Write(out interface{}) *MockInterface_Write_Call {
	return &MockInterface_Write_Call{Call: _e.mock.On("Write", out)}
}

func (_c *MockInterface_Write_Call) Run(run func(out credentialplugin.Output)) *MockInterface_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(credentialplugin.Output))
	})
	return _c
}

func (_c *MockInterface_Write_Call) Return(_a0 error) *MockInterface_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Write_Call) RunAndReturn(run func(credentialplugin.Output) error) *MockInterface_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
