package permissions

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type PermissionService struct {
	services.Service
}

func NewPermissionService(sling *sling.Sling, uriTemplate string) *PermissionService {
	return &PermissionService{
		Service: services.NewService(constants.ServicePermissionService, sling, uriTemplate),
	}
}
