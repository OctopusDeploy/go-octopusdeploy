package teamV1

import (
	service2 "github.com/OctopusDeploy/go-octopusdeploy/service"
)

const teamMembershipBasePath = "teammembership"

type teamMembershipServiceV1 struct {
	client *service2.AdminClient
	service2.AdminService
}

func NewTeamMembershipService(client *service2.AdminClient, previewTeamPath string) *teamMembershipServiceV1 {
	return &teamMembershipServiceV1{
		AdminService: service2.NewAdminService(service2.ServiceTeamMembershipService, teamMembershipBasePath, client),
	}
}
