package client

type MockClient struct {
	BaseUrl string
}

func (c *MockClient) LoadRoutes() (*[]string, error, int) {
	return &[]string{"/route1", "/route2"}, nil, 200
}

func NewMockClient(url string) ClientType {
	return &MockClient{BaseUrl: url}
}
