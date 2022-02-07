package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type machineRoleService struct {
	services.service
}

func newMachineRoleService(sling *sling.Sling, uriTemplate string) *machineRoleService {
	return &machineRoleService{
		service: services.newService(ServiceMachineRoleService, sling, uriTemplate),
	}
}
