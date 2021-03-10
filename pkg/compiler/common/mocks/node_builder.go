// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	common "github.com/flyteorg/flytepropeller/pkg/compiler/common"

	mock "github.com/stretchr/testify/mock"
)

// NodeBuilder is an autogenerated mock type for the NodeBuilder type
type NodeBuilder struct {
	mock.Mock
}

type NodeBuilder_GetBranchNode struct {
	*mock.Call
}

func (_m NodeBuilder_GetBranchNode) Return(_a0 *core.BranchNode) *NodeBuilder_GetBranchNode {
	return &NodeBuilder_GetBranchNode{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetBranchNode() *NodeBuilder_GetBranchNode {
	c := _m.On("GetBranchNode")
	return &NodeBuilder_GetBranchNode{Call: c}
}

func (_m *NodeBuilder) OnGetBranchNodeMatch(matchers ...interface{}) *NodeBuilder_GetBranchNode {
	c := _m.On("GetBranchNode", matchers...)
	return &NodeBuilder_GetBranchNode{Call: c}
}

// GetBranchNode provides a mock function with given fields:
func (_m *NodeBuilder) GetBranchNode() *core.BranchNode {
	ret := _m.Called()

	var r0 *core.BranchNode
	if rf, ok := ret.Get(0).(func() *core.BranchNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BranchNode)
		}
	}

	return r0
}

type NodeBuilder_GetCoreNode struct {
	*mock.Call
}

func (_m NodeBuilder_GetCoreNode) Return(_a0 *core.Node) *NodeBuilder_GetCoreNode {
	return &NodeBuilder_GetCoreNode{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetCoreNode() *NodeBuilder_GetCoreNode {
	c := _m.On("GetCoreNode")
	return &NodeBuilder_GetCoreNode{Call: c}
}

func (_m *NodeBuilder) OnGetCoreNodeMatch(matchers ...interface{}) *NodeBuilder_GetCoreNode {
	c := _m.On("GetCoreNode", matchers...)
	return &NodeBuilder_GetCoreNode{Call: c}
}

// GetCoreNode provides a mock function with given fields:
func (_m *NodeBuilder) GetCoreNode() *core.Node {
	ret := _m.Called()

	var r0 *core.Node
	if rf, ok := ret.Get(0).(func() *core.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Node)
		}
	}

	return r0
}

type NodeBuilder_GetId struct {
	*mock.Call
}

