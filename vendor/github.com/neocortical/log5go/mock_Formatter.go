package log5go

import "github.com/stretchr/testify/mock"

import "time"

type MockFormatter struct {
	mock.Mock
}

func (m *MockFormatter) Format(tstamp time.Time, level LogLevel, prefix string, caller string, line uint, msg string, data Data, out *[]byte) {
	m.Called(tstamp, level, prefix, caller, line, msg, data, out)
}
func (m *MockFormatter) SetTimeFormat(timeFormat string) {
	m.Called(timeFormat)
}
func (m *MockFormatter) SetLines(lines bool) {
	m.Called(lines)
}
