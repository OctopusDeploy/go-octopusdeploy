package service

type SpaceScopedService struct {
	spaceScopedClient *SpaceScopedClient
	service
	IService
}

func NewSpaceScopedService(name string, basePathRelativeToRoot string, client *SpaceScopedClient) SpaceScopedService {
	return SpaceScopedService{
		service:           *NewService(name, basePathRelativeToRoot),
		spaceScopedClient: client,
	}
}

func (s SpaceScopedService) GetClient() IClient {
	return s.spaceScopedClient
}
