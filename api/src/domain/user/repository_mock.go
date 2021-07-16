// Code generated by MockGen. DO NOT EDIT.
// Source: src/domain/user/repository.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, id int) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, id)
}

// GetByIDs mocks base method.
func (m *MockRepository) GetByIDs(ctx context.Context, ids []int) ([]*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, ids)
	ret0, _ := ret[0].([]*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockRepositoryMockRecorder) GetByIDs(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockRepository)(nil).GetByIDs), ctx, ids)
}

// GetByUID mocks base method.
func (m *MockRepository) GetByUID(ctx context.Context, uid string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUID", ctx, uid)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUID indicates an expected call of GetByUID.
func (mr *MockRepositoryMockRecorder) GetByUID(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUID", reflect.TypeOf((*MockRepository)(nil).GetByUID), ctx, uid)
}

// GetByUIDs mocks base method.
func (m *MockRepository) GetByUIDs(ctx context.Context, uids []string) ([]*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUIDs", ctx, uids)
	ret0, _ := ret[0].([]*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUIDs indicates an expected call of GetByUIDs.
func (mr *MockRepositoryMockRecorder) GetByUIDs(ctx, uids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUIDs", reflect.TypeOf((*MockRepository)(nil).GetByUIDs), ctx, uids)
}

// GetOnlineUsers mocks base method.
func (m *MockRepository) GetOnlineUsers(ctx context.Context) ([]*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnlineUsers", ctx)
	ret0, _ := ret[0].([]*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOnlineUsers indicates an expected call of GetOnlineUsers.
func (mr *MockRepositoryMockRecorder) GetOnlineUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnlineUsers", reflect.TypeOf((*MockRepository)(nil).GetOnlineUsers), ctx)
}

// Offline mocks base method.
func (m *MockRepository) Offline(ctx context.Context, u *User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offline", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Offline indicates an expected call of Offline.
func (mr *MockRepositoryMockRecorder) Offline(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offline", reflect.TypeOf((*MockRepository)(nil).Offline), ctx, u)
}

// Online mocks base method.
func (m *MockRepository) Online(ctx context.Context, u *User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Online", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Online indicates an expected call of Online.
func (mr *MockRepositoryMockRecorder) Online(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Online", reflect.TypeOf((*MockRepository)(nil).Online), ctx, u)
}

// Save mocks base method.
func (m *MockRepository) Save(ctx context.Context, u *User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, u)
}
