// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	icon "github.com/TonyZZhang/goharbor-client/v5/apiv2/internal/api/client/icon"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockIconClientService is an autogenerated mock type for the ClientService type
type MockIconClientService struct {
	mock.Mock
}

// GetIcon provides a mock function with given fields: params, authInfo
func (_m *MockIconClientService) GetIcon(params *icon.GetIconParams, authInfo runtime.ClientAuthInfoWriter) (*icon.GetIconOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *icon.GetIconOK
	if rf, ok := ret.Get(0).(func(*icon.GetIconParams, runtime.ClientAuthInfoWriter) *icon.GetIconOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*icon.GetIconOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*icon.GetIconParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockIconClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockIconClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIconClientService creates a new instance of MockIconClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIconClientService(t mockConstructorTestingTNewMockIconClientService) *MockIconClientService {
	mock := &MockIconClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
