// Code generated by mockery. DO NOT EDIT.

package translation

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	deployment "github.com/mongodb/mongodb-atlas-kubernetes/v2/internal/translation/deployment"
)

// AtlasDeploymentsServiceMock is an autogenerated mock type for the AtlasDeploymentsService type
type AtlasDeploymentsServiceMock struct {
	mock.Mock
}

type AtlasDeploymentsServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *AtlasDeploymentsServiceMock) EXPECT() *AtlasDeploymentsServiceMock_Expecter {
	return &AtlasDeploymentsServiceMock_Expecter{mock: &_m.Mock}
}

// ClusterExists provides a mock function with given fields: ctx, projectID, clusterName
func (_m *AtlasDeploymentsServiceMock) ClusterExists(ctx context.Context, projectID string, clusterName string) (bool, error) {
	ret := _m.Called(ctx, projectID, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for ClusterExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, projectID, clusterName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, projectID, clusterName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, projectID, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_ClusterExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterExists'
type AtlasDeploymentsServiceMock_ClusterExists_Call struct {
	*mock.Call
}

// ClusterExists is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
func (_e *AtlasDeploymentsServiceMock_Expecter) ClusterExists(ctx interface{}, projectID interface{}, clusterName interface{}) *AtlasDeploymentsServiceMock_ClusterExists_Call {
	return &AtlasDeploymentsServiceMock_ClusterExists_Call{Call: _e.mock.On("ClusterExists", ctx, projectID, clusterName)}
}

func (_c *AtlasDeploymentsServiceMock_ClusterExists_Call) Run(run func(ctx context.Context, projectID string, clusterName string)) *AtlasDeploymentsServiceMock_ClusterExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ClusterExists_Call) Return(_a0 bool, _a1 error) *AtlasDeploymentsServiceMock_ClusterExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ClusterExists_Call) RunAndReturn(run func(context.Context, string, string) (bool, error)) *AtlasDeploymentsServiceMock_ClusterExists_Call {
	_c.Call.Return(run)
	return _c
}

// ClusterWithProcessArgs provides a mock function with given fields: ctx, cluster
func (_m *AtlasDeploymentsServiceMock) ClusterWithProcessArgs(ctx context.Context, cluster *deployment.Cluster) error {
	ret := _m.Called(ctx, cluster)

	if len(ret) == 0 {
		panic("no return value specified for ClusterWithProcessArgs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *deployment.Cluster) error); ok {
		r0 = rf(ctx, cluster)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterWithProcessArgs'
type AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call struct {
	*mock.Call
}

// ClusterWithProcessArgs is a helper method to define mock.On call
//   - ctx context.Context
//   - cluster *deployment.Cluster
func (_e *AtlasDeploymentsServiceMock_Expecter) ClusterWithProcessArgs(ctx interface{}, cluster interface{}) *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call {
	return &AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call{Call: _e.mock.On("ClusterWithProcessArgs", ctx, cluster)}
}

func (_c *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call) Run(run func(ctx context.Context, cluster *deployment.Cluster)) *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*deployment.Cluster))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call) Return(_a0 error) *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call) RunAndReturn(run func(context.Context, *deployment.Cluster) error) *AtlasDeploymentsServiceMock_ClusterWithProcessArgs_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCustomZones provides a mock function with given fields: ctx, projectID, clusterName, mappings
func (_m *AtlasDeploymentsServiceMock) CreateCustomZones(ctx context.Context, projectID string, clusterName string, mappings []v1.CustomZoneMapping) (map[string]string, error) {
	ret := _m.Called(ctx, projectID, clusterName, mappings)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomZones")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []v1.CustomZoneMapping) (map[string]string, error)); ok {
		return rf(ctx, projectID, clusterName, mappings)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []v1.CustomZoneMapping) map[string]string); ok {
		r0 = rf(ctx, projectID, clusterName, mappings)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, []v1.CustomZoneMapping) error); ok {
		r1 = rf(ctx, projectID, clusterName, mappings)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_CreateCustomZones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomZones'
type AtlasDeploymentsServiceMock_CreateCustomZones_Call struct {
	*mock.Call
}

// CreateCustomZones is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
//   - mappings []v1.CustomZoneMapping
func (_e *AtlasDeploymentsServiceMock_Expecter) CreateCustomZones(ctx interface{}, projectID interface{}, clusterName interface{}, mappings interface{}) *AtlasDeploymentsServiceMock_CreateCustomZones_Call {
	return &AtlasDeploymentsServiceMock_CreateCustomZones_Call{Call: _e.mock.On("CreateCustomZones", ctx, projectID, clusterName, mappings)}
}

func (_c *AtlasDeploymentsServiceMock_CreateCustomZones_Call) Run(run func(ctx context.Context, projectID string, clusterName string, mappings []v1.CustomZoneMapping)) *AtlasDeploymentsServiceMock_CreateCustomZones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].([]v1.CustomZoneMapping))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_CreateCustomZones_Call) Return(_a0 map[string]string, _a1 error) *AtlasDeploymentsServiceMock_CreateCustomZones_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_CreateCustomZones_Call) RunAndReturn(run func(context.Context, string, string, []v1.CustomZoneMapping) (map[string]string, error)) *AtlasDeploymentsServiceMock_CreateCustomZones_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDeployment provides a mock function with given fields: ctx, _a1
func (_m *AtlasDeploymentsServiceMock) CreateDeployment(ctx context.Context, _a1 deployment.Deployment) (deployment.Deployment, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateDeployment")
	}

	var r0 deployment.Deployment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, deployment.Deployment) (deployment.Deployment, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, deployment.Deployment) deployment.Deployment); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(deployment.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, deployment.Deployment) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_CreateDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDeployment'
