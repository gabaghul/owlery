// Code generated by MockGen. DO NOT EDIT.
// Source: src/domain/emailing/services/interfaces.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"
	time "time"

	http "github.com/gabaghul/owlery/src/adapters/http"
	models "github.com/gabaghul/owlery/src/domain/emailing/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRedisAdapter is a mock of RedisAdapter interface.
type MockRedisAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockRedisAdapterMockRecorder
}

// MockRedisAdapterMockRecorder is the mock recorder for MockRedisAdapter.
type MockRedisAdapterMockRecorder struct {
	mock *MockRedisAdapter
}

// NewMockRedisAdapter creates a new mock instance.
func NewMockRedisAdapter(ctrl *gomock.Controller) *MockRedisAdapter {
	mock := &MockRedisAdapter{ctrl: ctrl}
	mock.recorder = &MockRedisAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisAdapter) EXPECT() *MockRedisAdapterMockRecorder {
	return m.recorder
}

// GetEmailingMemberListOffset mocks base method.
func (m *MockRedisAdapter) GetEmailingMemberListOffset(ctx context.Context, clientID int64, listID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailingMemberListOffset", ctx, clientID, listID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailingMemberListOffset indicates an expected call of GetEmailingMemberListOffset.
func (mr *MockRedisAdapterMockRecorder) GetEmailingMemberListOffset(ctx, clientID, listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailingMemberListOffset", reflect.TypeOf((*MockRedisAdapter)(nil).GetEmailingMemberListOffset), ctx, clientID, listID)
}

// Store mocks base method.
func (m *MockRedisAdapter) Store(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockRedisAdapterMockRecorder) Store(ctx, key, value, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRedisAdapter)(nil).Store), ctx, key, value, ttl)
}

// MockMailChimpAdapter is a mock of MailChimpAdapter interface.
type MockMailChimpAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockMailChimpAdapterMockRecorder
}

// MockMailChimpAdapterMockRecorder is the mock recorder for MockMailChimpAdapter.
type MockMailChimpAdapterMockRecorder struct {
	mock *MockMailChimpAdapter
}

// NewMockMailChimpAdapter creates a new mock instance.
func NewMockMailChimpAdapter(ctrl *gomock.Controller) *MockMailChimpAdapter {
	mock := &MockMailChimpAdapter{ctrl: ctrl}
	mock.recorder = &MockMailChimpAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailChimpAdapter) EXPECT() *MockMailChimpAdapterMockRecorder {
	return m.recorder
}

// GetContactsByListID mocks base method.
func (m *MockMailChimpAdapter) GetContactsByListID(ctx context.Context, listID string, offset, count int64) (http.GetContactsByListIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactsByListID", ctx, listID, offset, count)
	ret0, _ := ret[0].(http.GetContactsByListIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactsByListID indicates an expected call of GetContactsByListID.
func (mr *MockMailChimpAdapterMockRecorder) GetContactsByListID(ctx, listID, offset, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactsByListID", reflect.TypeOf((*MockMailChimpAdapter)(nil).GetContactsByListID), ctx, listID, offset, count)
}

// MockOmetriaAdapter is a mock of OmetriaAdapter interface.
type MockOmetriaAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockOmetriaAdapterMockRecorder
}

// MockOmetriaAdapterMockRecorder is the mock recorder for MockOmetriaAdapter.
type MockOmetriaAdapterMockRecorder struct {
	mock *MockOmetriaAdapter
}

// NewMockOmetriaAdapter creates a new mock instance.
func NewMockOmetriaAdapter(ctrl *gomock.Controller) *MockOmetriaAdapter {
	mock := &MockOmetriaAdapter{ctrl: ctrl}
	mock.recorder = &MockOmetriaAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOmetriaAdapter) EXPECT() *MockOmetriaAdapterMockRecorder {
	return m.recorder
}

// IngestContactRecords mocks base method.
func (m *MockOmetriaAdapter) IngestContactRecords(ctx context.Context, contacts []models.Contact) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestContactRecords", ctx, contacts)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngestContactRecords indicates an expected call of IngestContactRecords.
func (mr *MockOmetriaAdapterMockRecorder) IngestContactRecords(ctx, contacts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestContactRecords", reflect.TypeOf((*MockOmetriaAdapter)(nil).IngestContactRecords), ctx, contacts)
}

// MockPsqlAdapter is a mock of PsqlAdapter interface.
type MockPsqlAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockPsqlAdapterMockRecorder
}

// MockPsqlAdapterMockRecorder is the mock recorder for MockPsqlAdapter.
type MockPsqlAdapterMockRecorder struct {
	mock *MockPsqlAdapter
}

// NewMockPsqlAdapter creates a new mock instance.
func NewMockPsqlAdapter(ctrl *gomock.Controller) *MockPsqlAdapter {
	mock := &MockPsqlAdapter{ctrl: ctrl}
	mock.recorder = &MockPsqlAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPsqlAdapter) EXPECT() *MockPsqlAdapterMockRecorder {
	return m.recorder
}

// GetAllContactLists mocks base method.
func (m *MockPsqlAdapter) GetAllContactLists(ctx context.Context) ([]models.ContactList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllContactLists", ctx)
	ret0, _ := ret[0].([]models.ContactList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllContactLists indicates an expected call of GetAllContactLists.
func (mr *MockPsqlAdapterMockRecorder) GetAllContactLists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllContactLists", reflect.TypeOf((*MockPsqlAdapter)(nil).GetAllContactLists), ctx)
}

// GetAllEmailingConfigs mocks base method.
func (m *MockPsqlAdapter) GetAllEmailingConfigs(ctx context.Context) ([]models.EmailingConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmailingConfigs", ctx)
	ret0, _ := ret[0].([]models.EmailingConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEmailingConfigs indicates an expected call of GetAllEmailingConfigs.
func (mr *MockPsqlAdapterMockRecorder) GetAllEmailingConfigs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmailingConfigs", reflect.TypeOf((*MockPsqlAdapter)(nil).GetAllEmailingConfigs), ctx)
}

// GetContactListsByClientID mocks base method.
func (m *MockPsqlAdapter) GetContactListsByClientID(ctx context.Context, clientID int64) ([]models.ContactList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactListsByClientID", ctx, clientID)
	ret0, _ := ret[0].([]models.ContactList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactListsByClientID indicates an expected call of GetContactListsByClientID.
func (mr *MockPsqlAdapterMockRecorder) GetContactListsByClientID(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactListsByClientID", reflect.TypeOf((*MockPsqlAdapter)(nil).GetContactListsByClientID), ctx, clientID)
}

// GetEmailingConfigsByClientID mocks base method.
func (m *MockPsqlAdapter) GetEmailingConfigsByClientID(ctx context.Context, clientID int64) ([]models.EmailingConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailingConfigsByClientID", ctx, clientID)
	ret0, _ := ret[0].([]models.EmailingConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailingConfigsByClientID indicates an expected call of GetEmailingConfigsByClientID.
func (mr *MockPsqlAdapterMockRecorder) GetEmailingConfigsByClientID(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailingConfigsByClientID", reflect.TypeOf((*MockPsqlAdapter)(nil).GetEmailingConfigsByClientID), ctx, clientID)
}
