package assert

import "github.com/stretchr/testify/mock"

type MockTestingT struct {
	mock.Mock
}

func (m *MockTestingT) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}
