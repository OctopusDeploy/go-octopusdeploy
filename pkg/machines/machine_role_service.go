package machines

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type MachineRoleService struct {
	services.Service
}

func NewMachineRoleService(sling *sling.Sling, uriTemplate string) *MachineRoleService {
	return &MachineRoleService{
		Service: services.NewService(constants.ServiceMachineRoleService, sling, uriTemplate),
	}
}
