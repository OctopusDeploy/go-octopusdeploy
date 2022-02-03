package octopusdeploy

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const teamsV1BasePath = "teams"

type teamService struct {
	adminService
	canDeleteService
}

func NewTeamService(client AdminClient) *teamService {
	teamService := &teamService{
		adminService: newAdminService(ServiceTeamService, client),
	}

	return teamService
}

func (s teamService) getPagedResponse(path string) ([]*Team, error) {
	resources := []*Team{}
	loadNextPage := true

	for loadNextPage {
		resp, err := s.client.apiGetPaged(new(Teams))
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Teams)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new team.
func (s teamService) Add(resource *Team) (*Team, error) {

	resp, err := s.client.apiAdd(resource, new(Team))
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

// Delete will delete a team if it is not a built-in team (i.e. the field,
// CanBeDeleted is true). If the team cannot be deleted or an error occurs, it
// returns an error.
func (s teamService) Delete(team *Team) error {
	if team == nil {
		return createInvalidParameterError(OperationDelete, ParameterTeam)
	}

	if !team.CanBeDeleted {
		return createBuiltInTeamsCannotDeleteError()
	}

	return s.client.apiDelete(team.GetID())
}

// Query returns a collection of teams based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s teamService) Query(teamsQuery TeamsQuery) (*Teams, error) {
	template, err := uritemplates.Parse(fmt.Sprintf("%s{?skip,take,ids,partialName,spaces,includeSystem}", s.BasePath))
	path, err := s.getURITemplate().Expand(teamsQuery)
	if err != nil {
		return &Teams{}, err
	}

	response, err := s.client.apiQuery(new(Teams), template)
	if err != nil {
		return &Teams{}, err
	}

	return response.(*Teams), nil
}

// GetAll returns all teams. If none can be found or an error occurs, it
// returns an empty collection.
func (s teamService) GetAll() ([]*Team, error) {
	items := []*Team{}
	path, err := s.getAllPath()
	if err != nil {
		return items, err
	}

	_, err = s.client.apiGet(&items, path)
	return items, err
}

// GetByID returns the team that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s teamService) GetByID(id string) (*Team, error) {
	resp, err := s.client.apiGetByID(new(Team), id)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

// GetByPartialName performs a lookup and returns teams with a matching partial
// name.
func (s teamService) GetByPartialName(name string) ([]*Team, error) {
	path, err := s.getByPartialNamePath(name)
	if err != nil {
		return []*Team{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a team based on the one provided as input.
func (s teamService) Update(team *Team) (*Team, error) {
	path, err := s.getUpdatePath(team)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.apiUpdate(team, new(Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

func (s teamService) GetScopedUserRoles(team Team, query SkipTakeQuery) (*ScopedUserRoles, error) {
	template, _ := uritemplates.Parse(team.Links["ScopedUserRoles"])
	path, _ := template.Expand(query)

	resp, err := s.client.apiGet(new(ScopedUserRoles), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ScopedUserRoles), nil
}
