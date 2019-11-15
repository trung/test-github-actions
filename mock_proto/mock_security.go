// Code generated by MockGen. DO NOT EDIT.
// Source: proto/security.pb.go

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTLSConfigurationSourceClient is a mock of TLSConfigurationSourceClient interface
type MockTLSConfigurationSourceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTLSConfigurationSourceClientMockRecorder
}

// MockTLSConfigurationSourceClientMockRecorder is the mock recorder for MockTLSConfigurationSourceClient
type MockTLSConfigurationSourceClientMockRecorder struct {
	mock *MockTLSConfigurationSourceClient
}

// NewMockTLSConfigurationSourceClient creates a new mock instance
func NewMockTLSConfigurationSourceClient(ctrl *gomock.Controller) *MockTLSConfigurationSourceClient {
	mock := &MockTLSConfigurationSourceClient{ctrl: ctrl}
	mock.recorder = &MockTLSConfigurationSourceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTLSConfigurationSourceClient) EXPECT() *MockTLSConfigurationSourceClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTLSConfigurationSourceClient) Get(ctx context.Context, in *proto.TLSConfiguration_Request, opts ...grpc.CallOption) (*proto.TLSConfiguration_Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*proto.TLSConfiguration_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTLSConfigurationSourceClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTLSConfigurationSourceClient)(nil).Get), varargs...)
}

// MockTLSConfigurationSourceServer is a mock of TLSConfigurationSourceServer interface
type MockTLSConfigurationSourceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTLSConfigurationSourceServerMockRecorder
}

// MockTLSConfigurationSourceServerMockRecorder is the mock recorder for MockTLSConfigurationSourceServer
type MockTLSConfigurationSourceServerMockRecorder struct {
	mock *MockTLSConfigurationSourceServer
}

// NewMockTLSConfigurationSourceServer creates a new mock instance
func NewMockTLSConfigurationSourceServer(ctrl *gomock.Controller) *MockTLSConfigurationSourceServer {
	mock := &MockTLSConfigurationSourceServer{ctrl: ctrl}
	mock.recorder = &MockTLSConfigurationSourceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTLSConfigurationSourceServer) EXPECT() *MockTLSConfigurationSourceServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTLSConfigurationSourceServer) Get(arg0 context.Context, arg1 *proto.TLSConfiguration_Request) (*proto.TLSConfiguration_Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*proto.TLSConfiguration_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTLSConfigurationSourceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTLSConfigurationSourceServer)(nil).Get), arg0, arg1)
}

// MockAuthenticationManagerClient is a mock of AuthenticationManagerClient interface
type MockAuthenticationManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationManagerClientMockRecorder
}

// MockAuthenticationManagerClientMockRecorder is the mock recorder for MockAuthenticationManagerClient
type MockAuthenticationManagerClientMockRecorder struct {
	mock *MockAuthenticationManagerClient
}

// NewMockAuthenticationManagerClient creates a new mock instance
func NewMockAuthenticationManagerClient(ctrl *gomock.Controller) *MockAuthenticationManagerClient {
	mock := &MockAuthenticationManagerClient{ctrl: ctrl}
	mock.recorder = &MockAuthenticationManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationManagerClient) EXPECT() *MockAuthenticationManagerClientMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockAuthenticationManagerClient) Authenticate(ctx context.Context, in *proto.PreAuthenticatedAuthenticationToken, opts ...grpc.CallOption) (*proto.PreAuthenticatedAuthenticationToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Authenticate", varargs...)
	ret0, _ := ret[0].(*proto.PreAuthenticatedAuthenticationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAuthenticationManagerClientMockRecorder) Authenticate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationManagerClient)(nil).Authenticate), varargs...)
}

// MockAuthenticationManagerServer is a mock of AuthenticationManagerServer interface
type MockAuthenticationManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationManagerServerMockRecorder
}

// MockAuthenticationManagerServerMockRecorder is the mock recorder for MockAuthenticationManagerServer
type MockAuthenticationManagerServerMockRecorder struct {
	mock *MockAuthenticationManagerServer
}

// NewMockAuthenticationManagerServer creates a new mock instance
func NewMockAuthenticationManagerServer(ctrl *gomock.Controller) *MockAuthenticationManagerServer {
	mock := &MockAuthenticationManagerServer{ctrl: ctrl}
	mock.recorder = &MockAuthenticationManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationManagerServer) EXPECT() *MockAuthenticationManagerServerMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockAuthenticationManagerServer) Authenticate(arg0 context.Context, arg1 *proto.PreAuthenticatedAuthenticationToken) (*proto.PreAuthenticatedAuthenticationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(*proto.PreAuthenticatedAuthenticationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAuthenticationManagerServerMockRecorder) Authenticate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationManagerServer)(nil).Authenticate), arg0, arg1)
}

// MockAccountAccessDecisionManagerClient is a mock of AccountAccessDecisionManagerClient interface
type MockAccountAccessDecisionManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountAccessDecisionManagerClientMockRecorder
}

// MockAccountAccessDecisionManagerClientMockRecorder is the mock recorder for MockAccountAccessDecisionManagerClient
type MockAccountAccessDecisionManagerClientMockRecorder struct {
	mock *MockAccountAccessDecisionManagerClient
}

