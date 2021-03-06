// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/youtube/vitess/go/vt/discovery (interfaces: HealthCheck)

package txthrottler

import (
	gomock "github.com/golang/mock/gomock"
	discovery "github.com/youtube/vitess/go/vt/discovery"
	topodata "github.com/youtube/vitess/go/vt/proto/topodata"
	queryservice "github.com/youtube/vitess/go/vt/tabletserver/queryservice"
)

// Mock of HealthCheck interface
type MockHealthCheck struct {
	ctrl     *gomock.Controller
	recorder *_MockHealthCheckRecorder
}

// Recorder for MockHealthCheck (not exported)
type _MockHealthCheckRecorder struct {
	mock *MockHealthCheck
}

func NewMockHealthCheck(ctrl *gomock.Controller) *MockHealthCheck {
	mock := &MockHealthCheck{ctrl: ctrl}
	mock.recorder = &_MockHealthCheckRecorder{mock}
	return mock
}

func (_m *MockHealthCheck) EXPECT() *_MockHealthCheckRecorder {
	return _m.recorder
}

func (_m *MockHealthCheck) AddTablet(_param0 *topodata.Tablet, _param1 string) {
	_m.ctrl.Call(_m, "AddTablet", _param0, _param1)
}

func (_mr *_MockHealthCheckRecorder) AddTablet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTablet", arg0, arg1)
}

func (_m *MockHealthCheck) CacheStatus() discovery.TabletsCacheStatusList {
	ret := _m.ctrl.Call(_m, "CacheStatus")
	ret0, _ := ret[0].(discovery.TabletsCacheStatusList)
	return ret0
}

func (_mr *_MockHealthCheckRecorder) CacheStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CacheStatus")
}

func (_m *MockHealthCheck) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHealthCheckRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockHealthCheck) GetConnection(_param0 string) queryservice.QueryService {
	ret := _m.ctrl.Call(_m, "GetConnection", _param0)
	ret0, _ := ret[0].(queryservice.QueryService)
	return ret0
}

func (_mr *_MockHealthCheckRecorder) GetConnection(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConnection", arg0)
}

func (_m *MockHealthCheck) RegisterStats() {
	_m.ctrl.Call(_m, "RegisterStats")
}

func (_mr *_MockHealthCheckRecorder) RegisterStats() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterStats")
}

func (_m *MockHealthCheck) RemoveTablet(_param0 *topodata.Tablet) {
	_m.ctrl.Call(_m, "RemoveTablet", _param0)
}

func (_mr *_MockHealthCheckRecorder) RemoveTablet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveTablet", arg0)
}

func (_m *MockHealthCheck) SetListener(_param0 discovery.HealthCheckStatsListener, _param1 bool) {
	_m.ctrl.Call(_m, "SetListener", _param0, _param1)
}

func (_mr *_MockHealthCheckRecorder) SetListener(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetListener", arg0, arg1)
}

func (_m *MockHealthCheck) WaitForInitialStatsUpdates() {
	_m.ctrl.Call(_m, "WaitForInitialStatsUpdates")
}

func (_mr *_MockHealthCheckRecorder) WaitForInitialStatsUpdates() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitForInitialStatsUpdates")
}
