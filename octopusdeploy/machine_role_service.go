package octopusdeploy

import "github.com/dghubble/sling"

type machineRoleService struct {
	service
}

func newMachineRoleService(sling *sling.Sling, uriTemplate string) *machineRoleService {
	return &machineRoleService{
		service: newService(serviceMachineRoleService, sling, uriTemplate, nil),
	}
}
