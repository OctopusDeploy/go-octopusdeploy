package services

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/features/configuration/access_management/resources"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/access_management"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const teamsV1BasePath = "teams"

type teamServiceV1 struct {
	services.AdminService
	services.GetsByIDer[resources.Team]
	services.ResourceQueryer[resources.Team]
	services.CanAddService[resources.Team]
	services.CanUpdateService[resources.Team]
	services.CanDeleteService[resources.Team]
}

func NewTeamServiceV1(client *services.AdminClient) *teamServiceV1 {
	teamService := &teamServiceV1{
		AdminService: services.NewAdminService(services.ServiceTeamService, teamsV1BasePath, client),
	}

	return teamService
}

// Query returns a collection of teams based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s teamServiceV1) Query(teamsQuery services.TeamsQuery) (*access_management.Teams, error) {
	template, err := uritemplates.Parse(fmt.Sprintf("%s{?skip,take,ids,partialName,spaces,includeSystem}", s.BasePath))
	path, err := s.getURITemplate().Expand(teamsQuery)
	if err != nil {
		return &access_management.Teams{}, err
	}

	response, err := s.client.apiQuery(new(access_management.Teams), template)
	if err != nil {
		return &access_management.Teams{}, err
	}

	return response.(*access_management.Teams), nil
}

func (s teamServiceV1) GetScopedUserRoles(team resources.Team) (services.IPagedResultsHandler[resources.Team], error) {
	//TODO: include skip/take params in the path
	path := fmt.Sprintf("%s/teams/%s/scopeduserroles", teamsV1BasePath, team.GetID())

	pageHandler := services.NewPagedResultsHandler(s.GetClient(), 30, path)

	return pageHandler, nil
}
