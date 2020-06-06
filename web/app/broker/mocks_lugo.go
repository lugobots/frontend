// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lugobots/lugo4go/v2/lugo (interfaces: BroadcastClient,BroadcastServer,Broadcast_OnEventServer,Broadcast_OnEventClient)

// Package broker is a generated GoMock package.
package broker

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	lugo "github.com/lugobots/lugo4go/v2/lugo"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockBroadcastClient is a mock of BroadcastClient interface.
type MockBroadcastClient struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcastClientMockRecorder
}

// MockBroadcastClientMockRecorder is the mock recorder for MockBroadcastClient.
type MockBroadcastClientMockRecorder struct {
	mock *MockBroadcastClient
}

// NewMockBroadcastClient creates a new mock instance.
func NewMockBroadcastClient(ctrl *gomock.Controller) *MockBroadcastClient {
	mock := &MockBroadcastClient{ctrl: ctrl}
	mock.recorder = &MockBroadcastClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcastClient) EXPECT() *MockBroadcastClientMockRecorder {
	return m.recorder
}

// GetGameSetup mocks base method.
func (m *MockBroadcastClient) GetGameSetup(arg0 context.Context, arg1 *lugo.WatcherRequest, arg2 ...grpc.CallOption) (*lugo.GameSetup, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGameSetup", varargs...)
	ret0, _ := ret[0].(*lugo.GameSetup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameSetup indicates an expected call of GetGameSetup.
func (mr *MockBroadcastClientMockRecorder) GetGameSetup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameSetup", reflect.TypeOf((*MockBroadcastClient)(nil).GetGameSetup), varargs...)
}

// OnEvent mocks base method.
func (m *MockBroadcastClient) OnEvent(arg0 context.Context, arg1 *lugo.WatcherRequest, arg2 ...grpc.CallOption) (lugo.Broadcast_OnEventClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OnEvent", varargs...)
	ret0, _ := ret[0].(lugo.Broadcast_OnEventClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnEvent indicates an expected call of OnEvent.
func (mr *MockBroadcastClientMockRecorder) OnEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEvent", reflect.TypeOf((*MockBroadcastClient)(nil).OnEvent), varargs...)
}

// MockBroadcastServer is a mock of BroadcastServer interface.
type MockBroadcastServer struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcastServerMockRecorder
}

// MockBroadcastServerMockRecorder is the mock recorder for MockBroadcastServer.
type MockBroadcastServerMockRecorder struct {
	mock *MockBroadcastServer
}

// NewMockBroadcastServer creates a new mock instance.
func NewMockBroadcastServer(ctrl *gomock.Controller) *MockBroadcastServer {
	mock := &MockBroadcastServer{ctrl: ctrl}
	mock.recorder = &MockBroadcastServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcastServer) EXPECT() *MockBroadcastServerMockRecorder {
	return m.recorder
}

// GetGameSetup mocks base method.
func (m *MockBroadcastServer) GetGameSetup(arg0 context.Context, arg1 *lugo.WatcherRequest) (*lugo.GameSetup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameSetup", arg0, arg1)
	ret0, _ := ret[0].(*lugo.GameSetup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameSetup indicates an expected call of GetGameSetup.
func (mr *MockBroadcastServerMockRecorder) GetGameSetup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameSetup", reflect.TypeOf((*MockBroadcastServer)(nil).GetGameSetup), arg0, arg1)
}

// OnEvent mocks base method.
func (m *MockBroadcastServer) OnEvent(arg0 *lugo.WatcherRequest, arg1 lugo.Broadcast_OnEventServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnEvent indicates an expected call of OnEvent.
func (mr *MockBroadcastServerMockRecorder) OnEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEvent", reflect.TypeOf((*MockBroadcastServer)(nil).OnEvent), arg0, arg1)
}

// MockBroadcast_OnEventServer is a mock of Broadcast_OnEventServer interface.
type MockBroadcast_OnEventServer struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcast_OnEventServerMockRecorder
}

// MockBroadcast_OnEventServerMockRecorder is the mock recorder for MockBroadcast_OnEventServer.
type MockBroadcast_OnEventServerMockRecorder struct {
	mock *MockBroadcast_OnEventServer
}

// NewMockBroadcast_OnEventServer creates a new mock instance.
func NewMockBroadcast_OnEventServer(ctrl *gomock.Controller) *MockBroadcast_OnEventServer {
	mock := &MockBroadcast_OnEventServer{ctrl: ctrl}
	mock.recorder = &MockBroadcast_OnEventServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcast_OnEventServer) EXPECT() *MockBroadcast_OnEventServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBroadcast_OnEventServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBroadcast_OnEventServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockBroadcast_OnEventServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBroadcast_OnEventServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockBroadcast_OnEventServer) Send(arg0 *lugo.GameEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockBroadcast_OnEventServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockBroadcast_OnEventServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockBroadcast_OnEventServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockBroadcast_OnEventServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBroadcast_OnEventServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockBroadcast_OnEventServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockBroadcast_OnEventServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockBroadcast_OnEventServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockBroadcast_OnEventServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBroadcast_OnEventServer)(nil).SetTrailer), arg0)
}

// MockBroadcast_OnEventClient is a mock of Broadcast_OnEventClient interface.
type MockBroadcast_OnEventClient struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcast_OnEventClientMockRecorder
}

// MockBroadcast_OnEventClientMockRecorder is the mock recorder for MockBroadcast_OnEventClient.
type MockBroadcast_OnEventClientMockRecorder struct {
	mock *MockBroadcast_OnEventClient
}

// NewMockBroadcast_OnEventClient creates a new mock instance.
func NewMockBroadcast_OnEventClient(ctrl *gomock.Controller) *MockBroadcast_OnEventClient {
	mock := &MockBroadcast_OnEventClient{ctrl: ctrl}
	mock.recorder = &MockBroadcast_OnEventClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcast_OnEventClient) EXPECT() *MockBroadcast_OnEventClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockBroadcast_OnEventClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockBroadcast_OnEventClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockBroadcast_OnEventClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBroadcast_OnEventClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).Context))
}

// Header mocks base method.
func (m *MockBroadcast_OnEventClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockBroadcast_OnEventClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockBroadcast_OnEventClient) Recv() (*lugo.GameEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*lugo.GameEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockBroadcast_OnEventClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockBroadcast_OnEventClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBroadcast_OnEventClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockBroadcast_OnEventClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBroadcast_OnEventClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockBroadcast_OnEventClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockBroadcast_OnEventClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBroadcast_OnEventClient)(nil).Trailer))
}