// NewMockAccountAccessDecisionManagerClient creates a new mock instance
func NewMockAccountAccessDecisionManagerClient(ctrl *gomock.Controller) *MockAccountAccessDecisionManagerClient {
	mock := &MockAccountAccessDecisionManagerClient{ctrl: ctrl}
	mock.recorder = &MockAccountAccessDecisionManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountAccessDecisionManagerClient) EXPECT() *MockAccountAccessDecisionManagerClientMockRecorder {
	return m.recorder
}

// IsAuthorized mocks base method
func (m *MockAccountAccessDecisionManagerClient) IsAuthorized(ctx context.Context, in *proto.AccountAccess_Request, opts ...grpc.CallOption) (*proto.AccountAccess_Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsAuthorized", varargs...)
	ret0, _ := ret[0].(*proto.AccountAccess_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized
func (mr *MockAccountAccessDecisionManagerClientMockRecorder) IsAuthorized(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockAccountAccessDecisionManagerClient)(nil).IsAuthorized), varargs...)
}

// MockAccountAccessDecisionManagerServer is a mock of AccountAccessDecisionManagerServer interface
type MockAccountAccessDecisionManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountAccessDecisionManagerServerMockRecorder
}

// MockAccountAccessDecisionManagerServerMockRecorder is the mock recorder for MockAccountAccessDecisionManagerServer
type MockAccountAccessDecisionManagerServerMockRecorder struct {
	mock *MockAccountAccessDecisionManagerServer
}

// NewMockAccountAccessDecisionManagerServer creates a new mock instance
func NewMockAccountAccessDecisionManagerServer(ctrl *gomock.Controller) *MockAccountAccessDecisionManagerServer {
	mock := &MockAccountAccessDecisionManagerServer{ctrl: ctrl}
	mock.recorder = &MockAccountAccessDecisionManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountAccessDecisionManagerServer) EXPECT() *MockAccountAccessDecisionManagerServerMockRecorder {
	return m.recorder
}

// IsAuthorized mocks base method
func (m *MockAccountAccessDecisionManagerServer) IsAuthorized(arg0 context.Context, arg1 *proto.AccountAccess_Request) (*proto.AccountAccess_Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", arg0, arg1)
	ret0, _ := ret[0].(*proto.AccountAccess_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized
func (mr *MockAccountAccessDecisionManagerServerMockRecorder) IsAuthorized(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockAccountAccessDecisionManagerServer)(nil).IsAuthorized), arg0, arg1)
}

// MockContractAccessDecisionManagerClient is a mock of ContractAccessDecisionManagerClient interface
type MockContractAccessDecisionManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractAccessDecisionManagerClientMockRecorder
}

// MockContractAccessDecisionManagerClientMockRecorder is the mock recorder for MockContractAccessDecisionManagerClient
type MockContractAccessDecisionManagerClientMockRecorder struct {
	mock *MockContractAccessDecisionManagerClient
}

// NewMockContractAccessDecisionManagerClient creates a new mock instance
func NewMockContractAccessDecisionManagerClient(ctrl *gomock.Controller) *MockContractAccessDecisionManagerClient {
	mock := &MockContractAccessDecisionManagerClient{ctrl: ctrl}
	mock.recorder = &MockContractAccessDecisionManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractAccessDecisionManagerClient) EXPECT() *MockContractAccessDecisionManagerClientMockRecorder {
	return m.recorder
}

// IsAuthorized mocks base method
func (m *MockContractAccessDecisionManagerClient) IsAuthorized(ctx context.Context, in *proto.ContractAccess_Request, opts ...grpc.CallOption) (*proto.ContractAccess_Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsAuthorized", varargs...)
	ret0, _ := ret[0].(*proto.ContractAccess_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized
func (mr *MockContractAccessDecisionManagerClientMockRecorder) IsAuthorized(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockContractAccessDecisionManagerClient)(nil).IsAuthorized), varargs...)
}

// MockContractAccessDecisionManagerServer is a mock of ContractAccessDecisionManagerServer interface
type MockContractAccessDecisionManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockContractAccessDecisionManagerServerMockRecorder
}

// MockContractAccessDecisionManagerServerMockRecorder is the mock recorder for MockContractAccessDecisionManagerServer
type MockContractAccessDecisionManagerServerMockRecorder struct {
	mock *MockContractAccessDecisionManagerServer
}

// NewMockContractAccessDecisionManagerServer creates a new mock instance
func NewMockContractAccessDecisionManagerServer(ctrl *gomock.Controller) *MockContractAccessDecisionManagerServer {
	mock := &MockContractAccessDecisionManagerServer{ctrl: ctrl}
	mock.recorder = &MockContractAccessDecisionManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractAccessDecisionManagerServer) EXPECT() *MockContractAccessDecisionManagerServerMockRecorder {
	return m.recorder
}

// IsAuthorized mocks base method
func (m *MockContractAccessDecisionManagerServer) IsAuthorized(arg0 context.Context, arg1 *proto.ContractAccess_Request) (*proto.ContractAccess_Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", arg0, arg1)
	ret0, _ := ret[0].(*proto.ContractAccess_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized
func (mr *MockContractAccessDecisionManagerServerMockRecorder) IsAuthorized(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockContractAccessDecisionManagerServer)(nil).IsAuthorized), arg0, arg1)
}