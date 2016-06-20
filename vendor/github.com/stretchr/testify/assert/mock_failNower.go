package assert

import "github.com/stretchr/testify/mock"

type MockfailNower struct {
	mock.Mock
}

func (m *MockfailNower) FailNow() {
	m.Called()
}
