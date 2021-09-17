// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/sensor/common/compliance (interfaces: AuditLogCollectionManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	central "github.com/stackrox/rox/generated/internalapi/central"
	sensor "github.com/stackrox/rox/generated/internalapi/sensor"
	storage "github.com/stackrox/rox/generated/storage"
	centralsensor "github.com/stackrox/rox/pkg/centralsensor"
	reflect "reflect"
)

// MockAuditLogCollectionManager is a mock of AuditLogCollectionManager interface
type MockAuditLogCollectionManager struct {
	ctrl     *gomock.Controller
	recorder *MockAuditLogCollectionManagerMockRecorder
}

// MockAuditLogCollectionManagerMockRecorder is the mock recorder for MockAuditLogCollectionManager
type MockAuditLogCollectionManagerMockRecorder struct {
	mock *MockAuditLogCollectionManager
}

// NewMockAuditLogCollectionManager creates a new mock instance
func NewMockAuditLogCollectionManager(ctrl *gomock.Controller) *MockAuditLogCollectionManager {
	mock := &MockAuditLogCollectionManager{ctrl: ctrl}
	mock.recorder = &MockAuditLogCollectionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuditLogCollectionManager) EXPECT() *MockAuditLogCollectionManagerMockRecorder {
	return m.recorder
}

// AddEligibleComplianceNode mocks base method
func (m *MockAuditLogCollectionManager) AddEligibleComplianceNode(arg0 string, arg1 sensor.ComplianceService_CommunicateServer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEligibleComplianceNode", arg0, arg1)
}

// AddEligibleComplianceNode indicates an expected call of AddEligibleComplianceNode
func (mr *MockAuditLogCollectionManagerMockRecorder) AddEligibleComplianceNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEligibleComplianceNode", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).AddEligibleComplianceNode), arg0, arg1)
}

// AuditMessagesChan mocks base method
func (m *MockAuditLogCollectionManager) AuditMessagesChan() chan<- *sensor.MsgFromCompliance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuditMessagesChan")
	ret0, _ := ret[0].(chan<- *sensor.MsgFromCompliance)
	return ret0
}

// AuditMessagesChan indicates an expected call of AuditMessagesChan
func (mr *MockAuditLogCollectionManagerMockRecorder) AuditMessagesChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuditMessagesChan", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).AuditMessagesChan))
}

// Capabilities mocks base method
func (m *MockAuditLogCollectionManager) Capabilities() []centralsensor.SensorCapability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].([]centralsensor.SensorCapability)
	return ret0
}

// Capabilities indicates an expected call of Capabilities
func (mr *MockAuditLogCollectionManagerMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).Capabilities))
}

// DisableCollection mocks base method
func (m *MockAuditLogCollectionManager) DisableCollection() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisableCollection")
}

// DisableCollection indicates an expected call of DisableCollection
func (mr *MockAuditLogCollectionManagerMockRecorder) DisableCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableCollection", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).DisableCollection))
}

// EnableCollection mocks base method
func (m *MockAuditLogCollectionManager) EnableCollection() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableCollection")
}

// EnableCollection indicates an expected call of EnableCollection
func (mr *MockAuditLogCollectionManagerMockRecorder) EnableCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCollection", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).EnableCollection))
}

// ForceUpdate mocks base method
func (m *MockAuditLogCollectionManager) ForceUpdate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForceUpdate")
}

// ForceUpdate indicates an expected call of ForceUpdate
func (mr *MockAuditLogCollectionManagerMockRecorder) ForceUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUpdate", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).ForceUpdate))
}

// ProcessMessage mocks base method
func (m *MockAuditLogCollectionManager) ProcessMessage(arg0 *central.MsgToSensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage
func (mr *MockAuditLogCollectionManagerMockRecorder) ProcessMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).ProcessMessage), arg0)
}

// RemoveEligibleComplianceNode mocks base method
func (m *MockAuditLogCollectionManager) RemoveEligibleComplianceNode(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveEligibleComplianceNode", arg0)
}

// RemoveEligibleComplianceNode indicates an expected call of RemoveEligibleComplianceNode
func (mr *MockAuditLogCollectionManagerMockRecorder) RemoveEligibleComplianceNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEligibleComplianceNode", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).RemoveEligibleComplianceNode), arg0)
}

// ResponsesC mocks base method
func (m *MockAuditLogCollectionManager) ResponsesC() <-chan *central.MsgFromSensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponsesC")
	ret0, _ := ret[0].(<-chan *central.MsgFromSensor)
	return ret0
}

// ResponsesC indicates an expected call of ResponsesC
func (mr *MockAuditLogCollectionManagerMockRecorder) ResponsesC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponsesC", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).ResponsesC))
}

// SetAuditLogFileStateFromCentral mocks base method
func (m *MockAuditLogCollectionManager) SetAuditLogFileStateFromCentral(arg0 map[string]*storage.AuditLogFileState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAuditLogFileStateFromCentral", arg0)
}

// SetAuditLogFileStateFromCentral indicates an expected call of SetAuditLogFileStateFromCentral
func (mr *MockAuditLogCollectionManagerMockRecorder) SetAuditLogFileStateFromCentral(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAuditLogFileStateFromCentral", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).SetAuditLogFileStateFromCentral), arg0)
}

// Start mocks base method
func (m *MockAuditLogCollectionManager) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockAuditLogCollectionManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).Start))
}

// Stop mocks base method
func (m *MockAuditLogCollectionManager) Stop(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop
func (mr *MockAuditLogCollectionManagerMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAuditLogCollectionManager)(nil).Stop), arg0)
}