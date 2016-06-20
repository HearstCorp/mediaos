package log5go

import "github.com/stretchr/testify/mock"

import "time"

type MockAppender struct {
	mock.Mock
}

func (m *MockAppender) Append(msg *[]byte, level LogLevel, tstamp time.Time) error {
	ret := m.Called(msg, level, tstamp)

	r0 := ret.Error(0)

	return r0
}