func (_m NodeBuilder_GetId) Return(_a0 string) *NodeBuilder_GetId {
	return &NodeBuilder_GetId{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetId() *NodeBuilder_GetId {
	c := _m.On("GetId")
	return &NodeBuilder_GetId{Call: c}
}

func (_m *NodeBuilder) OnGetIdMatch(matchers ...interface{}) *NodeBuilder_GetId {
	c := _m.On("GetId", matchers...)
	return &NodeBuilder_GetId{Call: c}
}

// GetId provides a mock function with given fields:
func (_m *NodeBuilder) GetId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type NodeBuilder_GetInputs struct {
	*mock.Call
}

func (_m NodeBuilder_GetInputs) Return(_a0 []*core.Binding) *NodeBuilder_GetInputs {
	return &NodeBuilder_GetInputs{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetInputs() *NodeBuilder_GetInputs {
	c := _m.On("GetInputs")
	return &NodeBuilder_GetInputs{Call: c}
}

func (_m *NodeBuilder) OnGetInputsMatch(matchers ...interface{}) *NodeBuilder_GetInputs {
	c := _m.On("GetInputs", matchers...)
	return &NodeBuilder_GetInputs{Call: c}
}

// GetInputs provides a mock function with given fields:
func (_m *NodeBuilder) GetInputs() []*core.Binding {
	ret := _m.Called()

	var r0 []*core.Binding
	if rf, ok := ret.Get(0).(func() []*core.Binding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Binding)
		}
	}

	return r0
}

type NodeBuilder_GetInterface struct {
	*mock.Call
}

func (_m NodeBuilder_GetInterface) Return(_a0 *core.TypedInterface) *NodeBuilder_GetInterface {
	return &NodeBuilder_GetInterface{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetInterface() *NodeBuilder_GetInterface {
	c := _m.On("GetInterface")
	return &NodeBuilder_GetInterface{Call: c}
}

func (_m *NodeBuilder) OnGetInterfaceMatch(matchers ...interface{}) *NodeBuilder_GetInterface {
	c := _m.On("GetInterface", matchers...)
	return &NodeBuilder_GetInterface{Call: c}
}

// GetInterface provides a mock function with given fields:
func (_m *NodeBuilder) GetInterface() *core.TypedInterface {
	ret := _m.Called()

	var r0 *core.TypedInterface
	if rf, ok := ret.Get(0).(func() *core.TypedInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TypedInterface)
		}
	}

	return r0
}

type NodeBuilder_GetMetadata struct {
	*mock.Call
}

func (_m NodeBuilder_GetMetadata) Return(_a0 *core.NodeMetadata) *NodeBuilder_GetMetadata {
	return &NodeBuilder_GetMetadata{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetMetadata() *NodeBuilder_GetMetadata {
	c := _m.On("GetMetadata")
	return &NodeBuilder_GetMetadata{Call: c}
}

func (_m *NodeBuilder) OnGetMetadataMatch(matchers ...interface{}) *NodeBuilder_GetMetadata {
	c := _m.On("GetMetadata", matchers...)
	return &NodeBuilder_GetMetadata{Call: c}
}

// GetMetadata provides a mock function with given fields:
func (_m *NodeBuilder) GetMetadata() *core.NodeMetadata {
	ret := _m.Called()

	var r0 *core.NodeMetadata
	if rf, ok := ret.Get(0).(func() *core.NodeMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.NodeMetadata)
		}
	}

	return r0
}

type NodeBuilder_GetOutputAliases struct {
	*mock.Call
}

func (_m NodeBuilder_GetOutputAliases) Return(_a0 []*core.Alias) *NodeBuilder_GetOutputAliases {
	return &NodeBuilder_GetOutputAliases{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetOutputAliases() *NodeBuilder_GetOutputAliases {
	c := _m.On("GetOutputAliases")
	return &NodeBuilder_GetOutputAliases{Call: c}
}

func (_m *NodeBuilder) OnGetOutputAliasesMatch(matchers ...interface{}) *NodeBuilder_GetOutputAliases {
	c := _m.On("GetOutputAliases", matchers...)
	return &NodeBuilder_GetOutputAliases{Call: c}
}

// GetOutputAliases provides a mock function with given fields:
func (_m *NodeBuilder) GetOutputAliases() []*core.Alias {
	ret := _m.Called()

	var r0 []*core.Alias
	if rf, ok := ret.Get(0).(func() []*core.Alias); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Alias)
		}
	}

	return r0
}

type NodeBuilder_GetSubWorkflow struct {
	*mock.Call
}

func (_m NodeBuilder_GetSubWorkflow) Return(_a0 common.Workflow) *NodeBuilder_GetSubWorkflow {
	return &NodeBuilder_GetSubWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetSubWorkflow() *NodeBuilder_GetSubWorkflow {
	c := _m.On("GetSubWorkflow")
	return &NodeBuilder_GetSubWorkflow{Call: c}
}

func (_m *NodeBuilder) OnGetSubWorkflowMatch(matchers ...interface{}) *NodeBuilder_GetSubWorkflow {
	c := _m.On("GetSubWorkflow", matchers...)
	return &NodeBuilder_GetSubWorkflow{Call: c}
}

// GetSubWorkflow provides a mock function with given fields:
func (_m *NodeBuilder) GetSubWorkflow() common.Workflow {
	ret := _m.Called()

	var r0 common.Workflow
	if rf, ok := ret.Get(0).(func() common.Workflow); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Workflow)
		}
	}

	return r0
}

type NodeBuilder_GetTask struct {
	*mock.Call
}

func (_m NodeBuilder_GetTask) Return(_a0 common.Task) *NodeBuilder_GetTask {
	return &NodeBuilder_GetTask{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetTask() *NodeBuilder_GetTask {
	c := _m.On("GetTask")
	return &NodeBuilder_GetTask{Call: c}
}

func (_m *NodeBuilder) OnGetTaskMatch(matchers ...interface{}) *NodeBuilder_GetTask {
	c := _m.On("GetTask", matchers...)
	return &NodeBuilder_GetTask{Call: c}
}

// GetTask provides a mock function with given fields:
func (_m *NodeBuilder) GetTask() common.Task {
	ret := _m.Called()

	var r0 common.Task
	if rf, ok := ret.Get(0).(func() common.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Task)
		}
	}

	return r0
}

