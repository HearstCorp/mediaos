package mediaos

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetContent(publication PubData, key string, endpoint Endpoint, req Request) (ContentResponse, error) {
	ret := m.Called(publication, key, endpoint, req)

	r0 := ret.Get(0).(ContentResponse)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockClient) GetImages(publication PubData, key string, req Request) (ImageResponse, error) {
	ret := m.Called(publication, key, req)

	r0 := ret.Get(0).(ImageResponse)
	r1 := ret.Error(1)

	return r0, r1
}
