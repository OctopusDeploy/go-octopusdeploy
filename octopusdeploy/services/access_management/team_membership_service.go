package access_management

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
)

const teamMembershipV1BasePath = "teammembership"

type teamMembershipServiceV1 struct {
	client *services.AdminClient
	services.AdminService
}

func NewTeamMembershipService(client *services.AdminClient, previewTeamPath string) *teamMembershipServiceV1 {
	return &teamMembershipServiceV1{
		AdminService: services.NewAdminService(services.ServiceTeamMembershipService, teamMembershipV1BasePath, client),
	}
}
