package log5go

import "github.com/stretchr/testify/mock"

import "time"

type MockAppender struct {
	mock.Mock
}

// Append provides a mock function with given fields: msg, level, tstamp
func (_m *MockAppender) Append(msg *[]byte, level LogLevel, tstamp time.Time) error {
	ret := _m.Called(msg, level, tstamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(*[]byte, LogLevel, time.Time) error); ok {
		r0 = rf(msg, level, tstamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
