package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type invitationService struct {
	services.service
}

func newInvitationService(sling *sling.Sling, uriTemplate string) *invitationService {
	return &invitationService{
		service: services.newService(ServiceInvitationService, sling, uriTemplate),
	}
}
