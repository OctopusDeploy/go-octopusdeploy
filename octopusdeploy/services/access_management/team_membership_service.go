package access_management

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type teamMembershipService struct {
	previewTeamPath string

	octopusdeploy.service
}

func newTeamMembershipService(sling *sling.Sling, uriTemplate string, previewTeamPath string) *teamMembershipService {
	return &teamMembershipService{
		previewTeamPath: previewTeamPath,
		service:         octopusdeploy.newService(services.ServiceTeamMembershipService, sling, uriTemplate),
	}
}
