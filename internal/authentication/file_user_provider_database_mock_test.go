// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/authentication (interfaces: FileUserProviderDatabase)

// Package authentication is a generated GoMock package.
package authentication

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileUserDatabase is a mock of FileUserProviderDatabase interface.
type MockFileUserDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockFileUserDatabaseMockRecorder
}

// MockFileUserDatabaseMockRecorder is the mock recorder for MockFileUserDatabase.
type MockFileUserDatabaseMockRecorder struct {
	mock *MockFileUserDatabase
}

// NewMockFileUserDatabase creates a new mock instance.
func NewMockFileUserDatabase(ctrl *gomock.Controller) *MockFileUserDatabase {
	mock := &MockFileUserDatabase{ctrl: ctrl}
	mock.recorder = &MockFileUserDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileUserDatabase) EXPECT() *MockFileUserDatabaseMockRecorder {
	return m.recorder
}

// GetUserDetails mocks base method.
func (m *MockFileUserDatabase) GetUserDetails(arg0 string) (FileUserDatabaseUserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", arg0)
	ret0, _ := ret[0].(FileUserDatabaseUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockFileUserDatabaseMockRecorder) GetUserDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockFileUserDatabase)(nil).GetUserDetails), arg0)
}

// Load mocks base method.
func (m *MockFileUserDatabase) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockFileUserDatabaseMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockFileUserDatabase)(nil).Load))
}

// Save mocks base method.
func (m *MockFileUserDatabase) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockFileUserDatabaseMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileUserDatabase)(nil).Save))
}

// SetUserDetails mocks base method.
func (m *MockFileUserDatabase) SetUserDetails(arg0 string, arg1 *FileUserDatabaseUserDetails) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserDetails", arg0, arg1)
}

// SetUserDetails indicates an expected call of SetUserDetails.
func (mr *MockFileUserDatabaseMockRecorder) SetUserDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserDetails", reflect.TypeOf((*MockFileUserDatabase)(nil).SetUserDetails), arg0, arg1)
}
