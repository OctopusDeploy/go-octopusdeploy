package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// projectGroupService handles communication with ProjectGroup-related methods of the Octopus API.
type projectGroupService struct {
	service
}

// newProjectGroupService returns a projectGroupService with a preconfigured client.
func newProjectGroupService(sling *sling.Sling, uriTemplate string) *projectGroupService {
	projectGroupService := &projectGroupService{}
	projectGroupService.service = newService(serviceProjectGroupService, sling, uriTemplate, new(model.ProjectGroup))

	return projectGroupService
}

// Add creates a new project group.
func (s projectGroupService) Add(resource *model.ProjectGroup) (*model.ProjectGroup, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

// GetByID returns the project group that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s projectGroupService) GetByID(id string) (*model.ProjectGroup, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.ProjectGroup), nil
}

// GetAll returns all project groups. If none can be found or an error occurs,
// it returns an empty collection.
func (s projectGroupService) GetAll() ([]*model.ProjectGroup, error) {
	items := []*model.ProjectGroup{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// Update modifies a project group based on the one provided as input.
func (s projectGroupService) Update(resource model.ProjectGroup) (*model.ProjectGroup, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}
