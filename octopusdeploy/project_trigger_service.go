package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type projectTriggerService struct {
	canDeleteService
}

func newProjectTriggerService(sling *sling.Sling, uriTemplate string) *projectTriggerService {
	projectTriggerService := &projectTriggerService{}
	projectTriggerService.service = newService(serviceProjectTriggerService, sling, uriTemplate, new(ProjectTrigger))

	return projectTriggerService
}

func (s projectTriggerService) getPagedResponse(path string) ([]*ProjectTrigger, error) {
	resources := []*ProjectTrigger{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(ProjectTriggers), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*ProjectTriggers)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the project trigger that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s projectTriggerService) GetByID(id string) (*ProjectTrigger, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(ProjectTrigger), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*ProjectTrigger), nil
}

func (s projectTriggerService) GetByProjectID(id string) ([]*ProjectTrigger, error) {
	var triggersByProject []*ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, triggers...)

	return triggersByProject, nil
}

// GetAll returns all project triggers. If none can be found or an error
// occurs, it returns an empty collection.
func (s projectTriggerService) GetAll() ([]*ProjectTrigger, error) {
	path, err := getPath(s)
	if err != nil {
		return []*ProjectTrigger{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new project trigger.
func (s projectTriggerService) Add(resource *ProjectTrigger) (*ProjectTrigger, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

// Update modifies a project trigger based on the one provided as input.
func (s projectTriggerService) Update(resource ProjectTrigger) (*ProjectTrigger, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}
