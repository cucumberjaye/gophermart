// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cucumberjaye/gophermart/internal/app/service (interfaces: MartRepository)

// Package mock_service is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/cucumberjaye/gophermart/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMartRepository is a mock of MartRepository interface.
type MockMartRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMartRepositoryMockRecorder
}

// MockMartRepositoryMockRecorder is the mock recorder for MockMartRepository.
type MockMartRepositoryMockRecorder struct {
	mock *MockMartRepository
}

// NewMockMartRepository creates a new mock instance.
func NewMockMartRepository(ctrl *gomock.Controller) *MockMartRepository {
	mock := &MockMartRepository{ctrl: ctrl}
	mock.recorder = &MockMartRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMartRepository) EXPECT() *MockMartRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockMartRepository) CreateUser(arg0 string, arg1 models.RegisterUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockMartRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockMartRepository)(nil).CreateUser), arg0, arg1)
}

// GetBalance mocks base method.
func (m *MockMartRepository) GetBalance(arg0 string) (models.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0)
	ret0, _ := ret[0].(models.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockMartRepositoryMockRecorder) GetBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockMartRepository)(nil).GetBalance), arg0)
}

// GetOrders mocks base method.
func (m *MockMartRepository) GetOrders(arg0 string) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockMartRepositoryMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockMartRepository)(nil).GetOrders), arg0)
}

// GetUser mocks base method.
func (m *MockMartRepository) GetUser(arg0 models.LoginUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockMartRepositoryMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockMartRepository)(nil).GetUser), arg0)
}

// GetWithdrawals mocks base method.
func (m *MockMartRepository) GetWithdrawals(arg0 string) ([]models.Withdraw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithdrawals", arg0)
	ret0, _ := ret[0].([]models.Withdraw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithdrawals indicates an expected call of GetWithdrawals.
func (mr *MockMartRepositoryMockRecorder) GetWithdrawals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithdrawals", reflect.TypeOf((*MockMartRepository)(nil).GetWithdrawals), arg0)
}

// SetOrder mocks base method.
func (m *MockMartRepository) SetOrder(arg0 models.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOrder indicates an expected call of SetOrder.
func (mr *MockMartRepositoryMockRecorder) SetOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrder", reflect.TypeOf((*MockMartRepository)(nil).SetOrder), arg0)
}

// Withdraw mocks base method.
func (m *MockMartRepository) Withdraw(arg0 string, arg1 models.Withdraw) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockMartRepositoryMockRecorder) Withdraw(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockMartRepository)(nil).Withdraw), arg0, arg1)
}
