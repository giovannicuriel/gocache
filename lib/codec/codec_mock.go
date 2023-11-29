// Code generated by MockGen. DO NOT EDIT.
// Source: lib/codec/interface.go

// Package codec is a generated GoMock package.
package codec

import (
	context "context"
	reflect "reflect"
	time "time"

	store "github.com/giovannicuriel/gocache/lib/v4/store"
	gomock "github.com/golang/mock/gomock"
)

// MockCodecInterface is a mock of CodecInterface interface.
type MockCodecInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCodecInterfaceMockRecorder
}

// MockCodecInterfaceMockRecorder is the mock recorder for MockCodecInterface.
type MockCodecInterfaceMockRecorder struct {
	mock *MockCodecInterface
}

// NewMockCodecInterface creates a new mock instance.
func NewMockCodecInterface(ctrl *gomock.Controller) *MockCodecInterface {
	mock := &MockCodecInterface{ctrl: ctrl}
	mock.recorder = &MockCodecInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodecInterface) EXPECT() *MockCodecInterfaceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockCodecInterface) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockCodecInterfaceMockRecorder) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockCodecInterface)(nil).Clear), ctx)
}

// Delete mocks base method.
func (m *MockCodecInterface) Delete(ctx context.Context, key any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCodecInterfaceMockRecorder) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCodecInterface)(nil).Delete), ctx, key)
}

// Get mocks base method.
func (m *MockCodecInterface) Get(ctx context.Context, key any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCodecInterfaceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCodecInterface)(nil).Get), ctx, key)
}

// GetStats mocks base method.
func (m *MockCodecInterface) GetStats() *Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(*Stats)
	return ret0
}

// GetStats indicates an expected call of GetStats.
func (mr *MockCodecInterfaceMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockCodecInterface)(nil).GetStats))
}

// GetStore mocks base method.
func (m *MockCodecInterface) GetStore() store.StoreInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore")
	ret0, _ := ret[0].(store.StoreInterface)
	return ret0
}

// GetStore indicates an expected call of GetStore.
func (mr *MockCodecInterfaceMockRecorder) GetStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockCodecInterface)(nil).GetStore))
}

// GetWithTTL mocks base method.
func (m *MockCodecInterface) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithTTL", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWithTTL indicates an expected call of GetWithTTL.
func (mr *MockCodecInterfaceMockRecorder) GetWithTTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTTL", reflect.TypeOf((*MockCodecInterface)(nil).GetWithTTL), ctx, key)
}

// Invalidate mocks base method.
func (m *MockCodecInterface) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Invalidate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockCodecInterfaceMockRecorder) Invalidate(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockCodecInterface)(nil).Invalidate), varargs...)
}

// Set mocks base method.
func (m *MockCodecInterface) Set(ctx context.Context, key, value any, options ...store.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, value}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCodecInterfaceMockRecorder) Set(ctx, key, value interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, value}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCodecInterface)(nil).Set), varargs...)
}
