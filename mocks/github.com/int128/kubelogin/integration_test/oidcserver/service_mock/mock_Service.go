// Code generated by mockery v2.49.1. DO NOT EDIT.

package service_mock

import (
	service "github.com/int128/kubelogin/integration_test/oidcserver/service"
	testconfig "github.com/int128/kubelogin/integration_test/oidcserver/testconfig"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// AuthenticateCode provides a mock function with given fields: req
func (_m *MockService) AuthenticateCode(req service.AuthenticationRequest) (string, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for AuthenticateCode")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(service.AuthenticationRequest) (string, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(service.AuthenticationRequest) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(service.AuthenticationRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_AuthenticateCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthenticateCode'
type MockService_AuthenticateCode_Call struct {
	*mock.Call
}

// AuthenticateCode is a helper method to define mock.On call
//   - req service.AuthenticationRequest
func (_e *MockService_Expecter) AuthenticateCode(req interface{}) *MockService_AuthenticateCode_Call {
	return &MockService_AuthenticateCode_Call{Call: _e.mock.On("AuthenticateCode", req)}
}

func (_c *MockService_AuthenticateCode_Call) Run(run func(req service.AuthenticationRequest)) *MockService_AuthenticateCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(service.AuthenticationRequest))
	})
	return _c
}

func (_c *MockService_AuthenticateCode_Call) Return(code string, err error) *MockService_AuthenticateCode_Call {
	_c.Call.Return(code, err)
	return _c
}

func (_c *MockService_AuthenticateCode_Call) RunAndReturn(run func(service.AuthenticationRequest) (string, error)) *MockService_AuthenticateCode_Call {
	_c.Call.Return(run)
	return _c
}

// AuthenticatePassword provides a mock function with given fields: username, password, scope
func (_m *MockService) AuthenticatePassword(username string, password string, scope string) (*service.TokenResponse, error) {
	ret := _m.Called(username, password, scope)

	if len(ret) == 0 {
		panic("no return value specified for AuthenticatePassword")
	}

	var r0 *service.TokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*service.TokenResponse, error)); ok {
		return rf(username, password, scope)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *service.TokenResponse); ok {
		r0 = rf(username, password, scope)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(username, password, scope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_AuthenticatePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthenticatePassword'
type MockService_AuthenticatePassword_Call struct {
	*mock.Call
}

// AuthenticatePassword is a helper method to define mock.On call
//   - username string
//   - password string
//   - scope string
func (_e *MockService_Expecter) AuthenticatePassword(username interface{}, password interface{}, scope interface{}) *MockService_AuthenticatePassword_Call {
	return &MockService_AuthenticatePassword_Call{Call: _e.mock.On("AuthenticatePassword", username, password, scope)}
}

func (_c *MockService_AuthenticatePassword_Call) Run(run func(username string, password string, scope string)) *MockService_AuthenticatePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockService_AuthenticatePassword_Call) Return(_a0 *service.TokenResponse, _a1 error) *MockService_AuthenticatePassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_AuthenticatePassword_Call) RunAndReturn(run func(string, string, string) (*service.TokenResponse, error)) *MockService_AuthenticatePassword_Call {
	_c.Call.Return(run)
	return _c
}

// Discovery provides a mock function with given fields:
func (_m *MockService) Discovery() *service.DiscoveryResponse {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Discovery")
	}

	var r0 *service.DiscoveryResponse
	if rf, ok := ret.Get(0).(func() *service.DiscoveryResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.DiscoveryResponse)
		}
	}

	return r0
}

// MockService_Discovery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Discovery'
type MockService_Discovery_Call struct {
	*mock.Call
}

// Discovery is a helper method to define mock.On call
func (_e *MockService_Expecter) Discovery() *MockService_Discovery_Call {
	return &MockService_Discovery_Call{Call: _e.mock.On("Discovery")}
}

func (_c *MockService_Discovery_Call) Run(run func()) *MockService_Discovery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Discovery_Call) Return(_a0 *service.DiscoveryResponse) *MockService_Discovery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_Discovery_Call) RunAndReturn(run func() *service.DiscoveryResponse) *MockService_Discovery_Call {
	_c.Call.Return(run)
	return _c
}

// Exchange provides a mock function with given fields: req
func (_m *MockService) Exchange(req service.TokenRequest) (*service.TokenResponse, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Exchange")
	}

	var r0 *service.TokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(service.TokenRequest) (*service.TokenResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(service.TokenRequest) *service.TokenResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(service.TokenRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_Exchange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exchange'
type MockService_Exchange_Call struct {
	*mock.Call
}

// Exchange is a helper method to define mock.On call
//   - req service.TokenRequest
func (_e *MockService_Expecter) Exchange(req interface{}) *MockService_Exchange_Call {
	return &MockService_Exchange_Call{Call: _e.mock.On("Exchange", req)}
}

func (_c *MockService_Exchange_Call) Run(run func(req service.TokenRequest)) *MockService_Exchange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(service.TokenRequest))
	})
	return _c
}

