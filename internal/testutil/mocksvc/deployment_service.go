// Code generated by mockery. DO NOT EDIT.

package mocksvc

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805005/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// DeploymentService is an autogenerated mock type for the DeploymentService type
type DeploymentService struct {
	mock.Mock
}

// GetAtlasSearchDeployment provides a mock function with given fields: ctx, groupID, clusterName
func (_m *DeploymentService) GetAtlasSearchDeployment(ctx context.Context, groupID string, clusterName string) (*admin.ApiSearchDeploymentResponse, *http.Response, error) {
	ret := _m.Called(ctx, groupID, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetAtlasSearchDeployment")
	}

	var r0 *admin.ApiSearchDeploymentResponse
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*admin.ApiSearchDeploymentResponse, *http.Response, error)); ok {
		return rf(ctx, groupID, clusterName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *admin.ApiSearchDeploymentResponse); ok {
		r0 = rf(ctx, groupID, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ApiSearchDeploymentResponse)
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

// NewDeploymentService creates a new instance of DeploymentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeploymentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeploymentService {
	mock := &DeploymentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
