package log5go

import "github.com/stretchr/testify/mock"

import "io"

type MockLogBuilder struct {
	mock.Mock
}

func (m *MockLogBuilder) Clone() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithTimeFmt(format string) Log5Go {
	ret := m.Called(format)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) ToStdout() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) ToStderr() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) ToWriter(out io.Writer) Log5Go {
	ret := m.Called(out)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) ToFile(directory string, filename string) Log5Go {
	ret := m.Called(directory, filename)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) ToAppender(appender Appender) Log5Go {
	ret := m.Called(appender)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithRotation(frequency rollFrequency, keepNLogs int) Log5Go {
	ret := m.Called(frequency, keepNLogs)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithStderr() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithPrefix(prefix string) Log5Go {
	ret := m.Called(prefix)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithLongLines() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithShortLines() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) WithFmt(format string) Log5Go {
	ret := m.Called(format)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) Json() Log5Go {
	ret := m.Called()

	r0 := ret.Get(0).(Log5Go)

	return r0
}
func (m *MockLogBuilder) Register(key string) Log5Go {
	ret := m.Called(key)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
