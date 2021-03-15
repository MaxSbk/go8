// Code generated by MockGen. DO NOT EDIT.
// Source: ../usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	book "github.com/gmhafiz/go8/internal/domain/book"
	models "github.com/gmhafiz/go8/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUseCase) Create(ctx context.Context, book *models.Book) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUseCaseMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, book)
}

// All mocks base method
func (m *MockUseCase) All(ctx context.Context, f *book.Filter) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, f)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockUseCaseMockRecorder) All(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockUseCase)(nil).All), ctx, f)
}

// Read mocks base method
func (m *MockUseCase) Read(ctx context.Context, bookID int64) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, bookID)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockUseCaseMockRecorder) Read(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockUseCase)(nil).Read), ctx, bookID)
}

// Update mocks base method
func (m *MockUseCase) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, book)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUseCaseMockRecorder) Update(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), ctx, book)
}

// Delete mocks base method
func (m *MockUseCase) Delete(ctx context.Context, bookID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUseCaseMockRecorder) Delete(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, bookID)
}

// Search mocks base method
func (m *MockUseCase) Search(ctx context.Context, req *book.Filter) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, req)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockUseCaseMockRecorder) Search(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUseCase)(nil).Search), ctx, req)
}
