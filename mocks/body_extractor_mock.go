// Code generated by MockGen. DO NOT EDIT.
// Source: handlers/http/controllers/body_extractor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBodyExtractor is a mock of BodyExtractor interface.
type MockBodyExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockBodyExtractorMockRecorder
}

// MockBodyExtractorMockRecorder is the mock recorder for MockBodyExtractor.
type MockBodyExtractorMockRecorder struct {
	mock *MockBodyExtractor
}

// NewMockBodyExtractor creates a new mock instance.
func NewMockBodyExtractor(ctrl *gomock.Controller) *MockBodyExtractor {
	mock := &MockBodyExtractor{ctrl: ctrl}
	mock.recorder = &MockBodyExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBodyExtractor) EXPECT() *MockBodyExtractorMockRecorder {
	return m.recorder
}

// Extract mocks base method.
func (m *MockBodyExtractor) Extract(url string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extract", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Extract indicates an expected call of Extract.
func (mr *MockBodyExtractorMockRecorder) Extract(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extract", reflect.TypeOf((*MockBodyExtractor)(nil).Extract), url)
}