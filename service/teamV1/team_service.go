package teamV1

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/service"
)

const teamsBasePath = "teams"

type teamService struct {
	service.CanGetByIDService[Team]
	service.CanAddService[Team]
	service.CanUpdateService[Team]
	service.CanDeleteService[Team]
	service.IAdminService
}

type ITeamService interface {
	service.GetsByIDer[Team]
	// service.ResourceQueryer[Team]
	service.ResourceAdder[Team]
	service.ResourceUpdater[Team]
	service.DeleteByIDer[Team]
	GetScopedUserRoles(team Team) (service.IPagedResultsHandler[Team], error)
}

func NewTeamService(client service.IAdminClient) ITeamService {
	teamService := &teamService{
		IAdminService: service.NewAdminService(service.ServiceTeamService, teamsBasePath, client),
	}
	return teamService
}

func (s teamService) GetScopedUserRoles(team Team) (service.IPagedResultsHandler[Team], error) {
	//TODO: include skip/take params in the path
	path := fmt.Sprintf("%s/teams/%s/scopeduserroles", teamsBasePath, team.GetID())

	pageHandler := service.NewPagedResultsHandler[Team](s.GetClient(), 30, path)

	return pageHandler, nil
}