package teamV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/service"
)

const teamMembershipBasePath = "teammembership"

type teamMembershipServiceV1 struct {
	*service.AdminClient
	service.IAdminService
}

func NewTeamMembershipService(client *service.AdminClient, previewTeamPath string) *teamMembershipServiceV1 {
	return &teamMembershipServiceV1{
		IAdminService: service.NewAdminService(service.ServiceTeamMembershipService, teamMembershipBasePath, client),
	}
}