type NodeBuilder_GetTaskNode struct {
	*mock.Call
}

func (_m NodeBuilder_GetTaskNode) Return(_a0 *core.TaskNode) *NodeBuilder_GetTaskNode {
	return &NodeBuilder_GetTaskNode{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetTaskNode() *NodeBuilder_GetTaskNode {
	c := _m.On("GetTaskNode")
	return &NodeBuilder_GetTaskNode{Call: c}
}

func (_m *NodeBuilder) OnGetTaskNodeMatch(matchers ...interface{}) *NodeBuilder_GetTaskNode {
	c := _m.On("GetTaskNode", matchers...)
	return &NodeBuilder_GetTaskNode{Call: c}
}

// GetTaskNode provides a mock function with given fields:
func (_m *NodeBuilder) GetTaskNode() *core.TaskNode {
	ret := _m.Called()

	var r0 *core.TaskNode
	if rf, ok := ret.Get(0).(func() *core.TaskNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskNode)
		}
	}

	return r0
}

type NodeBuilder_GetUpstreamNodeIds struct {
	*mock.Call
}

func (_m NodeBuilder_GetUpstreamNodeIds) Return(_a0 []string) *NodeBuilder_GetUpstreamNodeIds {
	return &NodeBuilder_GetUpstreamNodeIds{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetUpstreamNodeIds() *NodeBuilder_GetUpstreamNodeIds {
	c := _m.On("GetUpstreamNodeIds")
	return &NodeBuilder_GetUpstreamNodeIds{Call: c}
}

func (_m *NodeBuilder) OnGetUpstreamNodeIdsMatch(matchers ...interface{}) *NodeBuilder_GetUpstreamNodeIds {
	c := _m.On("GetUpstreamNodeIds", matchers...)
	return &NodeBuilder_GetUpstreamNodeIds{Call: c}
}

// GetUpstreamNodeIds provides a mock function with given fields:
func (_m *NodeBuilder) GetUpstreamNodeIds() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

type NodeBuilder_GetWorkflowNode struct {
	*mock.Call
}

func (_m NodeBuilder_GetWorkflowNode) Return(_a0 *core.WorkflowNode) *NodeBuilder_GetWorkflowNode {
	return &NodeBuilder_GetWorkflowNode{Call: _m.Call.Return(_a0)}
}

func (_m *NodeBuilder) OnGetWorkflowNode() *NodeBuilder_GetWorkflowNode {
	c := _m.On("GetWorkflowNode")
	return &NodeBuilder_GetWorkflowNode{Call: c}
}

func (_m *NodeBuilder) OnGetWorkflowNodeMatch(matchers ...interface{}) *NodeBuilder_GetWorkflowNode {
	c := _m.On("GetWorkflowNode", matchers...)
	return &NodeBuilder_GetWorkflowNode{Call: c}
}

// GetWorkflowNode provides a mock function with given fields:
func (_m *NodeBuilder) GetWorkflowNode() *core.WorkflowNode {
	ret := _m.Called()

	var r0 *core.WorkflowNode
	if rf, ok := ret.Get(0).(func() *core.WorkflowNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.WorkflowNode)
		}
	}

	return r0
}

// SetID provides a mock function with given fields: id
func (_m *NodeBuilder) SetID(id string) {
	_m.Called(id)
}

// SetInputs provides a mock function with given fields: inputs
func (_m *NodeBuilder) SetInputs(inputs []*core.Binding) {
	_m.Called(inputs)
}

// SetInterface provides a mock function with given fields: iface
func (_m *NodeBuilder) SetInterface(iface *core.TypedInterface) {
	_m.Called(iface)
}

// SetSubWorkflow provides a mock function with given fields: wf
func (_m *NodeBuilder) SetSubWorkflow(wf common.Workflow) {
	_m.Called(wf)
}

// SetTask provides a mock function with given fields: task
func (_m *NodeBuilder) SetTask(task common.Task) {
	_m.Called(task)
}
