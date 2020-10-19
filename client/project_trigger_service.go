package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type projectTriggerService struct {
	service
}

func newProjectTriggerService(sling *sling.Sling, uriTemplate string) *projectTriggerService {
	projectTriggerService := &projectTriggerService{}
	projectTriggerService.service = newService(serviceProjectTriggerService, sling, uriTemplate, new(model.ProjectTrigger))

	return projectTriggerService
}

func (s projectTriggerService) getPagedResponse(path string) ([]*model.ProjectTrigger, error) {
	resources := []*model.ProjectTrigger{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.ProjectTriggers), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.ProjectTriggers)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the project trigger that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s projectTriggerService) GetByID(id string) (*model.ProjectTrigger, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ProjectTrigger), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s projectTriggerService) GetByProjectID(id string) ([]*model.ProjectTrigger, error) {
	var triggersByProject []*model.ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, triggers...)

	return triggersByProject, nil
}

// GetAll returns all project triggers. If none can be found or an error
// occurs, it returns an empty collection.
func (s projectTriggerService) GetAll() ([]*model.ProjectTrigger, error) {
	path, err := getPath(s)
	if err != nil {
		return []*model.ProjectTrigger{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new project trigger.
func (s projectTriggerService) Add(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

// Update modifies a project trigger based on the one provided as input.
func (s projectTriggerService) Update(resource model.ProjectTrigger) (*model.ProjectTrigger, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}
