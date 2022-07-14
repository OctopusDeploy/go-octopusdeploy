package teammembership

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type TeamMembershipService struct {
	previewTeamPath string

	services.Service
}

func NewTeamMembershipService(sling *sling.Sling, uriTemplate string, previewTeamPath string) *TeamMembershipService {
	return &TeamMembershipService{
		previewTeamPath: previewTeamPath,
		Service:         services.NewService(constants.ServiceTeamMembershipService, sling, uriTemplate),
	}
}
