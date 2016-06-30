package mediaos

import "github.com/stretchr/testify/mock"

type MockPubData struct {
	mock.Mock
}

func (m *MockPubData) Name() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) MosDomain() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) MosPort() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) MosDomainAndPort() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) RamsDomain() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) DisplayName() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) NotificationAlias() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockPubData) IUDomain() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
