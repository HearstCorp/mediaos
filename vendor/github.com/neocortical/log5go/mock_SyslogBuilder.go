package log5go

import "github.com/stretchr/testify/mock"

type MockSyslogBuilder struct {
	mock.Mock
}

// ToLocalSyslog provides a mock function with given fields: facility, tag
func (_m *MockSyslogBuilder) ToLocalSyslog(facility SyslogPriority, tag string) Log5Go {
	ret := _m.Called(facility, tag)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(SyslogPriority, string) Log5Go); ok {
		r0 = rf(facility, tag)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// ToRemoteSyslog provides a mock function with given fields: facility, tag, transport, addr
func (_m *MockSyslogBuilder) ToRemoteSyslog(facility SyslogPriority, tag string, transport string, addr string) Log5Go {
	ret := _m.Called(facility, tag, transport, addr)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(SyslogPriority, string, string, string) Log5Go); ok {
		r0 = rf(facility, tag, transport, addr)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}