func (_c *MockService_Exchange_Call) Return(_a0 *service.TokenResponse, _a1 error) *MockService_Exchange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_Exchange_Call) RunAndReturn(run func(service.TokenRequest) (*service.TokenResponse, error)) *MockService_Exchange_Call {
	_c.Call.Return(run)
	return _c
}

// GetCertificates provides a mock function with given fields:
func (_m *MockService) GetCertificates() *service.CertificatesResponse {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCertificates")
	}

	var r0 *service.CertificatesResponse
	if rf, ok := ret.Get(0).(func() *service.CertificatesResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CertificatesResponse)
		}
	}

	return r0
}

// MockService_GetCertificates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCertificates'
type MockService_GetCertificates_Call struct {
	*mock.Call
}

// GetCertificates is a helper method to define mock.On call
func (_e *MockService_Expecter) GetCertificates() *MockService_GetCertificates_Call {
	return &MockService_GetCertificates_Call{Call: _e.mock.On("GetCertificates")}
}

func (_c *MockService_GetCertificates_Call) Run(run func()) *MockService_GetCertificates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_GetCertificates_Call) Return(_a0 *service.CertificatesResponse) *MockService_GetCertificates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_GetCertificates_Call) RunAndReturn(run func() *service.CertificatesResponse) *MockService_GetCertificates_Call {
	_c.Call.Return(run)
	return _c
}

// IssuerURL provides a mock function with given fields:
func (_m *MockService) IssuerURL() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IssuerURL")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_IssuerURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IssuerURL'
type MockService_IssuerURL_Call struct {
	*mock.Call
}

// IssuerURL is a helper method to define mock.On call
func (_e *MockService_Expecter) IssuerURL() *MockService_IssuerURL_Call {
	return &MockService_IssuerURL_Call{Call: _e.mock.On("IssuerURL")}
}

func (_c *MockService_IssuerURL_Call) Run(run func()) *MockService_IssuerURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_IssuerURL_Call) Return(_a0 string) *MockService_IssuerURL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_IssuerURL_Call) RunAndReturn(run func() string) *MockService_IssuerURL_Call {
	_c.Call.Return(run)
	return _c
}

// LastTokenResponse provides a mock function with given fields:
func (_m *MockService) LastTokenResponse() *service.TokenResponse {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LastTokenResponse")
	}

	var r0 *service.TokenResponse
	if rf, ok := ret.Get(0).(func() *service.TokenResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TokenResponse)
		}
	}

	return r0
}

// MockService_LastTokenResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LastTokenResponse'
type MockService_LastTokenResponse_Call struct {
	*mock.Call
}

// LastTokenResponse is a helper method to define mock.On call
func (_e *MockService_Expecter) LastTokenResponse() *MockService_LastTokenResponse_Call {
	return &MockService_LastTokenResponse_Call{Call: _e.mock.On("LastTokenResponse")}
}

func (_c *MockService_LastTokenResponse_Call) Run(run func()) *MockService_LastTokenResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_LastTokenResponse_Call) Return(_a0 *service.TokenResponse) *MockService_LastTokenResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_LastTokenResponse_Call) RunAndReturn(run func() *service.TokenResponse) *MockService_LastTokenResponse_Call {
	_c.Call.Return(run)
	return _c
}

// Refresh provides a mock function with given fields: refreshToken
func (_m *MockService) Refresh(refreshToken string) (*service.TokenResponse, error) {
	ret := _m.Called(refreshToken)

	if len(ret) == 0 {
		panic("no return value specified for Refresh")
	}

	var r0 *service.TokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*service.TokenResponse, error)); ok {
		return rf(refreshToken)
	}
	if rf, ok := ret.Get(0).(func(string) *service.TokenResponse); ok {
		r0 = rf(refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_Refresh_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Refresh'
type MockService_Refresh_Call struct {
	*mock.Call
}

// Refresh is a helper method to define mock.On call
//   - refreshToken string
func (_e *MockService_Expecter) Refresh(refreshToken interface{}) *MockService_Refresh_Call {
	return &MockService_Refresh_Call{Call: _e.mock.On("Refresh", refreshToken)}
}

func (_c *MockService_Refresh_Call) Run(run func(refreshToken string)) *MockService_Refresh_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_Refresh_Call) Return(_a0 *service.TokenResponse, _a1 error) *MockService_Refresh_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_Refresh_Call) RunAndReturn(run func(string) (*service.TokenResponse, error)) *MockService_Refresh_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *MockService) SetConfig(config testconfig.TestConfig) {
	_m.Called(config)
}

// MockService_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type MockService_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - config testconfig.TestConfig
func (_e *MockService_Expecter) SetConfig(config interface{}) *MockService_SetConfig_Call {
	return &MockService_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *MockService_SetConfig_Call) Run(run func(config testconfig.TestConfig)) *MockService_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(testconfig.TestConfig))
	})
	return _c
}

func (_c *MockService_SetConfig_Call) Return() *MockService_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockService_SetConfig_Call) RunAndReturn(run func(testconfig.TestConfig)) *MockService_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
