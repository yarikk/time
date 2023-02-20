/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: time/ptp/sptp/client/stats.go

// Package client is a generated GoMock package.
package client

import (
	stats "github.com/facebook/time/ptp/sptp/stats"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStatsServer is a mock of StatsServer interface
type MockStatsServer struct {
	ctrl     *gomock.Controller
	recorder *MockStatsServerMockRecorder
}

// MockStatsServerMockRecorder is the mock recorder for MockStatsServer
type MockStatsServerMockRecorder struct {
	mock *MockStatsServer
}

// NewMockStatsServer creates a new mock instance
func NewMockStatsServer(ctrl *gomock.Controller) *MockStatsServer {
	mock := &MockStatsServer{ctrl: ctrl}
	mock.recorder = &MockStatsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatsServer) EXPECT() *MockStatsServerMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockStatsServer) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockStatsServerMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockStatsServer)(nil).Reset))
}

// SetCounter mocks base method
func (m *MockStatsServer) SetCounter(key string, val int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCounter", key, val)
}

// SetCounter indicates an expected call of SetCounter
func (mr *MockStatsServerMockRecorder) SetCounter(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCounter", reflect.TypeOf((*MockStatsServer)(nil).SetCounter), key, val)
}

// UpdateCounterBy mocks base method
func (m *MockStatsServer) UpdateCounterBy(key string, count int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateCounterBy", key, count)
}

// UpdateCounterBy indicates an expected call of UpdateCounterBy
func (mr *MockStatsServerMockRecorder) UpdateCounterBy(key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCounterBy", reflect.TypeOf((*MockStatsServer)(nil).UpdateCounterBy), key, count)
}

// SetGMStats mocks base method
func (m *MockStatsServer) SetGMStats(gm string, stats *stats.Stats) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGMStats", gm, stats)
}

// SetGMStats indicates an expected call of SetGMStats
func (mr *MockStatsServerMockRecorder) SetGMStats(gm, stats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGMStats", reflect.TypeOf((*MockStatsServer)(nil).SetGMStats), gm, stats)
}