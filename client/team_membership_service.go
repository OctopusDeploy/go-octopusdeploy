package client

import "github.com/dghubble/sling"

type teamMembershipService struct {
	previewTeamPath string

	service
}

func newTeamMembershipService(sling *sling.Sling, uriTemplate string, previewTeamPath string) *teamMembershipService {
	return &teamMembershipService{
		previewTeamPath: previewTeamPath,
		service:         newService(serviceTeamMembershipService, sling, uriTemplate, nil),
	}
}
