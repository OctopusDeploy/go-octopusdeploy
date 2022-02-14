package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

type SpaceScopedService struct {
	spaceScopedClient *octopusdeploy.SpaceScopedClient
	service
	IService
}

func NewSpaceScopedService(name string, basePathRelativeToRoot string, client *octopusdeploy.SpaceScopedClient) SpaceScopedService {
	return SpaceScopedService{
		service:           *NewService(name, basePathRelativeToRoot),
		spaceScopedClient: client,
	}
}

func (s SpaceScopedService) GetClient() octopusdeploy.IClient {
	return s.spaceScopedClient
}
