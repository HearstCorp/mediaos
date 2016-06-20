package log5go

import "github.com/stretchr/testify/mock"

type MockLog5Go struct {
	mock.Mock
}

func (m *MockLog5Go) Log(level LogLevel, format string, a ...interface{}) {
	m.Called(level, format, a)
}
func (m *MockLog5Go) Trace(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Debug(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Info(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Notice(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Warn(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Error(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Critical(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Alert(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) Fatal(format string, a ...interface{}) {
	m.Called(format, a)
}
func (m *MockLog5Go) LogLevel() LogLevel {
	ret := m.Called()

	r0 := ret.Get(0).(LogLevel)

	return r0
}
func (m *MockLog5Go) SetLogLevel(level LogLevel) {
	m.Called(level)
}
