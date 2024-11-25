// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rattatatat3426/maseyth2 (interfaces: SendConn)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/rattatatat3426/maseyth2 -destination mock_send_conn_test.go github.com/rattatatat3426/maseyth2 SendConn
//
// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"

	protocol "github.com/rattatatat3426/maseyth2/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockSendConn is a mock of SendConn interface.
type MockSendConn struct {
	ctrl     *gomock.Controller
	recorder *MockSendConnMockRecorder
}

// MockSendConnMockRecorder is the mock recorder for MockSendConn.
type MockSendConnMockRecorder struct {
	mock *MockSendConn
}

// NewMockSendConn creates a new mock instance.
func NewMockSendConn(ctrl *gomock.Controller) *MockSendConn {
	mock := &MockSendConn{ctrl: ctrl}
	mock.recorder = &MockSendConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendConn) EXPECT() *MockSendConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSendConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSendConnMockRecorder) Close() *SendConnCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSendConn)(nil).Close))
	return &SendConnCloseCall{Call: call}
}

// SendConnCloseCall wrap *gomock.Call
type SendConnCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendConnCloseCall) Return(arg0 error) *SendConnCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendConnCloseCall) Do(f func() error) *SendConnCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendConnCloseCall) DoAndReturn(f func() error) *SendConnCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockSendConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockSendConnMockRecorder) LocalAddr() *SendConnLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockSendConn)(nil).LocalAddr))
	return &SendConnLocalAddrCall{Call: call}
}

// SendConnLocalAddrCall wrap *gomock.Call
type SendConnLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendConnLocalAddrCall) Return(arg0 net.Addr) *SendConnLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendConnLocalAddrCall) Do(f func() net.Addr) *SendConnLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendConnLocalAddrCall) DoAndReturn(f func() net.Addr) *SendConnLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoteAddr mocks base method.
func (m *MockSendConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockSendConnMockRecorder) RemoteAddr() *SendConnRemoteAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockSendConn)(nil).RemoteAddr))
	return &SendConnRemoteAddrCall{Call: call}
}

// SendConnRemoteAddrCall wrap *gomock.Call
type SendConnRemoteAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendConnRemoteAddrCall) Return(arg0 net.Addr) *SendConnRemoteAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendConnRemoteAddrCall) Do(f func() net.Addr) *SendConnRemoteAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendConnRemoteAddrCall) DoAndReturn(f func() net.Addr) *SendConnRemoteAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockSendConn) Write(arg0 []byte, arg1 uint16, arg2 protocol.ECN) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockSendConnMockRecorder) Write(arg0, arg1, arg2 any) *SendConnWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSendConn)(nil).Write), arg0, arg1, arg2)
	return &SendConnWriteCall{Call: call}
}

// SendConnWriteCall wrap *gomock.Call
type SendConnWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendConnWriteCall) Return(arg0 error) *SendConnWriteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendConnWriteCall) Do(f func([]byte, uint16, protocol.ECN) error) *SendConnWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendConnWriteCall) DoAndReturn(f func([]byte, uint16, protocol.ECN) error) *SendConnWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// capabilities mocks base method.
func (m *MockSendConn) capabilities() connCapabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "capabilities")
	ret0, _ := ret[0].(connCapabilities)
	return ret0
}

// capabilities indicates an expected call of capabilities.
func (mr *MockSendConnMockRecorder) capabilities() *SendConncapabilitiesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "capabilities", reflect.TypeOf((*MockSendConn)(nil).capabilities))
	return &SendConncapabilitiesCall{Call: call}
}

// SendConncapabilitiesCall wrap *gomock.Call
type SendConncapabilitiesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendConncapabilitiesCall) Return(arg0 connCapabilities) *SendConncapabilitiesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendConncapabilitiesCall) Do(f func() connCapabilities) *SendConncapabilitiesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendConncapabilitiesCall) DoAndReturn(f func() connCapabilities) *SendConncapabilitiesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
