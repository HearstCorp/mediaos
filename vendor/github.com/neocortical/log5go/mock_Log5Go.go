package log5go

import "github.com/stretchr/testify/mock"

type MockLog5Go struct {
	mock.Mock
}

// Log provides a mock function with given fields: level, format, a
func (_m *MockLog5Go) Log(level LogLevel, format string, a ...interface{}) {
	_m.Called(level, format, a)
}

// Trace provides a mock function with given fields: format, a
func (_m *MockLog5Go) Trace(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Debug provides a mock function with given fields: format, a
func (_m *MockLog5Go) Debug(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Info provides a mock function with given fields: format, a
func (_m *MockLog5Go) Info(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Notice provides a mock function with given fields: format, a
func (_m *MockLog5Go) Notice(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Warn provides a mock function with given fields: format, a
func (_m *MockLog5Go) Warn(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Error provides a mock function with given fields: format, a
func (_m *MockLog5Go) Error(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Critical provides a mock function with given fields: format, a
func (_m *MockLog5Go) Critical(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Alert provides a mock function with given fields: format, a
func (_m *MockLog5Go) Alert(format string, a ...interface{}) {
	_m.Called(format, a)
}

// Fatal provides a mock function with given fields: format, a
func (_m *MockLog5Go) Fatal(format string, a ...interface{}) {
	_m.Called(format, a)
}

// LogLevel provides a mock function with given fields:
func (_m *MockLog5Go) LogLevel() LogLevel {
	ret := _m.Called()

	var r0 LogLevel
	if rf, ok := ret.Get(0).(func() LogLevel); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(LogLevel)
	}

	return r0
}

// SetLogLevel provides a mock function with given fields: level
func (_m *MockLog5Go) SetLogLevel(level LogLevel) {
	_m.Called(level)
}
