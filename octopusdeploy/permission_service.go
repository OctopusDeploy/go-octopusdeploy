package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type permissionService struct {
	services.service
}

func newPermissionService(sling *sling.Sling, uriTemplate string) *permissionService {
	return &permissionService{
		service: services.newService(ServicePermissionService, sling, uriTemplate),
	}
}
