// Code generated by MockGen. DO NOT EDIT.
// Source: src/rabbitmq/notification/service.go

// Package notification is a generated GoMock package.
package notification

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileNotification is a mock of FileNotification interface.
type MockFileNotification struct {
	ctrl     *gomock.Controller
	recorder *MockFileNotificationMockRecorder
}

// MockFileNotificationMockRecorder is the mock recorder for MockFileNotification.
type MockFileNotificationMockRecorder struct {
	mock *MockFileNotification
}

// NewMockFileNotification creates a new mock instance.
func NewMockFileNotification(ctrl *gomock.Controller) *MockFileNotification {
	mock := &MockFileNotification{ctrl: ctrl}
	mock.recorder = &MockFileNotificationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileNotification) EXPECT() *MockFileNotificationMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockFileNotification) Notify(ctx context.Context, model, serialNumber, fileName string, notification chan LzNotificationMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", ctx, model, serialNumber, fileName, notification)
}

// Notify indicates an expected call of Notify.
func (mr *MockFileNotificationMockRecorder) Notify(ctx, model, serialNumber, fileName, notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockFileNotification)(nil).Notify), ctx, model, serialNumber, fileName, notification)
}
