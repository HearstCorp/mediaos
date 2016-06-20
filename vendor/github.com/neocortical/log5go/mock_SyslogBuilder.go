package log5go

import "github.com/stretchr/testify/mock"

type MockSyslogBuilder struct {
	mock.Mock
}

func (m *MockSyslogBuilder) ToLocalSyslog(facility SyslogPriority, tag string) Log5Go {
	ret := m.Called(facility, tag)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockSyslogBuilder) ToRemoteSyslog(facility SyslogPriority, tag string, transport string, addr string) Log5Go {
	ret := m.Called(facility, tag, transport, addr)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