type AtlasDeploymentsServiceMock_CreateDeployment_Call struct {
	*mock.Call
}

// CreateDeployment is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 deployment.Deployment
func (_e *AtlasDeploymentsServiceMock_Expecter) CreateDeployment(ctx interface{}, _a1 interface{}) *AtlasDeploymentsServiceMock_CreateDeployment_Call {
	return &AtlasDeploymentsServiceMock_CreateDeployment_Call{Call: _e.mock.On("CreateDeployment", ctx, _a1)}
}

func (_c *AtlasDeploymentsServiceMock_CreateDeployment_Call) Run(run func(ctx context.Context, _a1 deployment.Deployment)) *AtlasDeploymentsServiceMock_CreateDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(deployment.Deployment))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_CreateDeployment_Call) Return(_a0 deployment.Deployment, _a1 error) *AtlasDeploymentsServiceMock_CreateDeployment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_CreateDeployment_Call) RunAndReturn(run func(context.Context, deployment.Deployment) (deployment.Deployment, error)) *AtlasDeploymentsServiceMock_CreateDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// CreateManagedNamespace provides a mock function with given fields: ctx, projectID, clusterName, ns
func (_m *AtlasDeploymentsServiceMock) CreateManagedNamespace(ctx context.Context, projectID string, clusterName string, ns *v1.ManagedNamespace) error {
	ret := _m.Called(ctx, projectID, clusterName, ns)

	if len(ret) == 0 {
		panic("no return value specified for CreateManagedNamespace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *v1.ManagedNamespace) error); ok {
		r0 = rf(ctx, projectID, clusterName, ns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AtlasDeploymentsServiceMock_CreateManagedNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateManagedNamespace'
type AtlasDeploymentsServiceMock_CreateManagedNamespace_Call struct {
	*mock.Call
}

// CreateManagedNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
//   - ns *v1.ManagedNamespace
func (_e *AtlasDeploymentsServiceMock_Expecter) CreateManagedNamespace(ctx interface{}, projectID interface{}, clusterName interface{}, ns interface{}) *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call {
	return &AtlasDeploymentsServiceMock_CreateManagedNamespace_Call{Call: _e.mock.On("CreateManagedNamespace", ctx, projectID, clusterName, ns)}
}

func (_c *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call) Run(run func(ctx context.Context, projectID string, clusterName string, ns *v1.ManagedNamespace)) *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*v1.ManagedNamespace))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call) Return(_a0 error) *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call) RunAndReturn(run func(context.Context, string, string, *v1.ManagedNamespace) error) *AtlasDeploymentsServiceMock_CreateManagedNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCustomZones provides a mock function with given fields: ctx, projectID, clusterName
func (_m *AtlasDeploymentsServiceMock) DeleteCustomZones(ctx context.Context, projectID string, clusterName string) error {
	ret := _m.Called(ctx, projectID, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomZones")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, projectID, clusterName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AtlasDeploymentsServiceMock_DeleteCustomZones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCustomZones'
type AtlasDeploymentsServiceMock_DeleteCustomZones_Call struct {
	*mock.Call
}

// DeleteCustomZones is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
func (_e *AtlasDeploymentsServiceMock_Expecter) DeleteCustomZones(ctx interface{}, projectID interface{}, clusterName interface{}) *AtlasDeploymentsServiceMock_DeleteCustomZones_Call {
	return &AtlasDeploymentsServiceMock_DeleteCustomZones_Call{Call: _e.mock.On("DeleteCustomZones", ctx, projectID, clusterName)}
}

func (_c *AtlasDeploymentsServiceMock_DeleteCustomZones_Call) Run(run func(ctx context.Context, projectID string, clusterName string)) *AtlasDeploymentsServiceMock_DeleteCustomZones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeleteCustomZones_Call) Return(_a0 error) *AtlasDeploymentsServiceMock_DeleteCustomZones_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeleteCustomZones_Call) RunAndReturn(run func(context.Context, string, string) error) *AtlasDeploymentsServiceMock_DeleteCustomZones_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDeployment provides a mock function with given fields: ctx, _a1
func (_m *AtlasDeploymentsServiceMock) DeleteDeployment(ctx context.Context, _a1 deployment.Deployment) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDeployment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, deployment.Deployment) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AtlasDeploymentsServiceMock_DeleteDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDeployment'
type AtlasDeploymentsServiceMock_DeleteDeployment_Call struct {
	*mock.Call
}

