// Code generated by MockGen. DO NOT EDIT.
// Source: user_account_repository.go
//
// Generated by this command:
//
//	mockgen -source=user_account_repository.go -destination=./mock/mock_user_account_repository.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
	model "mickamy.com/playground/internal/model"
	repository "mickamy.com/playground/internal/repository"
)

// MockUserAccount is a mock of UserAccount interface.
type MockUserAccount struct {
	ctrl     *gomock.Controller
	recorder *MockUserAccountMockRecorder
}

// MockUserAccountMockRecorder is the mock recorder for MockUserAccount.
type MockUserAccountMockRecorder struct {
	mock *MockUserAccount
}

// NewMockUserAccount creates a new mock instance.
func NewMockUserAccount(ctrl *gomock.Controller) *MockUserAccount {
	mock := &MockUserAccount{ctrl: ctrl}
	mock.recorder = &MockUserAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAccount) EXPECT() *MockUserAccountMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserAccount) Create(ctx context.Context, m *model.UserAccount) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserAccountMockRecorder) Create(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserAccount)(nil).Create), ctx, m)
}

// Get mocks base method.
func (m *MockUserAccount) Get(ctx context.Context, id model.UUID, scopes ...repository.Scope) (model.UserAccount, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(model.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserAccountMockRecorder) Get(ctx, id any, scopes ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserAccount)(nil).Get), varargs...)
}

// GetByIDToken mocks base method.
func (m *MockUserAccount) GetByIDToken(ctx context.Context, provider model.UserAccountProvider, uid string, scopes ...repository.Scope) (model.UserAccount, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, provider, uid}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByIDToken", varargs...)
	ret0, _ := ret[0].(model.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDToken indicates an expected call of GetByIDToken.
func (mr *MockUserAccountMockRecorder) GetByIDToken(ctx, provider, uid any, scopes ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, provider, uid}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDToken", reflect.TypeOf((*MockUserAccount)(nil).GetByIDToken), varargs...)
}

// GetBySlug mocks base method.
func (m *MockUserAccount) GetBySlug(ctx context.Context, slug string, scopes ...repository.Scope) (model.UserAccount, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, slug}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBySlug", varargs...)
	ret0, _ := ret[0].(model.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockUserAccountMockRecorder) GetBySlug(ctx, slug any, scopes ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, slug}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockUserAccount)(nil).GetBySlug), varargs...)
}

// WithTx mocks base method.
func (m *MockUserAccount) WithTx(tx *gorm.DB) repository.UserAccount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", tx)
	ret0, _ := ret[0].(repository.UserAccount)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockUserAccountMockRecorder) WithTx(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockUserAccount)(nil).WithTx), tx)
}
