package assert

import "github.com/stretchr/testify/mock"

type MockTestingT struct {
	mock.Mock
}

// Errorf provides a mock function with given fields: format, args
func (_m *MockTestingT) Errorf(format string, args ...interface{}) {
	_m.Called(format, args)
}
