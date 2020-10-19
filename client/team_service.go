package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type teamService struct {
	service
}

func newTeamService(sling *sling.Sling, uriTemplate string) *teamService {
	teamService := &teamService{}
	teamService.service = newService(serviceTeamService, sling, uriTemplate, new(model.Team))

	return teamService
}

func (s teamService) getPagedResponse(path string) ([]*model.Team, error) {
	resources := []*model.Team{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Teams), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Teams)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new team.
func (s teamService) Add(resource *model.Team) (*model.Team, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Team), nil
}

// GetByID returns the team that matches the input ID. If one cannot be found, it returns nil and an error.
func (s teamService) GetByID(id string) (*model.Team, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.Team), nil
}

// GetAll returns all teams. If none can be found or an error occurs, it returns an empty collection.
func (s teamService) GetAll() ([]*model.Team, error) {
	items := []*model.Team{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByPartialName performs a lookup and returns teams with a matching partial name.
func (s teamService) GetByPartialName(name string) ([]*model.Team, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.Team{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a team based on the one provided as input.
func (s teamService) Update(machinePolicy *model.Team) (*model.Team, error) {
	path, err := getUpdatePath(s, machinePolicy)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), machinePolicy, new(model.Team), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Team), nil
}
