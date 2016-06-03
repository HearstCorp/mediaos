package log5go

import "github.com/stretchr/testify/mock"

import "io"

type MockLogBuilder struct {
	mock.Mock
}

// Clone provides a mock function with given fields:
func (_m *MockLogBuilder) Clone() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithTimeFmt provides a mock function with given fields: format
func (_m *MockLogBuilder) WithTimeFmt(format string) Log5Go {
	ret := _m.Called(format)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(string) Log5Go); ok {
		r0 = rf(format)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// ToStdout provides a mock function with given fields:
func (_m *MockLogBuilder) ToStdout() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// ToStderr provides a mock function with given fields:
func (_m *MockLogBuilder) ToStderr() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// ToWriter provides a mock function with given fields: out
func (_m *MockLogBuilder) ToWriter(out io.Writer) Log5Go {
	ret := _m.Called(out)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(io.Writer) Log5Go); ok {
		r0 = rf(out)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// ToFile provides a mock function with given fields: directory, filename
func (_m *MockLogBuilder) ToFile(directory string, filename string) Log5Go {
	ret := _m.Called(directory, filename)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(string, string) Log5Go); ok {
		r0 = rf(directory, filename)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// ToAppender provides a mock function with given fields: appender
func (_m *MockLogBuilder) ToAppender(appender Appender) Log5Go {
	ret := _m.Called(appender)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(Appender) Log5Go); ok {
		r0 = rf(appender)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithRotation provides a mock function with given fields: frequency, keepNLogs
func (_m *MockLogBuilder) WithRotation(frequency rollFrequency, keepNLogs int) Log5Go {
	ret := _m.Called(frequency, keepNLogs)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(rollFrequency, int) Log5Go); ok {
		r0 = rf(frequency, keepNLogs)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithStderr provides a mock function with given fields:
func (_m *MockLogBuilder) WithStderr() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithPrefix provides a mock function with given fields: prefix
func (_m *MockLogBuilder) WithPrefix(prefix string) Log5Go {
	ret := _m.Called(prefix)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(string) Log5Go); ok {
		r0 = rf(prefix)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithLongLines provides a mock function with given fields:
func (_m *MockLogBuilder) WithLongLines() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithShortLines provides a mock function with given fields:
func (_m *MockLogBuilder) WithShortLines() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// WithFmt provides a mock function with given fields: format
func (_m *MockLogBuilder) WithFmt(format string) Log5Go {
	ret := _m.Called(format)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(string) Log5Go); ok {
		r0 = rf(format)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// Json provides a mock function with given fields:
func (_m *MockLogBuilder) Json() Log5Go {
	ret := _m.Called()

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func() Log5Go); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}

// Register provides a mock function with given fields: key
func (_m *MockLogBuilder) Register(key string) Log5Go {
	ret := _m.Called(key)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(string) Log5Go); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}