// DeleteDeployment is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 deployment.Deployment
func (_e *AtlasDeploymentsServiceMock_Expecter) DeleteDeployment(ctx interface{}, _a1 interface{}) *AtlasDeploymentsServiceMock_DeleteDeployment_Call {
	return &AtlasDeploymentsServiceMock_DeleteDeployment_Call{Call: _e.mock.On("DeleteDeployment", ctx, _a1)}
}

func (_c *AtlasDeploymentsServiceMock_DeleteDeployment_Call) Run(run func(ctx context.Context, _a1 deployment.Deployment)) *AtlasDeploymentsServiceMock_DeleteDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(deployment.Deployment))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeleteDeployment_Call) Return(_a0 error) *AtlasDeploymentsServiceMock_DeleteDeployment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeleteDeployment_Call) RunAndReturn(run func(context.Context, deployment.Deployment) error) *AtlasDeploymentsServiceMock_DeleteDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteManagedNamespace provides a mock function with given fields: ctx, projectID, clusterName, ns
func (_m *AtlasDeploymentsServiceMock) DeleteManagedNamespace(ctx context.Context, projectID string, clusterName string, ns *v1.ManagedNamespace) error {
	ret := _m.Called(ctx, projectID, clusterName, ns)

	if len(ret) == 0 {
		panic("no return value specified for DeleteManagedNamespace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *v1.ManagedNamespace) error); ok {
		r0 = rf(ctx, projectID, clusterName, ns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteManagedNamespace'
type AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call struct {
	*mock.Call
}

// DeleteManagedNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
//   - ns *v1.ManagedNamespace
func (_e *AtlasDeploymentsServiceMock_Expecter) DeleteManagedNamespace(ctx interface{}, projectID interface{}, clusterName interface{}, ns interface{}) *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call {
	return &AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call{Call: _e.mock.On("DeleteManagedNamespace", ctx, projectID, clusterName, ns)}
}

func (_c *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call) Run(run func(ctx context.Context, projectID string, clusterName string, ns *v1.ManagedNamespace)) *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*v1.ManagedNamespace))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call) Return(_a0 error) *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call) RunAndReturn(run func(context.Context, string, string, *v1.ManagedNamespace) error) *AtlasDeploymentsServiceMock_DeleteManagedNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// DeploymentIsReady provides a mock function with given fields: ctx, projectID, deploymentName
func (_m *AtlasDeploymentsServiceMock) DeploymentIsReady(ctx context.Context, projectID string, deploymentName string) (bool, error) {
	ret := _m.Called(ctx, projectID, deploymentName)

	if len(ret) == 0 {
		panic("no return value specified for DeploymentIsReady")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, projectID, deploymentName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, projectID, deploymentName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, projectID, deploymentName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_DeploymentIsReady_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeploymentIsReady'
type AtlasDeploymentsServiceMock_DeploymentIsReady_Call struct {
	*mock.Call
}

// DeploymentIsReady is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - deploymentName string
func (_e *AtlasDeploymentsServiceMock_Expecter) DeploymentIsReady(ctx interface{}, projectID interface{}, deploymentName interface{}) *AtlasDeploymentsServiceMock_DeploymentIsReady_Call {
	return &AtlasDeploymentsServiceMock_DeploymentIsReady_Call{Call: _e.mock.On("DeploymentIsReady", ctx, projectID, deploymentName)}
}

func (_c *AtlasDeploymentsServiceMock_DeploymentIsReady_Call) Run(run func(ctx context.Context, projectID string, deploymentName string)) *AtlasDeploymentsServiceMock_DeploymentIsReady_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeploymentIsReady_Call) Return(_a0 bool, _a1 error) *AtlasDeploymentsServiceMock_DeploymentIsReady_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_DeploymentIsReady_Call) RunAndReturn(run func(context.Context, string, string) (bool, error)) *AtlasDeploymentsServiceMock_DeploymentIsReady_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomZones provides a mock function with given fields: ctx, projectID, clusterName
func (_m *AtlasDeploymentsServiceMock) GetCustomZones(ctx context.Context, projectID string, clusterName string) (map[string]string, error) {
	ret := _m.Called(ctx, projectID, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomZones")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (map[string]string, error)); ok {
		return rf(ctx, projectID, clusterName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]string); ok {
		r0 = rf(ctx, projectID, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, projectID, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_GetCustomZones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomZones'
type AtlasDeploymentsServiceMock_GetCustomZones_Call struct {
	*mock.Call
}

// GetCustomZones is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
func (_e *AtlasDeploymentsServiceMock_Expecter) GetCustomZones(ctx interface{}, projectID interface{}, clusterName interface{}) *AtlasDeploymentsServiceMock_GetCustomZones_Call {
	return &AtlasDeploymentsServiceMock_GetCustomZones_Call{Call: _e.mock.On("GetCustomZones", ctx, projectID, clusterName)}
}

func (_c *AtlasDeploymentsServiceMock_GetCustomZones_Call) Run(run func(ctx context.Context, projectID string, clusterName string)) *AtlasDeploymentsServiceMock_GetCustomZones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetCustomZones_Call) Return(_a0 map[string]string, _a1 error) *AtlasDeploymentsServiceMock_GetCustomZones_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetCustomZones_Call) RunAndReturn(run func(context.Context, string, string) (map[string]string, error)) *AtlasDeploymentsServiceMock_GetCustomZones_Call {
	_c.Call.Return(run)
	return _c
}

// GetDeployment provides a mock function with given fields: ctx, projectID, _a2
func (_m *AtlasDeploymentsServiceMock) GetDeployment(ctx context.Context, projectID string, _a2 *v1.AtlasDeployment) (deployment.Deployment, error) {
	ret := _m.Called(ctx, projectID, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetDeployment")
	}

	var r0 deployment.Deployment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *v1.AtlasDeployment) (deployment.Deployment, error)); ok {
		return rf(ctx, projectID, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *v1.AtlasDeployment) deployment.Deployment); ok {
		r0 = rf(ctx, projectID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(deployment.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *v1.AtlasDeployment) error); ok {
		r1 = rf(ctx, projectID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_GetDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDeployment'
type AtlasDeploymentsServiceMock_GetDeployment_Call struct {
	*mock.Call
}

// GetDeployment is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - _a2 *v1.AtlasDeployment
func (_e *AtlasDeploymentsServiceMock_Expecter) GetDeployment(ctx interface{}, projectID interface{}, _a2 interface{}) *AtlasDeploymentsServiceMock_GetDeployment_Call {
	return &AtlasDeploymentsServiceMock_GetDeployment_Call{Call: _e.mock.On("GetDeployment", ctx, projectID, _a2)}
}

func (_c *AtlasDeploymentsServiceMock_GetDeployment_Call) Run(run func(ctx context.Context, projectID string, _a2 *v1.AtlasDeployment)) *AtlasDeploymentsServiceMock_GetDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*v1.AtlasDeployment))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetDeployment_Call) Return(_a0 deployment.Deployment, _a1 error) *AtlasDeploymentsServiceMock_GetDeployment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetDeployment_Call) RunAndReturn(run func(context.Context, string, *v1.AtlasDeployment) (deployment.Deployment, error)) *AtlasDeploymentsServiceMock_GetDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// GetManagedNamespaces provides a mock function with given fields: ctx, projectID, clusterName
func (_m *AtlasDeploymentsServiceMock) GetManagedNamespaces(ctx context.Context, projectID string, clusterName string) ([]v1.ManagedNamespace, error) {
	ret := _m.Called(ctx, projectID, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetManagedNamespaces")
	}

	var r0 []v1.ManagedNamespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]v1.ManagedNamespace, error)); ok {
		return rf(ctx, projectID, clusterName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []v1.ManagedNamespace); ok {
		r0 = rf(ctx, projectID, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.ManagedNamespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, projectID, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_GetManagedNamespaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManagedNamespaces'
type AtlasDeploymentsServiceMock_GetManagedNamespaces_Call struct {
	*mock.Call
}

// GetManagedNamespaces is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - clusterName string
func (_e *AtlasDeploymentsServiceMock_Expecter) GetManagedNamespaces(ctx interface{}, projectID interface{}, clusterName interface{}) *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call {
	return &AtlasDeploymentsServiceMock_GetManagedNamespaces_Call{Call: _e.mock.On("GetManagedNamespaces", ctx, projectID, clusterName)}
}

func (_c *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call) Run(run func(ctx context.Context, projectID string, clusterName string)) *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call) Return(_a0 []v1.ManagedNamespace, _a1 error) *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call) RunAndReturn(run func(context.Context, string, string) ([]v1.ManagedNamespace, error)) *AtlasDeploymentsServiceMock_GetManagedNamespaces_Call {
	_c.Call.Return(run)
	return _c
}

// GetZoneMapping provides a mock function with given fields: ctx, projectID, deploymentName
func (_m *AtlasDeploymentsServiceMock) GetZoneMapping(ctx context.Context, projectID string, deploymentName string) (map[string]string, error) {
	ret := _m.Called(ctx, projectID, deploymentName)

	if len(ret) == 0 {
		panic("no return value specified for GetZoneMapping")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (map[string]string, error)); ok {
		return rf(ctx, projectID, deploymentName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]string); ok {
		r0 = rf(ctx, projectID, deploymentName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, projectID, deploymentName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_GetZoneMapping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetZoneMapping'
type AtlasDeploymentsServiceMock_GetZoneMapping_Call struct {
	*mock.Call
}

// GetZoneMapping is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - deploymentName string
func (_e *AtlasDeploymentsServiceMock_Expecter) GetZoneMapping(ctx interface{}, projectID interface{}, deploymentName interface{}) *AtlasDeploymentsServiceMock_GetZoneMapping_Call {
	return &AtlasDeploymentsServiceMock_GetZoneMapping_Call{Call: _e.mock.On("GetZoneMapping", ctx, projectID, deploymentName)}
}

func (_c *AtlasDeploymentsServiceMock_GetZoneMapping_Call) Run(run func(ctx context.Context, projectID string, deploymentName string)) *AtlasDeploymentsServiceMock_GetZoneMapping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetZoneMapping_Call) Return(_a0 map[string]string, _a1 error) *AtlasDeploymentsServiceMock_GetZoneMapping_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_GetZoneMapping_Call) RunAndReturn(run func(context.Context, string, string) (map[string]string, error)) *AtlasDeploymentsServiceMock_GetZoneMapping_Call {
	_c.Call.Return(run)
	return _c
}

// ListDeploymentConnections provides a mock function with given fields: ctx, projectID
func (_m *AtlasDeploymentsServiceMock) ListDeploymentConnections(ctx context.Context, projectID string) ([]deployment.Connection, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for ListDeploymentConnections")
	}

	var r0 []deployment.Connection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]deployment.Connection, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []deployment.Connection); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]deployment.Connection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_ListDeploymentConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDeploymentConnections'
