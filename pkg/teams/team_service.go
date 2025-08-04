package teams

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/userroles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type TeamService struct {
	services.CanDeleteService
}

func NewTeamService(sling *sling.Sling, uriTemplate string) *TeamService {
	return &TeamService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceTeamService, sling, uriTemplate),
		},
	}
}

// Add creates a new team.
func (s *TeamService) Add(team *Team) (*Team, error) {
	if IsNil(team) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterTeam)
	}

	path, err := services.GetAddPath(s, team)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), team, new(Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

// Delete will delete a team if it is not a built-in team (i.e. the field,
// CanBeDeleted is true). If the team cannot be deleted or an error occurs, it
// returns an error.
func (s *TeamService) Delete(team *Team) error {
	if team == nil {
		return internal.CreateInvalidParameterError(constants.OperationDelete, constants.ParameterTeam)
	}

	if !team.CanBeDeleted {
		return internal.CreateBuiltInTeamsCannotDeleteError()
	}

	path := s.GetBasePath() + "/" + team.GetID()
	return services.ApiDelete(s.GetClient(), path)
}

// Get returns a collection of teams based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s *TeamService) Get(teamsQuery TeamsQuery) (*resources.Resources[*Team], error) {
	path, err := s.GetURITemplate().Expand(teamsQuery)
	if err != nil {
		return &resources.Resources[*Team]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Team]), path)
	if err != nil {
		return &resources.Resources[*Team]{}, err
	}

	return response.(*resources.Resources[*Team]), nil
}

// GetAll returns all teams. If none can be found or an error occurs, it
// returns an empty collection.
func (s *TeamService) GetAll() ([]*Team, error) {
	items := []*Team{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the team that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s *TeamService) GetByID(id string) (*Team, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

// GetByPartialName performs a lookup and returns teams with a matching partial
// name.
func (s *TeamService) GetByPartialName(partialName string) ([]*Team, error) {
	if internal.IsEmpty(partialName) {
		return []*Team{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Team{}, err
	}

	return services.GetPagedResponse[Team](s, path)
}

// Update modifies a team based on the one provided as input.
func (s *TeamService) Update(team *Team) (*Team, error) {
	path, err := services.GetUpdatePath(s, team)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), team, new(Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

func (s *TeamService) GetScopedUserRoles(team Team, query core.SkipTakeQuery) (*resources.Resources[*userroles.ScopedUserRole], error) {
	template, _ := uritemplates.Parse(team.Links["ScopedUserRoles"])
	path, _ := template.Expand(query)

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*userroles.ScopedUserRole]), path)
	if err != nil {
		return nil, err
	}

	return resp.(*resources.Resources[*userroles.ScopedUserRole]), nil
}

// --- new ---

const teamsTemplate = "/api/{spaceId}/teams{/id}{?skip,take,ids,partialName}"

func Add(client newclient.Client, team *Team) (*Team, error) {
	return newclient.Add[Team](client, teamsTemplate, team.SpaceID, team)
}

func GetByID(client newclient.Client, spaceId string, id string) (*Team, error) {
	return newclient.GetByID[Team](client, teamsTemplate, spaceId, id)
}

func DeleteByID(client newclient.Client, spaceId string, id string) error {
	return newclient.DeleteByID(client, teamsTemplate, spaceId, id)
}
