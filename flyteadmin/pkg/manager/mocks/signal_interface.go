// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	mock "github.com/stretchr/testify/mock"
)

// SignalInterface is an autogenerated mock type for the SignalInterface type
type SignalInterface struct {
	mock.Mock
}

type SignalInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SignalInterface) EXPECT() *SignalInterface_Expecter {
	return &SignalInterface_Expecter{mock: &_m.Mock}
}

// GetOrCreateSignal provides a mock function with given fields: ctx, request
func (_m *SignalInterface) GetOrCreateSignal(ctx context.Context, request admin.SignalGetOrCreateRequest) (*admin.Signal, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetOrCreateSignal")
	}

	var r0 *admin.Signal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, admin.SignalGetOrCreateRequest) (*admin.Signal, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, admin.SignalGetOrCreateRequest) *admin.Signal); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Signal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, admin.SignalGetOrCreateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignalInterface_GetOrCreateSignal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrCreateSignal'
type SignalInterface_GetOrCreateSignal_Call struct {
	*mock.Call
}

// GetOrCreateSignal is a helper method to define mock.On call
//   - ctx context.Context
//   - request admin.SignalGetOrCreateRequest
func (_e *SignalInterface_Expecter) GetOrCreateSignal(ctx interface{}, request interface{}) *SignalInterface_GetOrCreateSignal_Call {
	return &SignalInterface_GetOrCreateSignal_Call{Call: _e.mock.On("GetOrCreateSignal", ctx, request)}
}

func (_c *SignalInterface_GetOrCreateSignal_Call) Run(run func(ctx context.Context, request admin.SignalGetOrCreateRequest)) *SignalInterface_GetOrCreateSignal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(admin.SignalGetOrCreateRequest))
	})
	return _c
}

func (_c *SignalInterface_GetOrCreateSignal_Call) Return(_a0 *admin.Signal, _a1 error) *SignalInterface_GetOrCreateSignal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SignalInterface_GetOrCreateSignal_Call) RunAndReturn(run func(context.Context, admin.SignalGetOrCreateRequest) (*admin.Signal, error)) *SignalInterface_GetOrCreateSignal_Call {
	_c.Call.Return(run)
	return _c
}

// ListSignals provides a mock function with given fields: ctx, request
func (_m *SignalInterface) ListSignals(ctx context.Context, request admin.SignalListRequest) (*admin.SignalList, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListSignals")
	}

	var r0 *admin.SignalList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, admin.SignalListRequest) (*admin.SignalList, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, admin.SignalListRequest) *admin.SignalList); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.SignalList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, admin.SignalListRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignalInterface_ListSignals_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSignals'
type SignalInterface_ListSignals_Call struct {
	*mock.Call
}

// ListSignals is a helper method to define mock.On call
//   - ctx context.Context
//   - request admin.SignalListRequest
func (_e *SignalInterface_Expecter) ListSignals(ctx interface{}, request interface{}) *SignalInterface_ListSignals_Call {
	return &SignalInterface_ListSignals_Call{Call: _e.mock.On("ListSignals", ctx, request)}
}

func (_c *SignalInterface_ListSignals_Call) Run(run func(ctx context.Context, request admin.SignalListRequest)) *SignalInterface_ListSignals_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(admin.SignalListRequest))
	})
	return _c
}

func (_c *SignalInterface_ListSignals_Call) Return(_a0 *admin.SignalList, _a1 error) *SignalInterface_ListSignals_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SignalInterface_ListSignals_Call) RunAndReturn(run func(context.Context, admin.SignalListRequest) (*admin.SignalList, error)) *SignalInterface_ListSignals_Call {
	_c.Call.Return(run)
	return _c
}

// SetSignal provides a mock function with given fields: ctx, request
func (_m *SignalInterface) SetSignal(ctx context.Context, request admin.SignalSetRequest) (*admin.SignalSetResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SetSignal")
	}

	var r0 *admin.SignalSetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, admin.SignalSetRequest) (*admin.SignalSetResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, admin.SignalSetRequest) *admin.SignalSetResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.SignalSetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, admin.SignalSetRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignalInterface_SetSignal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSignal'
type SignalInterface_SetSignal_Call struct {
	*mock.Call
}

// SetSignal is a helper method to define mock.On call
//   - ctx context.Context
//   - request admin.SignalSetRequest
func (_e *SignalInterface_Expecter) SetSignal(ctx interface{}, request interface{}) *SignalInterface_SetSignal_Call {
	return &SignalInterface_SetSignal_Call{Call: _e.mock.On("SetSignal", ctx, request)}
}

func (_c *SignalInterface_SetSignal_Call) Run(run func(ctx context.Context, request admin.SignalSetRequest)) *SignalInterface_SetSignal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(admin.SignalSetRequest))
	})
	return _c
}

func (_c *SignalInterface_SetSignal_Call) Return(_a0 *admin.SignalSetResponse, _a1 error) *SignalInterface_SetSignal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SignalInterface_SetSignal_Call) RunAndReturn(run func(context.Context, admin.SignalSetRequest) (*admin.SignalSetResponse, error)) *SignalInterface_SetSignal_Call {
	_c.Call.Return(run)
	return _c
}

// NewSignalInterface creates a new instance of SignalInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSignalInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SignalInterface {
	mock := &SignalInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
