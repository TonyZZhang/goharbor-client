// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	health "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/health"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockHealthClientService is an autogenerated mock type for the ClientService type
type MockHealthClientService struct {
	mock.Mock
}

// GetHealth provides a mock function with given fields: params, authInfo
func (_m *MockHealthClientService) GetHealth(params *health.GetHealthParams, authInfo runtime.ClientAuthInfoWriter) (*health.GetHealthOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *health.GetHealthOK
	if rf, ok := ret.Get(0).(func(*health.GetHealthParams, runtime.ClientAuthInfoWriter) *health.GetHealthOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*health.GetHealthOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*health.GetHealthParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockHealthClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
