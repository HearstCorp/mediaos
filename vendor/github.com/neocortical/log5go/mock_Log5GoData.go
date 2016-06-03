package log5go

import "github.com/stretchr/testify/mock"

type MockLog5GoData struct {
	mock.Mock
}

// WithData provides a mock function with given fields: d
func (_m *MockLog5GoData) WithData(d Data) Log5Go {
	ret := _m.Called(d)

	var r0 Log5Go
	if rf, ok := ret.Get(0).(func(Data) Log5Go); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Get(0).(Log5Go)
	}

	return r0
}