type AtlasDeploymentsServiceMock_ListDeploymentConnections_Call struct {
	*mock.Call
}

// ListDeploymentConnections is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
func (_e *AtlasDeploymentsServiceMock_Expecter) ListDeploymentConnections(ctx interface{}, projectID interface{}) *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call {
	return &AtlasDeploymentsServiceMock_ListDeploymentConnections_Call{Call: _e.mock.On("ListDeploymentConnections", ctx, projectID)}
}

func (_c *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call) Run(run func(ctx context.Context, projectID string)) *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call) Return(_a0 []deployment.Connection, _a1 error) *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call) RunAndReturn(run func(context.Context, string) ([]deployment.Connection, error)) *AtlasDeploymentsServiceMock_ListDeploymentConnections_Call {
	_c.Call.Return(run)
	return _c
}

// ListDeploymentNames provides a mock function with given fields: ctx, projectID
func (_m *AtlasDeploymentsServiceMock) ListDeploymentNames(ctx context.Context, projectID string) ([]string, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for ListDeploymentNames")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_ListDeploymentNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDeploymentNames'
type AtlasDeploymentsServiceMock_ListDeploymentNames_Call struct {
	*mock.Call
}

// ListDeploymentNames is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
func (_e *AtlasDeploymentsServiceMock_Expecter) ListDeploymentNames(ctx interface{}, projectID interface{}) *AtlasDeploymentsServiceMock_ListDeploymentNames_Call {
	return &AtlasDeploymentsServiceMock_ListDeploymentNames_Call{Call: _e.mock.On("ListDeploymentNames", ctx, projectID)}
}

func (_c *AtlasDeploymentsServiceMock_ListDeploymentNames_Call) Run(run func(ctx context.Context, projectID string)) *AtlasDeploymentsServiceMock_ListDeploymentNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ListDeploymentNames_Call) Return(_a0 []string, _a1 error) *AtlasDeploymentsServiceMock_ListDeploymentNames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_ListDeploymentNames_Call) RunAndReturn(run func(context.Context, string) ([]string, error)) *AtlasDeploymentsServiceMock_ListDeploymentNames_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDeployment provides a mock function with given fields: ctx, _a1
func (_m *AtlasDeploymentsServiceMock) UpdateDeployment(ctx context.Context, _a1 deployment.Deployment) (deployment.Deployment, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeployment")
	}

	var r0 deployment.Deployment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, deployment.Deployment) (deployment.Deployment, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, deployment.Deployment) deployment.Deployment); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(deployment.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, deployment.Deployment) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtlasDeploymentsServiceMock_UpdateDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDeployment'
