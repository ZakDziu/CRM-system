// Code generated by MockGen. DO NOT EDIT.
// Source: crm-system/pkg/authmiddleware (interfaces: AuthMiddleware)

// Package mockauthmiddleware is a generated GoMock package.
package mockauthmiddleware

import (
	authmiddleware "crm-system/pkg/authmiddleware"
	model "crm-system/pkg/model"
	http "net/http"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

// MockAuthMiddleware is a mock of AuthMiddleware interface.
type MockAuthMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMiddlewareMockRecorder
}

// MockAuthMiddlewareMockRecorder is the mock recorder for MockAuthMiddleware.
type MockAuthMiddlewareMockRecorder struct {
	mock *MockAuthMiddleware
}

// NewMockAuthMiddleware creates a new mock instance.
func NewMockAuthMiddleware(ctrl *gomock.Controller) *MockAuthMiddleware {
	mock := &MockAuthMiddleware{ctrl: ctrl}
	mock.recorder = &MockAuthMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthMiddleware) EXPECT() *MockAuthMiddlewareMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthMiddleware) Authorize(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Authorize", arg0)
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthMiddlewareMockRecorder) Authorize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthMiddleware)(nil).Authorize), arg0)
}

// CreateTokens mocks base method.
func (m *MockAuthMiddleware) CreateTokens(arg0 uuid.UUID, arg1 model.UserRole) (*authmiddleware.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTokens", arg0, arg1)
	ret0, _ := ret[0].(*authmiddleware.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTokens indicates an expected call of CreateTokens.
func (mr *MockAuthMiddlewareMockRecorder) CreateTokens(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTokens", reflect.TypeOf((*MockAuthMiddleware)(nil).CreateTokens), arg0, arg1)
}

// ExtractToken mocks base method.
func (m *MockAuthMiddleware) ExtractToken(arg0 *http.Request) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractToken", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtractToken indicates an expected call of ExtractToken.
func (mr *MockAuthMiddlewareMockRecorder) ExtractToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractToken", reflect.TypeOf((*MockAuthMiddleware)(nil).ExtractToken), arg0)
}

// GetUserID mocks base method.
func (m *MockAuthMiddleware) GetUserID(arg0 string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", arg0)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockAuthMiddlewareMockRecorder) GetUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockAuthMiddleware)(nil).GetUserID), arg0)
}

// GetUserRole mocks base method.
func (m *MockAuthMiddleware) GetUserRole(arg0 string) (model.UserRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", arg0)
	ret0, _ := ret[0].(model.UserRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockAuthMiddlewareMockRecorder) GetUserRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockAuthMiddleware)(nil).GetUserRole), arg0)
}

// Refresh mocks base method.
func (m *MockAuthMiddleware) Refresh(arg0 authmiddleware.Tokens) (*authmiddleware.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0)
	ret0, _ := ret[0].(*authmiddleware.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockAuthMiddlewareMockRecorder) Refresh(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockAuthMiddleware)(nil).Refresh), arg0)
}

// Validate mocks base method.
func (m *MockAuthMiddleware) Validate(arg0 string) (*authmiddleware.AccessClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(*authmiddleware.AccessClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockAuthMiddlewareMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAuthMiddleware)(nil).Validate), arg0)
}
