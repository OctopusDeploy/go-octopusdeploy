package teammembership

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
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
