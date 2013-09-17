// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: OutputRunner)

package testsupport

import (
	pipeline "github.com/mozilla-services/heka/pipeline"
	sync "sync"
	time "time"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of OutputRunner interface
type MockOutputRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockOutputRunnerRecorder
}

// Recorder for MockOutputRunner (not exported)
type _MockOutputRunnerRecorder struct {
	mock *MockOutputRunner
}

func NewMockOutputRunner(ctrl *gomock.Controller) *MockOutputRunner {
	mock := &MockOutputRunner{ctrl: ctrl}
	mock.recorder = &_MockOutputRunnerRecorder{mock}
	return mock
}

func (_m *MockOutputRunner) EXPECT() *_MockOutputRunnerRecorder {
	return _m.recorder
}

func (_m *MockOutputRunner) InChan() chan *pipeline.PipelinePack {
	ret := _m.ctrl.Call(_m, "InChan")
	ret0, _ := ret[0].(chan *pipeline.PipelinePack)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) InChan() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InChan")
}

func (_m *MockOutputRunner) LeakCount() int {
	ret := _m.ctrl.Call(_m, "LeakCount")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) LeakCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LeakCount")
}

func (_m *MockOutputRunner) LogError(_param0 error) {
	_m.ctrl.Call(_m, "LogError", _param0)
}

func (_mr *_MockOutputRunnerRecorder) LogError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogError", arg0)
}

func (_m *MockOutputRunner) LogMessage(_param0 string) {
	_m.ctrl.Call(_m, "LogMessage", _param0)
}

func (_mr *_MockOutputRunnerRecorder) LogMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogMessage", arg0)
}

func (_m *MockOutputRunner) MatchRunner() *pipeline.MatchRunner {
	ret := _m.ctrl.Call(_m, "MatchRunner")
	ret0, _ := ret[0].(*pipeline.MatchRunner)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) MatchRunner() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MatchRunner")
}

func (_m *MockOutputRunner) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockOutputRunner) Output() pipeline.Output {
	ret := _m.ctrl.Call(_m, "Output")
	ret0, _ := ret[0].(pipeline.Output)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Output() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Output")
}

func (_m *MockOutputRunner) Plugin() pipeline.Plugin {
	ret := _m.ctrl.Call(_m, "Plugin")
	ret0, _ := ret[0].(pipeline.Plugin)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Plugin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Plugin")
}

func (_m *MockOutputRunner) PluginGlobals() *pipeline.PluginGlobals {
	ret := _m.ctrl.Call(_m, "PluginGlobals")
	ret0, _ := ret[0].(*pipeline.PluginGlobals)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) PluginGlobals() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PluginGlobals")
}

func (_m *MockOutputRunner) RetainPack(_param0 *pipeline.PipelinePack) {
	_m.ctrl.Call(_m, "RetainPack", _param0)
}

func (_mr *_MockOutputRunnerRecorder) RetainPack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetainPack", arg0)
}

func (_m *MockOutputRunner) SetLeakCount(_param0 int) {
	_m.ctrl.Call(_m, "SetLeakCount", _param0)
}

func (_mr *_MockOutputRunnerRecorder) SetLeakCount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetLeakCount", arg0)
}

func (_m *MockOutputRunner) SetName(_param0 string) {
	_m.ctrl.Call(_m, "SetName", _param0)
}

func (_mr *_MockOutputRunnerRecorder) SetName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetName", arg0)
}

func (_m *MockOutputRunner) Start(_param0 pipeline.PluginHelper, _param1 *sync.WaitGroup) error {
	ret := _m.ctrl.Call(_m, "Start", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0, arg1)
}

func (_m *MockOutputRunner) Ticker() <-chan time.Time {
	ret := _m.ctrl.Call(_m, "Ticker")
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Ticker() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ticker")
}
