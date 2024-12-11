// Code generated by mockery v2.50.0. DO NOT EDIT.

package client_mock

import (
	context "context"

	client "github.com/int128/kubelogin/pkg/oidc/client"

	mock "github.com/stretchr/testify/mock"

	oauth2dev "github.com/int128/oauth2dev"

	oidc "github.com/int128/kubelogin/pkg/oidc"
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

// ExchangeAuthCode provides a mock function with given fields: ctx, in
func (_m *MockInterface) ExchangeAuthCode(ctx context.Context, in client.ExchangeAuthCodeInput) (*oidc.TokenSet, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeAuthCode")
	}

	var r0 *oidc.TokenSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ExchangeAuthCodeInput) (*oidc.TokenSet, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.ExchangeAuthCodeInput) *oidc.TokenSet); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.TokenSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.ExchangeAuthCodeInput) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ExchangeAuthCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeAuthCode'
type MockInterface_ExchangeAuthCode_Call struct {
	*mock.Call
}

// ExchangeAuthCode is a helper method to define mock.On call
//   - ctx context.Context
//   - in client.ExchangeAuthCodeInput
func (_e *MockInterface_Expecter) ExchangeAuthCode(ctx interface{}, in interface{}) *MockInterface_ExchangeAuthCode_Call {
	return &MockInterface_ExchangeAuthCode_Call{Call: _e.mock.On("ExchangeAuthCode", ctx, in)}
}

func (_c *MockInterface_ExchangeAuthCode_Call) Run(run func(ctx context.Context, in client.ExchangeAuthCodeInput)) *MockInterface_ExchangeAuthCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.ExchangeAuthCodeInput))
	})
	return _c
}

func (_c *MockInterface_ExchangeAuthCode_Call) Return(_a0 *oidc.TokenSet, _a1 error) *MockInterface_ExchangeAuthCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ExchangeAuthCode_Call) RunAndReturn(run func(context.Context, client.ExchangeAuthCodeInput) (*oidc.TokenSet, error)) *MockInterface_ExchangeAuthCode_Call {
	_c.Call.Return(run)
	return _c
}

// ExchangeDeviceCode provides a mock function with given fields: ctx, authResponse
func (_m *MockInterface) ExchangeDeviceCode(ctx context.Context, authResponse *oauth2dev.AuthorizationResponse) (*oidc.TokenSet, error) {
	ret := _m.Called(ctx, authResponse)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeDeviceCode")
	}

	var r0 *oidc.TokenSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2dev.AuthorizationResponse) (*oidc.TokenSet, error)); ok {
		return rf(ctx, authResponse)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2dev.AuthorizationResponse) *oidc.TokenSet); ok {
		r0 = rf(ctx, authResponse)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.TokenSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oauth2dev.AuthorizationResponse) error); ok {
		r1 = rf(ctx, authResponse)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ExchangeDeviceCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeDeviceCode'
type MockInterface_ExchangeDeviceCode_Call struct {
	*mock.Call
}

// ExchangeDeviceCode is a helper method to define mock.On call
//   - ctx context.Context
//   - authResponse *oauth2dev.AuthorizationResponse
func (_e *MockInterface_Expecter) ExchangeDeviceCode(ctx interface{}, authResponse interface{}) *MockInterface_ExchangeDeviceCode_Call {
	return &MockInterface_ExchangeDeviceCode_Call{Call: _e.mock.On("ExchangeDeviceCode", ctx, authResponse)}
}

func (_c *MockInterface_ExchangeDeviceCode_Call) Run(run func(ctx context.Context, authResponse *oauth2dev.AuthorizationResponse)) *MockInterface_ExchangeDeviceCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*oauth2dev.AuthorizationResponse))
	})
	return _c
}

