package permissions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type PermissionService struct {
	services.Service
}

func NewPermissionService(sling *sling.Sling, uriTemplate string) *PermissionService {
	return &PermissionService{
		Service: services.NewService(constants.ServicePermissionService, sling, uriTemplate),
	}
}
