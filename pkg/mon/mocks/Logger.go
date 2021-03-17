// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import mon "github.com/applike/gosoline/pkg/mon"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: args
func (_m *Logger) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: format, args
func (_m *Logger) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: err, msg
func (_m *Logger) Error(err error, msg string) {
	_m.Called(err, msg)
}

// Errorf provides a mock function with given fields: err, format, args
func (_m *Logger) Errorf(err error, format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: args
func (_m *Logger) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: format, args
func (_m *Logger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: args
func (_m *Logger) Warn(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warnf provides a mock function with given fields: format, args
func (_m *Logger) Warnf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithChannel provides a mock function with given fields: channel
func (_m *Logger) WithChannel(channel string) mon.Logger {
	ret := _m.Called(channel)

	var r0 mon.Logger
	if rf, ok := ret.Get(0).(func(string) mon.Logger); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mon.Logger)
		}
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Logger) WithContext(ctx context.Context) mon.Logger {
	ret := _m.Called(ctx)

	var r0 mon.Logger
	if rf, ok := ret.Get(0).(func(context.Context) mon.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mon.Logger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: fields
func (_m *Logger) WithFields(fields mon.Fields) mon.Logger {
	ret := _m.Called(fields)

	var r0 mon.Logger
	if rf, ok := ret.Get(0).(func(mon.Fields) mon.Logger); ok {
		r0 = rf(fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mon.Logger)
		}
	}

	return r0
}