func (_c *MockInterface_ExchangeDeviceCode_Call) Return(_a0 *oidc.TokenSet, _a1 error) *MockInterface_ExchangeDeviceCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ExchangeDeviceCode_Call) RunAndReturn(run func(context.Context, *oauth2dev.AuthorizationResponse) (*oidc.TokenSet, error)) *MockInterface_ExchangeDeviceCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthCodeURL provides a mock function with given fields: in
func (_m *MockInterface) GetAuthCodeURL(in client.AuthCodeURLInput) string {
	ret := _m.Called(in)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthCodeURL")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(client.AuthCodeURLInput) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockInterface_GetAuthCodeURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthCodeURL'
type MockInterface_GetAuthCodeURL_Call struct {
	*mock.Call
}

// GetAuthCodeURL is a helper method to define mock.On call
//   - in client.AuthCodeURLInput
func (_e *MockInterface_Expecter) GetAuthCodeURL(in interface{}) *MockInterface_GetAuthCodeURL_Call {
	return &MockInterface_GetAuthCodeURL_Call{Call: _e.mock.On("GetAuthCodeURL", in)}
}

func (_c *MockInterface_GetAuthCodeURL_Call) Run(run func(in client.AuthCodeURLInput)) *MockInterface_GetAuthCodeURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.AuthCodeURLInput))
	})
	return _c
}

func (_c *MockInterface_GetAuthCodeURL_Call) Return(_a0 string) *MockInterface_GetAuthCodeURL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_GetAuthCodeURL_Call) RunAndReturn(run func(client.AuthCodeURLInput) string) *MockInterface_GetAuthCodeURL_Call {
	_c.Call.Return(run)
	return _c
}

// GetDeviceAuthorization provides a mock function with given fields: ctx
func (_m *MockInterface) GetDeviceAuthorization(ctx context.Context) (*oauth2dev.AuthorizationResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceAuthorization")
	}

	var r0 *oauth2dev.AuthorizationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*oauth2dev.AuthorizationResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *oauth2dev.AuthorizationResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2dev.AuthorizationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetDeviceAuthorization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDeviceAuthorization'
type MockInterface_GetDeviceAuthorization_Call struct {
	*mock.Call
}

// GetDeviceAuthorization is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInterface_Expecter) GetDeviceAuthorization(ctx interface{}) *MockInterface_GetDeviceAuthorization_Call {
	return &MockInterface_GetDeviceAuthorization_Call{Call: _e.mock.On("GetDeviceAuthorization", ctx)}
}

func (_c *MockInterface_GetDeviceAuthorization_Call) Run(run func(ctx context.Context)) *MockInterface_GetDeviceAuthorization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInterface_GetDeviceAuthorization_Call) Return(_a0 *oauth2dev.AuthorizationResponse, _a1 error) *MockInterface_GetDeviceAuthorization_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetDeviceAuthorization_Call) RunAndReturn(run func(context.Context) (*oauth2dev.AuthorizationResponse, error)) *MockInterface_GetDeviceAuthorization_Call {
	_c.Call.Return(run)
	return _c
}

func (_m *MockInterface) GetAuthCode(ctx context.Context, in client.GetTokenByAuthCodeInput, localServerReadyChan chan<- string) (string, error) {
	return "", nil
}
// GetTokenByAuthCode provides a mock function with given fields: ctx, in, localServerReadyChan
func (_m *MockInterface) GetTokenByAuthCode(ctx context.Context, in client.GetTokenByAuthCodeInput, localServerReadyChan chan<- string) (*oidc.TokenSet, error) {
	ret := _m.Called(ctx, in, localServerReadyChan)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenByAuthCode")
	}

	var r0 *oidc.TokenSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.GetTokenByAuthCodeInput, chan<- string) (*oidc.TokenSet, error)); ok {
		return rf(ctx, in, localServerReadyChan)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.GetTokenByAuthCodeInput, chan<- string) *oidc.TokenSet); ok {
		r0 = rf(ctx, in, localServerReadyChan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.TokenSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.GetTokenByAuthCodeInput, chan<- string) error); ok {
		r1 = rf(ctx, in, localServerReadyChan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetTokenByAuthCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenByAuthCode'
type MockInterface_GetTokenByAuthCode_Call struct {
	*mock.Call
}

// GetTokenByAuthCode is a helper method to define mock.On call
//   - ctx context.Context
//   - in client.GetTokenByAuthCodeInput
//   - localServerReadyChan chan<- string
func (_e *MockInterface_Expecter) GetTokenByAuthCode(ctx interface{}, in interface{}, localServerReadyChan interface{}) *MockInterface_GetTokenByAuthCode_Call {
	return &MockInterface_GetTokenByAuthCode_Call{Call: _e.mock.On("GetTokenByAuthCode", ctx, in, localServerReadyChan)}
}

func (_c *MockInterface_GetTokenByAuthCode_Call) Run(run func(ctx context.Context, in client.GetTokenByAuthCodeInput, localServerReadyChan chan<- string)) *MockInterface_GetTokenByAuthCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(client.GetTokenByAuthCodeInput), args[2].(chan<- string))
	})
	return _c
}

