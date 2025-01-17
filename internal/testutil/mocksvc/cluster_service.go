// Code generated by mockery. DO NOT EDIT.

package mocksvc

import (
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"

	context "context"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ClusterService is an autogenerated mock type for the ClusterService type
type ClusterService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, groupID, clusterName
func (_m *ClusterService) Get(ctx context.Context, groupID string, clusterName string) (*admin.AdvancedClusterDescription, *http.Response, error) {
	ret := _m.Called(ctx, groupID, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *admin.AdvancedClusterDescription
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*admin.AdvancedClusterDescription, *http.Response, error)); ok {
		return rf(ctx, groupID, clusterName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *admin.AdvancedClusterDescription); ok {
		r0 = rf(ctx, groupID, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AdvancedClusterDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *http.Response); ok {
		r1 = rf(ctx, groupID, clusterName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, groupID, clusterName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// List provides a mock function with given fields: ctx, options
func (_m *ClusterService) List(ctx context.Context, options *admin.ListClustersApiParams) (*admin.PaginatedAdvancedClusterDescription, *http.Response, error) {
	ret := _m.Called(ctx, options)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *admin.PaginatedAdvancedClusterDescription
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListClustersApiParams) (*admin.PaginatedAdvancedClusterDescription, *http.Response, error)); ok {
		return rf(ctx, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListClustersApiParams) *admin.PaginatedAdvancedClusterDescription); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedAdvancedClusterDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admin.ListClustersApiParams) *http.Response); ok {
		r1 = rf(ctx, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *admin.ListClustersApiParams) error); ok {
		r2 = rf(ctx, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewClusterService creates a new instance of ClusterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClusterService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClusterService {
	mock := &ClusterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
