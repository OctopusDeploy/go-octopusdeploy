package machines

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
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

func (s *MachineRoleService) GetAll() ([]*string, error) {
	items := []*string{}
	path, err := services.GetPath(s)
	if err != nil {
		return nil, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}
