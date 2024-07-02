// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	container "github.com/docker/docker/api/types/container"

	io "io"

	mock "github.com/stretchr/testify/mock"

	network "github.com/docker/docker/api/types/network"

	types "github.com/docker/docker/api/types"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	volume "github.com/docker/docker/api/types/volume"
)

// Docker is an autogenerated mock type for the Docker type
type Docker struct {
	mock.Mock
}

type Docker_ContainerCreate struct {
	*mock.Call
}

func (_m Docker_ContainerCreate) Return(_a0 container.CreateResponse, _a1 error) *Docker_ContainerCreate {
	return &Docker_ContainerCreate{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) *Docker_ContainerCreate {
	c_call := _m.On("ContainerCreate", ctx, config, hostConfig, networkingConfig, platform, containerName)
	return &Docker_ContainerCreate{Call: c_call}
}

func (_m *Docker) OnContainerCreateMatch(matchers ...interface{}) *Docker_ContainerCreate {
	c_call := _m.On("ContainerCreate", matchers...)
	return &Docker_ContainerCreate{Call: c_call}
}

// ContainerCreate provides a mock function with given fields: ctx, config, hostConfig, networkingConfig, platform, containerName
func (_m *Docker) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.CreateResponse, error) {
	ret := _m.Called(ctx, config, hostConfig, networkingConfig, platform, containerName)

	var r0 container.CreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *v1.Platform, string) container.CreateResponse); ok {
		r0 = rf(ctx, config, hostConfig, networkingConfig, platform, containerName)
	} else {
		r0 = ret.Get(0).(container.CreateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *v1.Platform, string) error); ok {
		r1 = rf(ctx, config, hostConfig, networkingConfig, platform, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerExecAttach struct {
	*mock.Call
}

func (_m Docker_ContainerExecAttach) Return(_a0 types.HijackedResponse, _a1 error) *Docker_ContainerExecAttach {
	return &Docker_ContainerExecAttach{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) *Docker_ContainerExecAttach {
	c_call := _m.On("ContainerExecAttach", ctx, execID, config)
	return &Docker_ContainerExecAttach{Call: c_call}
}

func (_m *Docker) OnContainerExecAttachMatch(matchers ...interface{}) *Docker_ContainerExecAttach {
	c_call := _m.On("ContainerExecAttach", matchers...)
	return &Docker_ContainerExecAttach{Call: c_call}
}

// ContainerExecAttach provides a mock function with given fields: ctx, execID, config
func (_m *Docker) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, execID, config)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecStartCheck) types.HijackedResponse); ok {
		r0 = rf(ctx, execID, config)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecStartCheck) error); ok {
		r1 = rf(ctx, execID, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerExecCreate struct {
	*mock.Call
}

func (_m Docker_ContainerExecCreate) Return(_a0 types.IDResponse, _a1 error) *Docker_ContainerExecCreate {
	return &Docker_ContainerExecCreate{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerExecCreate(ctx context.Context, _a1 string, config types.ExecConfig) *Docker_ContainerExecCreate {
	c_call := _m.On("ContainerExecCreate", ctx, _a1, config)
	return &Docker_ContainerExecCreate{Call: c_call}
}

func (_m *Docker) OnContainerExecCreateMatch(matchers ...interface{}) *Docker_ContainerExecCreate {
	c_call := _m.On("ContainerExecCreate", matchers...)
	return &Docker_ContainerExecCreate{Call: c_call}
}

// ContainerExecCreate provides a mock function with given fields: ctx, _a1, config
func (_m *Docker) ContainerExecCreate(ctx context.Context, _a1 string, config types.ExecConfig) (types.IDResponse, error) {
	ret := _m.Called(ctx, _a1, config)

	var r0 types.IDResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecConfig) types.IDResponse); ok {
		r0 = rf(ctx, _a1, config)
	} else {
		r0 = ret.Get(0).(types.IDResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecConfig) error); ok {
		r1 = rf(ctx, _a1, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerExecInspect struct {
	*mock.Call
}

func (_m Docker_ContainerExecInspect) Return(_a0 types.ContainerExecInspect, _a1 error) *Docker_ContainerExecInspect {
	return &Docker_ContainerExecInspect{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerExecInspect(ctx context.Context, execID string) *Docker_ContainerExecInspect {
	c_call := _m.On("ContainerExecInspect", ctx, execID)
	return &Docker_ContainerExecInspect{Call: c_call}
}

func (_m *Docker) OnContainerExecInspectMatch(matchers ...interface{}) *Docker_ContainerExecInspect {
	c_call := _m.On("ContainerExecInspect", matchers...)
	return &Docker_ContainerExecInspect{Call: c_call}
}

// ContainerExecInspect provides a mock function with given fields: ctx, execID
func (_m *Docker) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error) {
	ret := _m.Called(ctx, execID)

	var r0 types.ContainerExecInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerExecInspect); ok {
		r0 = rf(ctx, execID)
	} else {
		r0 = ret.Get(0).(types.ContainerExecInspect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, execID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerList struct {
	*mock.Call
}

func (_m Docker_ContainerList) Return(_a0 []types.Container, _a1 error) *Docker_ContainerList {
	return &Docker_ContainerList{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerList(ctx context.Context, options types.ContainerListOptions) *Docker_ContainerList {
	c_call := _m.On("ContainerList", ctx, options)
	return &Docker_ContainerList{Call: c_call}
}

func (_m *Docker) OnContainerListMatch(matchers ...interface{}) *Docker_ContainerList {
	c_call := _m.On("ContainerList", matchers...)
	return &Docker_ContainerList{Call: c_call}
}

// ContainerList provides a mock function with given fields: ctx, options
func (_m *Docker) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.Container
	if rf, ok := ret.Get(0).(func(context.Context, types.ContainerListOptions) []types.Container); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ContainerListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerLogs struct {
	*mock.Call
}

func (_m Docker_ContainerLogs) Return(_a0 io.ReadCloser, _a1 error) *Docker_ContainerLogs {
	return &Docker_ContainerLogs{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerLogs(ctx context.Context, _a1 string, options types.ContainerLogsOptions) *Docker_ContainerLogs {
	c_call := _m.On("ContainerLogs", ctx, _a1, options)
	return &Docker_ContainerLogs{Call: c_call}
}

func (_m *Docker) OnContainerLogsMatch(matchers ...interface{}) *Docker_ContainerLogs {
	c_call := _m.On("ContainerLogs", matchers...)
	return &Docker_ContainerLogs{Call: c_call}
}

// ContainerLogs provides a mock function with given fields: ctx, _a1, options
func (_m *Docker) ContainerLogs(ctx context.Context, _a1 string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerRemove struct {
	*mock.Call
}

func (_m Docker_ContainerRemove) Return(_a0 error) *Docker_ContainerRemove {
	return &Docker_ContainerRemove{Call: _m.Call.Return(_a0)}
}

func (_m *Docker) OnContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) *Docker_ContainerRemove {
	c_call := _m.On("ContainerRemove", ctx, containerID, options)
	return &Docker_ContainerRemove{Call: c_call}
}

func (_m *Docker) OnContainerRemoveMatch(matchers ...interface{}) *Docker_ContainerRemove {
	c_call := _m.On("ContainerRemove", matchers...)
	return &Docker_ContainerRemove{Call: c_call}
}

// ContainerRemove provides a mock function with given fields: ctx, containerID, options
func (_m *Docker) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerRemoveOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Docker_ContainerStart struct {
	*mock.Call
}

func (_m Docker_ContainerStart) Return(_a0 error) *Docker_ContainerStart {
	return &Docker_ContainerStart{Call: _m.Call.Return(_a0)}
}

func (_m *Docker) OnContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) *Docker_ContainerStart {
	c_call := _m.On("ContainerStart", ctx, containerID, options)
	return &Docker_ContainerStart{Call: c_call}
}

func (_m *Docker) OnContainerStartMatch(matchers ...interface{}) *Docker_ContainerStart {
	c_call := _m.On("ContainerStart", matchers...)
	return &Docker_ContainerStart{Call: c_call}
}

// ContainerStart provides a mock function with given fields: ctx, containerID, options
func (_m *Docker) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerStartOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Docker_ContainerStatPath struct {
	*mock.Call
}

func (_m Docker_ContainerStatPath) Return(_a0 types.ContainerPathStat, _a1 error) *Docker_ContainerStatPath {
	return &Docker_ContainerStatPath{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerStatPath(ctx context.Context, containerID string, path string) *Docker_ContainerStatPath {
	c_call := _m.On("ContainerStatPath", ctx, containerID, path)
	return &Docker_ContainerStatPath{Call: c_call}
}

func (_m *Docker) OnContainerStatPathMatch(matchers ...interface{}) *Docker_ContainerStatPath {
	c_call := _m.On("ContainerStatPath", matchers...)
	return &Docker_ContainerStatPath{Call: c_call}
}

// ContainerStatPath provides a mock function with given fields: ctx, containerID, path
func (_m *Docker) ContainerStatPath(ctx context.Context, containerID string, path string) (types.ContainerPathStat, error) {
	ret := _m.Called(ctx, containerID, path)

	var r0 types.ContainerPathStat
	if rf, ok := ret.Get(0).(func(context.Context, string, string) types.ContainerPathStat); ok {
		r0 = rf(ctx, containerID, path)
	} else {
		r0 = ret.Get(0).(types.ContainerPathStat)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, containerID, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ContainerWait struct {
	*mock.Call
}

func (_m Docker_ContainerWait) Return(_a0 <-chan container.WaitResponse, _a1 <-chan error) *Docker_ContainerWait {
	return &Docker_ContainerWait{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnContainerWait(ctx context.Context, _a1 string, condition container.WaitCondition) *Docker_ContainerWait {
	c_call := _m.On("ContainerWait", ctx, _a1, condition)
	return &Docker_ContainerWait{Call: c_call}
}

func (_m *Docker) OnContainerWaitMatch(matchers ...interface{}) *Docker_ContainerWait {
	c_call := _m.On("ContainerWait", matchers...)
	return &Docker_ContainerWait{Call: c_call}
}

// ContainerWait provides a mock function with given fields: ctx, _a1, condition
func (_m *Docker) ContainerWait(ctx context.Context, _a1 string, condition container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ret := _m.Called(ctx, _a1, condition)

	var r0 <-chan container.WaitResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, container.WaitCondition) <-chan container.WaitResponse); ok {
		r0 = rf(ctx, _a1, condition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan container.WaitResponse)
		}
	}

	var r1 <-chan error
	if rf, ok := ret.Get(1).(func(context.Context, string, container.WaitCondition) <-chan error); ok {
		r1 = rf(ctx, _a1, condition)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

type Docker_CopyFromContainer struct {
	*mock.Call
}

func (_m Docker_CopyFromContainer) Return(_a0 io.ReadCloser, _a1 types.ContainerPathStat, _a2 error) *Docker_CopyFromContainer {
	return &Docker_CopyFromContainer{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *Docker) OnCopyFromContainer(ctx context.Context, containerID string, srcPath string) *Docker_CopyFromContainer {
	c_call := _m.On("CopyFromContainer", ctx, containerID, srcPath)
	return &Docker_CopyFromContainer{Call: c_call}
}

func (_m *Docker) OnCopyFromContainerMatch(matchers ...interface{}) *Docker_CopyFromContainer {
	c_call := _m.On("CopyFromContainer", matchers...)
	return &Docker_CopyFromContainer{Call: c_call}
}

// CopyFromContainer provides a mock function with given fields: ctx, containerID, srcPath
func (_m *Docker) CopyFromContainer(ctx context.Context, containerID string, srcPath string) (io.ReadCloser, types.ContainerPathStat, error) {
	ret := _m.Called(ctx, containerID, srcPath)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, string) io.ReadCloser); ok {
		r0 = rf(ctx, containerID, srcPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 types.ContainerPathStat
	if rf, ok := ret.Get(1).(func(context.Context, string, string) types.ContainerPathStat); ok {
		r1 = rf(ctx, containerID, srcPath)
	} else {
		r1 = ret.Get(1).(types.ContainerPathStat)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, containerID, srcPath)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type Docker_ImageList struct {
	*mock.Call
}

func (_m Docker_ImageList) Return(_a0 []types.ImageSummary, _a1 error) *Docker_ImageList {
	return &Docker_ImageList{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnImageList(ctx context.Context, listOption types.ImageListOptions) *Docker_ImageList {
	c_call := _m.On("ImageList", ctx, listOption)
	return &Docker_ImageList{Call: c_call}
}

func (_m *Docker) OnImageListMatch(matchers ...interface{}) *Docker_ImageList {
	c_call := _m.On("ImageList", matchers...)
	return &Docker_ImageList{Call: c_call}
}

// ImageList provides a mock function with given fields: ctx, listOption
func (_m *Docker) ImageList(ctx context.Context, listOption types.ImageListOptions) ([]types.ImageSummary, error) {
	ret := _m.Called(ctx, listOption)

	var r0 []types.ImageSummary
	if rf, ok := ret.Get(0).(func(context.Context, types.ImageListOptions) []types.ImageSummary); ok {
		r0 = rf(ctx, listOption)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ImageSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ImageListOptions) error); ok {
		r1 = rf(ctx, listOption)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_ImagePull struct {
	*mock.Call
}

func (_m Docker_ImagePull) Return(_a0 io.ReadCloser, _a1 error) *Docker_ImagePull {
	return &Docker_ImagePull{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) *Docker_ImagePull {
	c_call := _m.On("ImagePull", ctx, refStr, options)
	return &Docker_ImagePull{Call: c_call}
}

func (_m *Docker) OnImagePullMatch(matchers ...interface{}) *Docker_ImagePull {
	c_call := _m.On("ImagePull", matchers...)
	return &Docker_ImagePull{Call: c_call}
}

// ImagePull provides a mock function with given fields: ctx, refStr, options
func (_m *Docker) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, refStr, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) io.ReadCloser); ok {
		r0 = rf(ctx, refStr, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImagePullOptions) error); ok {
		r1 = rf(ctx, refStr, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_VolumeCreate struct {
	*mock.Call
}

func (_m Docker_VolumeCreate) Return(_a0 volume.Volume, _a1 error) *Docker_VolumeCreate {
	return &Docker_VolumeCreate{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnVolumeCreate(ctx context.Context, options volume.CreateOptions) *Docker_VolumeCreate {
	c_call := _m.On("VolumeCreate", ctx, options)
	return &Docker_VolumeCreate{Call: c_call}
}

func (_m *Docker) OnVolumeCreateMatch(matchers ...interface{}) *Docker_VolumeCreate {
	c_call := _m.On("VolumeCreate", matchers...)
	return &Docker_VolumeCreate{Call: c_call}
}

// VolumeCreate provides a mock function with given fields: ctx, options
func (_m *Docker) VolumeCreate(ctx context.Context, options volume.CreateOptions) (volume.Volume, error) {
	ret := _m.Called(ctx, options)

	var r0 volume.Volume
	if rf, ok := ret.Get(0).(func(context.Context, volume.CreateOptions) volume.Volume); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Get(0).(volume.Volume)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, volume.CreateOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_VolumeList struct {
	*mock.Call
}

func (_m Docker_VolumeList) Return(_a0 volume.ListResponse, _a1 error) *Docker_VolumeList {
	return &Docker_VolumeList{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Docker) OnVolumeList(ctx context.Context, options volume.ListOptions) *Docker_VolumeList {
	c_call := _m.On("VolumeList", ctx, options)
	return &Docker_VolumeList{Call: c_call}
}

func (_m *Docker) OnVolumeListMatch(matchers ...interface{}) *Docker_VolumeList {
	c_call := _m.On("VolumeList", matchers...)
	return &Docker_VolumeList{Call: c_call}
}

// VolumeList provides a mock function with given fields: ctx, options
func (_m *Docker) VolumeList(ctx context.Context, options volume.ListOptions) (volume.ListResponse, error) {
	ret := _m.Called(ctx, options)

	var r0 volume.ListResponse
	if rf, ok := ret.Get(0).(func(context.Context, volume.ListOptions) volume.ListResponse); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Get(0).(volume.ListResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, volume.ListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Docker_VolumeRemove struct {
	*mock.Call
}

func (_m Docker_VolumeRemove) Return(_a0 error) *Docker_VolumeRemove {
	return &Docker_VolumeRemove{Call: _m.Call.Return(_a0)}
}

func (_m *Docker) OnVolumeRemove(ctx context.Context, volumeID string, force bool) *Docker_VolumeRemove {
	c_call := _m.On("VolumeRemove", ctx, volumeID, force)
	return &Docker_VolumeRemove{Call: c_call}
}

func (_m *Docker) OnVolumeRemoveMatch(matchers ...interface{}) *Docker_VolumeRemove {
	c_call := _m.On("VolumeRemove", matchers...)
	return &Docker_VolumeRemove{Call: c_call}
}

// VolumeRemove provides a mock function with given fields: ctx, volumeID, force
func (_m *Docker) VolumeRemove(ctx context.Context, volumeID string, force bool) error {
	ret := _m.Called(ctx, volumeID, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, volumeID, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
