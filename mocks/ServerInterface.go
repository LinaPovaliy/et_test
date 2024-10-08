// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"
	et_test "et_test/proto/et_test"

	mock "github.com/stretchr/testify/mock"
)

// ServerInterface is an autogenerated mock type for the ServerInterface type
type ServerInterface struct {
	mock.Mock
}

// GetThumbnail provides a mock function with given fields: ctx, req
func (_m *ServerInterface) GetThumbnail(ctx context.Context, req *et_test.GetThumbnailRequest) (*et_test.GetThumbnailResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetThumbnail")
	}

	var r0 *et_test.GetThumbnailResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *et_test.GetThumbnailRequest) (*et_test.GetThumbnailResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *et_test.GetThumbnailRequest) *et_test.GetThumbnailResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*et_test.GetThumbnailResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *et_test.GetThumbnailRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServerInterface creates a new instance of ServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerInterface {
	mock := &ServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
