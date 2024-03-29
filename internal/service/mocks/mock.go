// Code generated by MockGen. DO NOT EDIT.
// Source: userStorage.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/andy-ahmedov/audit_log_server/pkg/domain"
	domain0 "github.com/andy-ahmedov/crud_service/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockAuditClient is a mock of AuditClient interface.
type MockAuditClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuditClientMockRecorder
}

// MockAuditClientMockRecorder is the mock recorder for MockAuditClient.
type MockAuditClientMockRecorder struct {
	mock *MockAuditClient
}

// NewMockAuditClient creates a new mock instance.
func NewMockAuditClient(ctrl *gomock.Controller) *MockAuditClient {
	mock := &MockAuditClient{ctrl: ctrl}
	mock.recorder = &MockAuditClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditClient) EXPECT() *MockAuditClientMockRecorder {
	return m.recorder
}

// SendLogRequest mocks base method.
func (m *MockAuditClient) SendLogRequest(ctx context.Context, req domain.LogItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendLogRequest", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendLogRequest indicates an expected call of SendLogRequest.
func (mr *MockAuditClientMockRecorder) SendLogRequest(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendLogRequest", reflect.TypeOf((*MockAuditClient)(nil).SendLogRequest), ctx, req)
}

// MockPasswordHasher is a mock of PasswordHasher interface.
type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher.
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance.
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockPasswordHasher) Hash(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockPasswordHasherMockRecorder) Hash(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockPasswordHasher)(nil).Hash), password)
}

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserStorage) CreateUser(ctx context.Context, inp domain0.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserStorageMockRecorder) CreateUser(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserStorage)(nil).CreateUser), ctx, inp)
}

// GetByCredential mocks base method.
func (m *MockUserStorage) GetByCredential(ctx context.Context, email, passwords string) (domain0.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCredential", ctx, email, passwords)
	ret0, _ := ret[0].(domain0.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCredential indicates an expected call of GetByCredential.
func (mr *MockUserStorageMockRecorder) GetByCredential(ctx, email, passwords interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCredential", reflect.TypeOf((*MockUserStorage)(nil).GetByCredential), ctx, email, passwords)
}

// MockSessionRepository is a mock of SessionRepository interface.
type MockSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryMockRecorder
}

// MockSessionRepositoryMockRecorder is the mock recorder for MockSessionRepository.
type MockSessionRepositoryMockRecorder struct {
	mock *MockSessionRepository
}

// NewMockSessionRepository creates a new mock instance.
func NewMockSessionRepository(ctrl *gomock.Controller) *MockSessionRepository {
	mock := &MockSessionRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRepository) EXPECT() *MockSessionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSessionRepository) Create(ctx context.Context, token domain0.RefreshSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSessionRepositoryMockRecorder) Create(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionRepository)(nil).Create), ctx, token)
}

// Get mocks base method.
func (m *MockSessionRepository) Get(ctx context.Context, token string) (domain0.RefreshSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, token)
	ret0, _ := ret[0].(domain0.RefreshSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionRepositoryMockRecorder) Get(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionRepository)(nil).Get), ctx, token)
}
