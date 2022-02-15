package teamV1

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/service"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const teamsBasePath = "teams"

type teamService struct {
	service.AdminService
	service.CanGetByIDService[Team]
	service.CanAddService[Team]
	service.CanUpdateService[Team]
	service.CanDeleteService[Team]
	ITeamService
}

type ITeamService interface {
	service.GetsByIDer[Team]
	service.ResourceQueryer[Team]
	service.ResourceAdder[Team]
	service.ResourceUpdater[Team]
	service.DeleteByIDer[Team]
	GetScopedUserRoles(team Team) (service.IPagedResultsHandler[Team], error)
}

func NewTeamService(client *service.AdminClient) ITeamService {
	teamService := &teamService{
		AdminService: service.NewAdminService(service.ServiceTeamService, teamsBasePath, client),
	}

	return teamService
}

// Query returns a collection of teams based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s teamService) Query(teamsQuery service.TeamsQuery) (service.IPagedResultsHandler[Team], error) {
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

func (s teamService) GetScopedUserRoles(team Team) (service.IPagedResultsHandler[Team], error) {
	//TODO: include skip/take params in the path
	path := fmt.Sprintf("%s/teams/%s/scopeduserroles", teamsBasePath, team.GetID())

	pageHandler := service.NewPagedResultsHandler(s.GetClient(), 30, path)

	return pageHandler, nil
}
