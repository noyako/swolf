// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/noyako/swolf (interfaces: MasterController,TenantController,Pool)

// Package mock_swolf is a generated GoMock package.
package mock_swolf

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMasterController is a mock of MasterController interface
type MockMasterController struct {
	ctrl     *gomock.Controller
	recorder *MockMasterControllerMockRecorder
}

// MockMasterControllerMockRecorder is the mock recorder for MockMasterController
type MockMasterControllerMockRecorder struct {
	mock *MockMasterController
}

// NewMockMasterController creates a new mock instance
func NewMockMasterController(ctrl *gomock.Controller) *MockMasterController {
	mock := &MockMasterController{ctrl: ctrl}
	mock.recorder = &MockMasterControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMasterController) EXPECT() *MockMasterControllerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMasterController) Create(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMasterControllerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterController)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockMasterController) Delete(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockMasterControllerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterController)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockMasterController) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMasterControllerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMasterController)(nil).Get), arg0)
}

// MockTenantController is a mock of TenantController interface
type MockTenantController struct {
	ctrl     *gomock.Controller
	recorder *MockTenantControllerMockRecorder
}

// MockTenantControllerMockRecorder is the mock recorder for MockTenantController
type MockTenantControllerMockRecorder struct {
	mock *MockTenantController
}

// NewMockTenantController creates a new mock instance
func NewMockTenantController(ctrl *gomock.Controller) *MockTenantController {
	mock := &MockTenantController{ctrl: ctrl}
	mock.recorder = &MockTenantControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTenantController) EXPECT() *MockTenantControllerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTenantController) Create(arg0 string) (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTenantControllerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTenantController)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockTenantController) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTenantControllerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTenantController)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockTenantController) Get(arg0 string) (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTenantControllerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTenantController)(nil).Get), arg0)
}

// MockPool is a mock of Pool interface
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPool) Get(arg0 string) (*sql.DB, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPoolMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPool)(nil).Get), arg0)
}

// Put mocks base method
func (m *MockPool) Put(arg0 string, arg1 *sql.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put
func (mr *MockPoolMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPool)(nil).Put), arg0, arg1)
}

// SetMaxConn mocks base method
func (m *MockPool) SetMaxConn(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxConn", arg0)
}

// SetMaxConn indicates an expected call of SetMaxConn
func (mr *MockPoolMockRecorder) SetMaxConn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxConn", reflect.TypeOf((*MockPool)(nil).SetMaxConn), arg0)
}
