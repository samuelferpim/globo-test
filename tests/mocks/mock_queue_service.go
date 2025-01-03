// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/queue.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/ports/queue.go -destination=tests/mocks/mock_queue_service.go -package=mocks QueueService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	ports "bbb-voting/internal/core/ports"
	context "context"
	reflect "reflect"

	amqp "github.com/streadway/amqp"
	gomock "go.uber.org/mock/gomock"
)

// MockQueueService is a mock of QueueService interface.
type MockQueueService struct {
	ctrl     *gomock.Controller
	recorder *MockQueueServiceMockRecorder
	isgomock struct{}
}

// MockQueueServiceMockRecorder is the mock recorder for MockQueueService.
type MockQueueServiceMockRecorder struct {
	mock *MockQueueService
}

// NewMockQueueService creates a new mock instance.
func NewMockQueueService(ctrl *gomock.Controller) *MockQueueService {
	mock := &MockQueueService{ctrl: ctrl}
	mock.recorder = &MockQueueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueService) EXPECT() *MockQueueServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockQueueService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockQueueServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQueueService)(nil).Close))
}

// Publish mocks base method.
func (m *MockQueueService) Publish(ctx context.Context, topic string, message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, topic, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockQueueServiceMockRecorder) Publish(ctx, topic, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockQueueService)(nil).Publish), ctx, topic, message)
}

// PublishError mocks base method.
func (m *MockQueueService) PublishError(ctx context.Context, originalMessage []byte, errorMsg string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishError", ctx, originalMessage, errorMsg)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishError indicates an expected call of PublishError.
func (mr *MockQueueServiceMockRecorder) PublishError(ctx, originalMessage, errorMsg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishError", reflect.TypeOf((*MockQueueService)(nil).PublishError), ctx, originalMessage, errorMsg)
}

// MockAMQPConnection is a mock of AMQPConnection interface.
type MockAMQPConnection struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPConnectionMockRecorder
	isgomock struct{}
}

// MockAMQPConnectionMockRecorder is the mock recorder for MockAMQPConnection.
type MockAMQPConnectionMockRecorder struct {
	mock *MockAMQPConnection
}

// NewMockAMQPConnection creates a new mock instance.
func NewMockAMQPConnection(ctrl *gomock.Controller) *MockAMQPConnection {
	mock := &MockAMQPConnection{ctrl: ctrl}
	mock.recorder = &MockAMQPConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPConnection) EXPECT() *MockAMQPConnectionMockRecorder {
	return m.recorder
}

// Channel mocks base method.
func (m *MockAMQPConnection) Channel() (ports.AMQPChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Channel")
	ret0, _ := ret[0].(ports.AMQPChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Channel indicates an expected call of Channel.
func (mr *MockAMQPConnectionMockRecorder) Channel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Channel", reflect.TypeOf((*MockAMQPConnection)(nil).Channel))
}

// Close mocks base method.
func (m *MockAMQPConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPConnection)(nil).Close))
}

// MockAMQPChannel is a mock of AMQPChannel interface.
type MockAMQPChannel struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPChannelMockRecorder
	isgomock struct{}
}

// MockAMQPChannelMockRecorder is the mock recorder for MockAMQPChannel.
type MockAMQPChannelMockRecorder struct {
	mock *MockAMQPChannel
}

// NewMockAMQPChannel creates a new mock instance.
func NewMockAMQPChannel(ctrl *gomock.Controller) *MockAMQPChannel {
	mock := &MockAMQPChannel{ctrl: ctrl}
	mock.recorder = &MockAMQPChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPChannel) EXPECT() *MockAMQPChannelMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAMQPChannel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPChannel)(nil).Close))
}

// Publish mocks base method.
func (m *MockAMQPChannel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", exchange, key, mandatory, immediate, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockAMQPChannelMockRecorder) Publish(exchange, key, mandatory, immediate, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockAMQPChannel)(nil).Publish), exchange, key, mandatory, immediate, msg)
}
