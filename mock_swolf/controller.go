// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/noyako/swolf (interfaces: MasterController,TenantController)

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