func (_c *MockInterface_GetTokenByAuthCode_Call) Return(_a0 *oidc.TokenSet, _a1 error) *MockInterface_GetTokenByAuthCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetTokenByAuthCode_Call) RunAndReturn(run func(context.Context, client.GetTokenByAuthCodeInput, chan<- string) (*oidc.TokenSet, error)) *MockInterface_GetTokenByAuthCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokenByROPC provides a mock function with given fields: ctx, username, password
func (_m *MockInterface) GetTokenByROPC(ctx context.Context, username string, password string) (*oidc.TokenSet, error) {
	ret := _m.Called(ctx, username, password)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenByROPC")
	}

	var r0 *oidc.TokenSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*oidc.TokenSet, error)); ok {
		return rf(ctx, username, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *oidc.TokenSet); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.TokenSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetTokenByROPC_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenByROPC'
type MockInterface_GetTokenByROPC_Call struct {
	*mock.Call
}

// GetTokenByROPC is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - password string
func (_e *MockInterface_Expecter) GetTokenByROPC(ctx interface{}, username interface{}, password interface{}) *MockInterface_GetTokenByROPC_Call {
	return &MockInterface_GetTokenByROPC_Call{Call: _e.mock.On("GetTokenByROPC", ctx, username, password)}
}

func (_c *MockInterface_GetTokenByROPC_Call) Run(run func(ctx context.Context, username string, password string)) *MockInterface_GetTokenByROPC_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockInterface_GetTokenByROPC_Call) Return(_a0 *oidc.TokenSet, _a1 error) *MockInterface_GetTokenByROPC_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetTokenByROPC_Call) RunAndReturn(run func(context.Context, string, string) (*oidc.TokenSet, error)) *MockInterface_GetTokenByROPC_Call {
	_c.Call.Return(run)
	return _c
}

// Refresh provides a mock function with given fields: ctx, refreshToken
func (_m *MockInterface) Refresh(ctx context.Context, refreshToken string) (*oidc.TokenSet, error) {
	ret := _m.Called(ctx, refreshToken)

	if len(ret) == 0 {
		panic("no return value specified for Refresh")
	}

	var r0 *oidc.TokenSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*oidc.TokenSet, error)); ok {
		return rf(ctx, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *oidc.TokenSet); ok {
		r0 = rf(ctx, refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.TokenSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Refresh_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Refresh'
type MockInterface_Refresh_Call struct {
	*mock.Call
}

// Refresh is a helper method to define mock.On call
//   - ctx context.Context
//   - refreshToken string
func (_e *MockInterface_Expecter) Refresh(ctx interface{}, refreshToken interface{}) *MockInterface_Refresh_Call {
	return &MockInterface_Refresh_Call{Call: _e.mock.On("Refresh", ctx, refreshToken)}
}

func (_c *MockInterface_Refresh_Call) Run(run func(ctx context.Context, refreshToken string)) *MockInterface_Refresh_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockInterface_Refresh_Call) Return(_a0 *oidc.TokenSet, _a1 error) *MockInterface_Refresh_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Refresh_Call) RunAndReturn(run func(context.Context, string) (*oidc.TokenSet, error)) *MockInterface_Refresh_Call {
	_c.Call.Return(run)
	return _c
}

// SupportedPKCEMethods provides a mock function with no fields
func (_m *MockInterface) SupportedPKCEMethods() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SupportedPKCEMethods")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockInterface_SupportedPKCEMethods_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SupportedPKCEMethods'
type MockInterface_SupportedPKCEMethods_Call struct {
	*mock.Call
}

// SupportedPKCEMethods is a helper method to define mock.On call
func (_e *MockInterface_Expecter) SupportedPKCEMethods() *MockInterface_SupportedPKCEMethods_Call {
	return &MockInterface_SupportedPKCEMethods_Call{Call: _e.mock.On("SupportedPKCEMethods")}
}

func (_c *MockInterface_SupportedPKCEMethods_Call) Run(run func()) *MockInterface_SupportedPKCEMethods_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_SupportedPKCEMethods_Call) Return(_a0 []string) *MockInterface_SupportedPKCEMethods_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_SupportedPKCEMethods_Call) RunAndReturn(run func() []string) *MockInterface_SupportedPKCEMethods_Call {
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
