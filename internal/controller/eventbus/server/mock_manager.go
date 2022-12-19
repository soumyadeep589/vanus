// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	vanus "github.com/linkall-labs/vanus/internal/primitive/vanus"
	segment "github.com/linkall-labs/vanus/proto/pkg/segment"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddServer mocks base method.
func (m *MockManager) AddServer(ctx context.Context, srv Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServer", ctx, srv)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddServer indicates an expected call of AddServer.
func (mr *MockManagerMockRecorder) AddServer(ctx, srv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServer", reflect.TypeOf((*MockManager)(nil).AddServer), ctx, srv)
}

// CanCreateEventbus mocks base method.
func (m *MockManager) CanCreateEventbus(ctx context.Context, replicaNum int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateEventbus", ctx, replicaNum)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanCreateEventbus indicates an expected call of CanCreateEventbus.
func (mr *MockManagerMockRecorder) CanCreateEventbus(ctx, replicaNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateEventbus", reflect.TypeOf((*MockManager)(nil).CanCreateEventbus), ctx, replicaNum)
}

// GetServerByAddress mocks base method.
func (m *MockManager) GetServerByAddress(addr string) Server {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerByAddress", addr)
	ret0, _ := ret[0].(Server)
	return ret0
}

// GetServerByAddress indicates an expected call of GetServerByAddress.
func (mr *MockManagerMockRecorder) GetServerByAddress(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerByAddress", reflect.TypeOf((*MockManager)(nil).GetServerByAddress), addr)
}

// GetServerByServerID mocks base method.
func (m *MockManager) GetServerByServerID(id vanus.ID) Server {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerByServerID", id)
	ret0, _ := ret[0].(Server)
	return ret0
}

// GetServerByServerID indicates an expected call of GetServerByServerID.
func (mr *MockManagerMockRecorder) GetServerByServerID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerByServerID", reflect.TypeOf((*MockManager)(nil).GetServerByServerID), id)
}

// RemoveServer mocks base method.
func (m *MockManager) RemoveServer(ctx context.Context, srv Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveServer", ctx, srv)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveServer indicates an expected call of RemoveServer.
func (mr *MockManagerMockRecorder) RemoveServer(ctx, srv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServer", reflect.TypeOf((*MockManager)(nil).RemoveServer), ctx, srv)
}

// Run mocks base method.
func (m *MockManager) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockManagerMockRecorder) Run(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockManager)(nil).Run), ctx)
}

// Stop mocks base method.
func (m *MockManager) Stop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", ctx)
}

// Stop indicates an expected call of Stop.
func (mr *MockManagerMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop), ctx)
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Address mocks base method.
func (m *MockServer) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address.
func (mr *MockServerMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockServer)(nil).Address))
}

// Close mocks base method.
func (m *MockServer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockServerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockServer)(nil).Close))
}

// GetClient mocks base method.
func (m *MockServer) GetClient() segment.SegmentServerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(segment.SegmentServerClient)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockServerMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockServer)(nil).GetClient))
}

// ID mocks base method.
func (m *MockServer) ID() vanus.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(vanus.ID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockServerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockServer)(nil).ID))
}

// IsActive mocks base method.
func (m *MockServer) IsActive(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActive indicates an expected call of IsActive.
func (mr *MockServerMockRecorder) IsActive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockServer)(nil).IsActive), ctx)
}

// Polish mocks base method.
func (m *MockServer) Polish() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Polish")
}

// Polish indicates an expected call of Polish.
func (mr *MockServerMockRecorder) Polish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Polish", reflect.TypeOf((*MockServer)(nil).Polish))
}

// RemoteStart mocks base method.
func (m *MockServer) RemoteStart(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteStart", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoteStart indicates an expected call of RemoteStart.
func (mr *MockServerMockRecorder) RemoteStart(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteStart", reflect.TypeOf((*MockServer)(nil).RemoteStart), ctx)
}

// RemoteStop mocks base method.
func (m *MockServer) RemoteStop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoteStop", ctx)
}

// RemoteStop indicates an expected call of RemoteStop.
func (mr *MockServerMockRecorder) RemoteStop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteStop", reflect.TypeOf((*MockServer)(nil).RemoteStop), ctx)
}

// Uptime mocks base method.
func (m *MockServer) Uptime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uptime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Uptime indicates an expected call of Uptime.
func (mr *MockServerMockRecorder) Uptime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uptime", reflect.TypeOf((*MockServer)(nil).Uptime))
}
