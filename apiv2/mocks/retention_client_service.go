// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	retention "github.com/TonyZZhang/goharbor-client/v5/apiv2/internal/api/client/retention"
	mock "github.com/stretchr/testify/mock"
)

// MockRetentionClientService is an autogenerated mock type for the ClientService type
type MockRetentionClientService struct {
	mock.Mock
}

// CreateRetention provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) CreateRetention(params *retention.CreateRetentionParams, authInfo runtime.ClientAuthInfoWriter) (*retention.CreateRetentionCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.CreateRetentionCreated
	if rf, ok := ret.Get(0).(func(*retention.CreateRetentionParams, runtime.ClientAuthInfoWriter) *retention.CreateRetentionCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.CreateRetentionCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.CreateRetentionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRetention provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) DeleteRetention(params *retention.DeleteRetentionParams, authInfo runtime.ClientAuthInfoWriter) (*retention.DeleteRetentionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.DeleteRetentionOK
	if rf, ok := ret.Get(0).(func(*retention.DeleteRetentionParams, runtime.ClientAuthInfoWriter) *retention.DeleteRetentionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.DeleteRetentionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.DeleteRetentionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRentenitionMetadata provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) GetRentenitionMetadata(params *retention.GetRentenitionMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*retention.GetRentenitionMetadataOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.GetRentenitionMetadataOK
	if rf, ok := ret.Get(0).(func(*retention.GetRentenitionMetadataParams, runtime.ClientAuthInfoWriter) *retention.GetRentenitionMetadataOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.GetRentenitionMetadataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.GetRentenitionMetadataParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetention provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) GetRetention(params *retention.GetRetentionParams, authInfo runtime.ClientAuthInfoWriter) (*retention.GetRetentionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.GetRetentionOK
	if rf, ok := ret.Get(0).(func(*retention.GetRetentionParams, runtime.ClientAuthInfoWriter) *retention.GetRetentionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.GetRetentionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.GetRetentionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetentionTaskLog provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) GetRetentionTaskLog(params *retention.GetRetentionTaskLogParams, authInfo runtime.ClientAuthInfoWriter) (*retention.GetRetentionTaskLogOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.GetRetentionTaskLogOK
	if rf, ok := ret.Get(0).(func(*retention.GetRetentionTaskLogParams, runtime.ClientAuthInfoWriter) *retention.GetRetentionTaskLogOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.GetRetentionTaskLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.GetRetentionTaskLogParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRetentionExecutions provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) ListRetentionExecutions(params *retention.ListRetentionExecutionsParams, authInfo runtime.ClientAuthInfoWriter) (*retention.ListRetentionExecutionsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.ListRetentionExecutionsOK
	if rf, ok := ret.Get(0).(func(*retention.ListRetentionExecutionsParams, runtime.ClientAuthInfoWriter) *retention.ListRetentionExecutionsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.ListRetentionExecutionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.ListRetentionExecutionsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRetentionTasks provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) ListRetentionTasks(params *retention.ListRetentionTasksParams, authInfo runtime.ClientAuthInfoWriter) (*retention.ListRetentionTasksOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.ListRetentionTasksOK
	if rf, ok := ret.Get(0).(func(*retention.ListRetentionTasksParams, runtime.ClientAuthInfoWriter) *retention.ListRetentionTasksOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.ListRetentionTasksOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.ListRetentionTasksParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OperateRetentionExecution provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) OperateRetentionExecution(params *retention.OperateRetentionExecutionParams, authInfo runtime.ClientAuthInfoWriter) (*retention.OperateRetentionExecutionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.OperateRetentionExecutionOK
	if rf, ok := ret.Get(0).(func(*retention.OperateRetentionExecutionParams, runtime.ClientAuthInfoWriter) *retention.OperateRetentionExecutionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.OperateRetentionExecutionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.OperateRetentionExecutionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockRetentionClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// TriggerRetentionExecution provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) TriggerRetentionExecution(params *retention.TriggerRetentionExecutionParams, authInfo runtime.ClientAuthInfoWriter) (*retention.TriggerRetentionExecutionOK, *retention.TriggerRetentionExecutionCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.TriggerRetentionExecutionOK
	if rf, ok := ret.Get(0).(func(*retention.TriggerRetentionExecutionParams, runtime.ClientAuthInfoWriter) *retention.TriggerRetentionExecutionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.TriggerRetentionExecutionOK)
		}
	}

	var r1 *retention.TriggerRetentionExecutionCreated
	if rf, ok := ret.Get(1).(func(*retention.TriggerRetentionExecutionParams, runtime.ClientAuthInfoWriter) *retention.TriggerRetentionExecutionCreated); ok {
		r1 = rf(params, authInfo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*retention.TriggerRetentionExecutionCreated)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*retention.TriggerRetentionExecutionParams, runtime.ClientAuthInfoWriter) error); ok {
		r2 = rf(params, authInfo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateRetention provides a mock function with given fields: params, authInfo
func (_m *MockRetentionClientService) UpdateRetention(params *retention.UpdateRetentionParams, authInfo runtime.ClientAuthInfoWriter) (*retention.UpdateRetentionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *retention.UpdateRetentionOK
	if rf, ok := ret.Get(0).(func(*retention.UpdateRetentionParams, runtime.ClientAuthInfoWriter) *retention.UpdateRetentionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.UpdateRetentionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*retention.UpdateRetentionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRetentionClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRetentionClientService creates a new instance of MockRetentionClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRetentionClientService(t mockConstructorTestingTNewMockRetentionClientService) *MockRetentionClientService {
	mock := &MockRetentionClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