type AtlasDeploymentsServiceMock_UpdateDeployment_Call struct {
	*mock.Call
}

// UpdateDeployment is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 deployment.Deployment
func (_e *AtlasDeploymentsServiceMock_Expecter) UpdateDeployment(ctx interface{}, _a1 interface{}) *AtlasDeploymentsServiceMock_UpdateDeployment_Call {
	return &AtlasDeploymentsServiceMock_UpdateDeployment_Call{Call: _e.mock.On("UpdateDeployment", ctx, _a1)}
}

func (_c *AtlasDeploymentsServiceMock_UpdateDeployment_Call) Run(run func(ctx context.Context, _a1 deployment.Deployment)) *AtlasDeploymentsServiceMock_UpdateDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(deployment.Deployment))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_UpdateDeployment_Call) Return(_a0 deployment.Deployment, _a1 error) *AtlasDeploymentsServiceMock_UpdateDeployment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_UpdateDeployment_Call) RunAndReturn(run func(context.Context, deployment.Deployment) (deployment.Deployment, error)) *AtlasDeploymentsServiceMock_UpdateDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProcessArgs provides a mock function with given fields: ctx, cluster
func (_m *AtlasDeploymentsServiceMock) UpdateProcessArgs(ctx context.Context, cluster *deployment.Cluster) error {
	ret := _m.Called(ctx, cluster)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProcessArgs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *deployment.Cluster) error); ok {
		r0 = rf(ctx, cluster)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AtlasDeploymentsServiceMock_UpdateProcessArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProcessArgs'
type AtlasDeploymentsServiceMock_UpdateProcessArgs_Call struct {
	*mock.Call
}

// UpdateProcessArgs is a helper method to define mock.On call
//   - ctx context.Context
//   - cluster *deployment.Cluster
func (_e *AtlasDeploymentsServiceMock_Expecter) UpdateProcessArgs(ctx interface{}, cluster interface{}) *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call {
	return &AtlasDeploymentsServiceMock_UpdateProcessArgs_Call{Call: _e.mock.On("UpdateProcessArgs", ctx, cluster)}
}

func (_c *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call) Run(run func(ctx context.Context, cluster *deployment.Cluster)) *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*deployment.Cluster))
	})
	return _c
}

func (_c *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call) Return(_a0 error) *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call) RunAndReturn(run func(context.Context, *deployment.Cluster) error) *AtlasDeploymentsServiceMock_UpdateProcessArgs_Call {
	_c.Call.Return(run)
	return _c
}

// NewAtlasDeploymentsServiceMock creates a new instance of AtlasDeploymentsServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAtlasDeploymentsServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *AtlasDeploymentsServiceMock {
	mock := &AtlasDeploymentsServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
