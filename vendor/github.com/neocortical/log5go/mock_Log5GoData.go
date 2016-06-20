package log5go

import "github.com/stretchr/testify/mock"

type MockLog5GoData struct {
	mock.Mock
}

func (m *MockLog5GoData) WithData(d Data) Log5Go {
	ret := m.Called(d)

	r0 := ret.Get(0).(Log5Go)

	return r0
}
