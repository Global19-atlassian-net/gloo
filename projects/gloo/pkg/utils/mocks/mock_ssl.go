// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/utils (interfaces: SslConfigTranslator)

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	utils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
)

// MockSslConfigTranslator is a mock of SslConfigTranslator interface
type MockSslConfigTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockSslConfigTranslatorMockRecorder
}

// MockSslConfigTranslatorMockRecorder is the mock recorder for MockSslConfigTranslator
type MockSslConfigTranslatorMockRecorder struct {
	mock *MockSslConfigTranslator
}

// NewMockSslConfigTranslator creates a new mock instance
func NewMockSslConfigTranslator(ctrl *gomock.Controller) *MockSslConfigTranslator {
	mock := &MockSslConfigTranslator{ctrl: ctrl}
	mock.recorder = &MockSslConfigTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSslConfigTranslator) EXPECT() *MockSslConfigTranslatorMockRecorder {
	return m.recorder
}

// ResolveCommonSslConfig mocks base method
func (m *MockSslConfigTranslator) ResolveCommonSslConfig(arg0 utils.CertSource, arg1 v1.SecretList, arg2 bool) (*envoy_api_v2_auth.CommonTlsContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveCommonSslConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(*envoy_api_v2_auth.CommonTlsContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveCommonSslConfig indicates an expected call of ResolveCommonSslConfig
func (mr *MockSslConfigTranslatorMockRecorder) ResolveCommonSslConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveCommonSslConfig", reflect.TypeOf((*MockSslConfigTranslator)(nil).ResolveCommonSslConfig), arg0, arg1, arg2)
}

// ResolveDownstreamSslConfig mocks base method
func (m *MockSslConfigTranslator) ResolveDownstreamSslConfig(arg0 v1.SecretList, arg1 *v1.SslConfig) (*envoy_api_v2_auth.DownstreamTlsContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveDownstreamSslConfig", arg0, arg1)
	ret0, _ := ret[0].(*envoy_api_v2_auth.DownstreamTlsContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveDownstreamSslConfig indicates an expected call of ResolveDownstreamSslConfig
func (mr *MockSslConfigTranslatorMockRecorder) ResolveDownstreamSslConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveDownstreamSslConfig", reflect.TypeOf((*MockSslConfigTranslator)(nil).ResolveDownstreamSslConfig), arg0, arg1)
}

// ResolveUpstreamSslConfig mocks base method
func (m *MockSslConfigTranslator) ResolveUpstreamSslConfig(arg0 v1.SecretList, arg1 *v1.UpstreamSslConfig) (*envoy_api_v2_auth.UpstreamTlsContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveUpstreamSslConfig", arg0, arg1)
	ret0, _ := ret[0].(*envoy_api_v2_auth.UpstreamTlsContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveUpstreamSslConfig indicates an expected call of ResolveUpstreamSslConfig
func (mr *MockSslConfigTranslatorMockRecorder) ResolveUpstreamSslConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveUpstreamSslConfig", reflect.TypeOf((*MockSslConfigTranslator)(nil).ResolveUpstreamSslConfig), arg0, arg1)
}
