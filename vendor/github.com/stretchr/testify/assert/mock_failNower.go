package assert

import "github.com/stretchr/testify/mock"

type mockFailNower struct {
	mock.Mock
}

// FailNow provides a mock function with given fields:
func (_m *mockFailNower) FailNow() {
	_m.Called()
}
