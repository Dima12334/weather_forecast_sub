// Code generated by MockGen. DO NOT EDIT.
// Source: clients.go
//
// Generated by this command:
//
//	mockgen -source=clients.go -destination=mocks/mock_clients.go
//

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	reflect "reflect"
	clients "weather_forecast_sub/pkg/clients"

	gomock "go.uber.org/mock/gomock"
)

// MockWeatherClient is a mock of WeatherClient interface.
type MockWeatherClient struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherClientMockRecorder
	isgomock struct{}
}

// MockWeatherClientMockRecorder is the mock recorder for MockWeatherClient.
type MockWeatherClientMockRecorder struct {
	mock *MockWeatherClient
}

// NewMockWeatherClient creates a new mock instance.
func NewMockWeatherClient(ctrl *gomock.Controller) *MockWeatherClient {
	mock := &MockWeatherClient{ctrl: ctrl}
	mock.recorder = &MockWeatherClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherClient) EXPECT() *MockWeatherClientMockRecorder {
	return m.recorder
}

// GetAPICurrentWeather mocks base method.
func (m *MockWeatherClient) GetAPICurrentWeather(city string) (*clients.WeatherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPICurrentWeather", city)
	ret0, _ := ret[0].(*clients.WeatherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPICurrentWeather indicates an expected call of GetAPICurrentWeather.
func (mr *MockWeatherClientMockRecorder) GetAPICurrentWeather(city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPICurrentWeather", reflect.TypeOf((*MockWeatherClient)(nil).GetAPICurrentWeather), city)
}

// GetAPIDayWeather mocks base method.
func (m *MockWeatherClient) GetAPIDayWeather(city string) (*clients.DayWeatherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIDayWeather", city)
	ret0, _ := ret[0].(*clients.DayWeatherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIDayWeather indicates an expected call of GetAPIDayWeather.
func (mr *MockWeatherClientMockRecorder) GetAPIDayWeather(city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIDayWeather", reflect.TypeOf((*MockWeatherClient)(nil).GetAPIDayWeather), city)
}
