package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

type AdminService struct {
	adminClient *octopusdeploy.AdminClient
	service
	Adminer
}

func NewAdminService(name string, basePathRelativeToRoot string, client *octopusdeploy.AdminClient) AdminService {
	return AdminService{
		service:     *NewService(name, basePathRelativeToRoot),
		adminClient: client,
	}
}

func (s AdminService) GetClient() octopusdeploy.IClient {
	return s.adminClient
}
