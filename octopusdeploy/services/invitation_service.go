package services

import (
	"github.com/dghubble/sling"
)

type invitationService struct {
	service
}

func newInvitationService(sling *sling.Sling, uriTemplate string) *invitationService {
	return &invitationService{
		service: newService(ServiceInvitationService, sling, uriTemplate),
	}
}
