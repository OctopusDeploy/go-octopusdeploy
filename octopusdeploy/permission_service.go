package octopusdeploy

import "github.com/dghubble/sling"

type permissionService struct {
	service
}

func newPermissionService(sling *sling.Sling, uriTemplate string) *permissionService {
	return &permissionService{
		service: newService(servicePermissionService, sling, uriTemplate, nil),
	}
}
