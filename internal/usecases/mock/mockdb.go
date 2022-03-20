// Code generated by MockGen. DO NOT EDIT.
// Source: goCleanArch/internal/usecases (interfaces: UseCaseLogic)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	model "goCleanArch/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUseCaseLogic is a mock of UseCaseLogic interface.
type MockUseCaseLogic struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseLogicMockRecorder
}

// MockUseCaseLogicMockRecorder is the mock recorder for MockUseCaseLogic.
type MockUseCaseLogicMockRecorder struct {
	mock *MockUseCaseLogic
}

// NewMockUseCaseLogic creates a new mock instance.
func NewMockUseCaseLogic(ctrl *gomock.Controller) *MockUseCaseLogic {
	mock := &MockUseCaseLogic{ctrl: ctrl}
	mock.recorder = &MockUseCaseLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCaseLogic) EXPECT() *MockUseCaseLogicMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCaseLogic) Create(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseLogicMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCaseLogic)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockUseCaseLogic) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseLogicMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCaseLogic)(nil).Delete), arg0)
}

// Edit mocks base method.
func (m *MockUseCaseLogic) Edit(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockUseCaseLogicMockRecorder) Edit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockUseCaseLogic)(nil).Edit), arg0)
}

// FindById mocks base method.
func (m *MockUseCaseLogic) FindById(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUseCaseLogicMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUseCaseLogic)(nil).FindById), arg0)
}

// Validation mocks base method.
func (m *MockUseCaseLogic) Validation(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validation indicates an expected call of Validation.
func (mr *MockUseCaseLogicMockRecorder) Validation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validation", reflect.TypeOf((*MockUseCaseLogic)(nil).Validation), arg0)
}