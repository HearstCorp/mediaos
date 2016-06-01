package log5go

import "github.com/stretchr/testify/mock"

import "time"

type MockFormatter struct {
	mock.Mock
}

// Format provides a mock function with given fields: tstamp, level, prefix, caller, line, msg, data, out
func (_m *MockFormatter) Format(tstamp time.Time, level LogLevel, prefix string, caller string, line uint, msg string, data Data, out *[]byte) {
	_m.Called(tstamp, level, prefix, caller, line, msg, data, out)
}

// SetTimeFormat provides a mock function with given fields: timeFormat
func (_m *MockFormatter) SetTimeFormat(timeFormat string) {
	_m.Called(timeFormat)
}

// SetLines provides a mock function with given fields: lines
func (_m *MockFormatter) SetLines(lines bool) {
	_m.Called(lines)
}
