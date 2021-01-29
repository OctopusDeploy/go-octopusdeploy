package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type teamService struct {
	canDeleteService
}

func newTeamService(sling *sling.Sling, uriTemplate string) *teamService {
	teamService := &teamService{}
	teamService.service = newService(ServiceTeamService, sling, uriTemplate)

	return teamService
}

func (s teamService) getPagedResponse(path string) ([]*Team, error) {
	resources := []*Team{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Teams), path)
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
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(Team), path)
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

	path := s.getBasePath() + "/" + team.GetID()
	return apiDelete(s.getClient(), path)
}

// Get returns a collection of teams based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s teamService) Get(teamsQuery TeamsQuery) (*Teams, error) {
	path, err := s.getURITemplate().Expand(teamsQuery)
	if err != nil {
		return &Teams{}, err
	}

	response, err := apiGet(s.getClient(), new(Teams), path)
	if err != nil {
		return &Teams{}, err
	}

	return response.(*Teams), nil
}

// GetAll returns all teams. If none can be found or an error occurs, it
// returns an empty collection.
func (s teamService) GetAll() ([]*Team, error) {
	items := []*Team{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the team that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s teamService) GetByID(id string) (*Team, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}

// GetByPartialName performs a lookup and returns teams with a matching partial
// name.
func (s teamService) GetByPartialName(name string) ([]*Team, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Team{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a team based on the one provided as input.
func (s teamService) Update(machinePolicy *Team) (*Team, error) {
	path, err := getUpdatePath(s, machinePolicy)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), machinePolicy, new(Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Team), nil
}
