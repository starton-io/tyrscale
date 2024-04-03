// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/starton-io/tyrscale/manager/api/network/dto"
	mock "github.com/stretchr/testify/mock"

	network "github.com/starton-io/tyrscale/manager/pkg/pb/network"
)

// INetworkService is an autogenerated mock type for the INetworkService type
type INetworkService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *INetworkService) Create(ctx context.Context, _a1 *dto.Network) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.Network) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, name
func (_m *INetworkService) Delete(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, filterParams
func (_m *INetworkService) List(ctx context.Context, filterParams *dto.ListNetworkReq) ([]*network.NetworkModel, error) {
	ret := _m.Called(ctx, filterParams)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*network.NetworkModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListNetworkReq) ([]*network.NetworkModel, error)); ok {
		return rf(ctx, filterParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListNetworkReq) []*network.NetworkModel); ok {
		r0 = rf(ctx, filterParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*network.NetworkModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListNetworkReq) error); ok {
		r1 = rf(ctx, filterParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewINetworkService creates a new instance of INetworkService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewINetworkService(t interface {
	mock.TestingT
	Cleanup(func())
}) *INetworkService {
	mock := &INetworkService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
