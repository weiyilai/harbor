// Code generated by mockery v2.53.3. DO NOT EDIT.

package export

import (
	context "context"

	export "github.com/goharbor/harbor/src/pkg/scan/export"
	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, params
func (_m *Manager) Fetch(ctx context.Context, params export.Params) ([]export.Data, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []export.Data
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, export.Params) ([]export.Data, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, export.Params) []export.Data); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]export.Data)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, export.Params) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
