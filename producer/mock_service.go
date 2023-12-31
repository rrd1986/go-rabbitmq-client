// Code generated by MockGen. DO NOT EDIT.
// Source: src/rabbitmq/producer/service.go

// Package producer is a generated GoMock package.
package producer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProducerType is a mock of ProducerType interface.
type MockProducerType struct {
	ctrl     *gomock.Controller
	recorder *MockProducerTypeMockRecorder
}

// MockProducerTypeMockRecorder is the mock recorder for MockProducerType.
type MockProducerTypeMockRecorder struct {
	mock *MockProducerType
}

// NewMockProducerType creates a new mock instance.
func NewMockProducerType(ctrl *gomock.Controller) *MockProducerType {
	mock := &MockProducerType{ctrl: ctrl}
	mock.recorder = &MockProducerTypeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerType) EXPECT() *MockProducerTypeMockRecorder {
	return m.recorder
}

// CloseChannel mocks base method.
func (m *MockProducerType) CloseChannel(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseChannel", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseChannel indicates an expected call of CloseChannel.
func (mr *MockProducerTypeMockRecorder) CloseChannel(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseChannel", reflect.TypeOf((*MockProducerType)(nil).CloseChannel), ctx)
}

// PublishMessage mocks base method.
func (m *MockProducerType) PublishMessage(ctx context.Context, msg interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishMessage indicates an expected call of PublishMessage.
func (mr *MockProducerTypeMockRecorder) PublishMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMessage", reflect.TypeOf((*MockProducerType)(nil).PublishMessage), ctx, msg)
}
