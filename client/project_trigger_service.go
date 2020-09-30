package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type projectTriggerService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newProjectTriggerService(sling *sling.Sling, uriTemplate string) *projectTriggerService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &projectTriggerService{
		name:        serviceProjectTriggerService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s projectTriggerService) getClient() *sling.Sling {
	return s.sling
}

func (s projectTriggerService) getName() string {
	return s.name
}

func (s projectTriggerService) getPagedResponse(path string) ([]model.ProjectTrigger, error) {
	resources := []model.ProjectTrigger{}
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

func (s projectTriggerService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
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
		return nil, createResourceNotFoundError("project trigger", "ID", id)
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s projectTriggerService) GetByProjectID(id string) (*[]model.ProjectTrigger, error) {
	var triggersByProject []model.ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, triggers...)

	return &triggersByProject, nil
}

// GetAll returns all project triggers. If none can be found or an error
// occurs, it returns an empty collection.
func (s projectTriggerService) GetAll() ([]model.ProjectTrigger, error) {
	path, err := getPath(s)
	if err != nil {
		return []model.ProjectTrigger{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new project trigger.
func (s projectTriggerService) Add(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

// DeleteByID deletes the project trigger that matches the input ID.
func (s projectTriggerService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("project trigger", "ID", id)
	}

	return err
}

// Update modifies a project trigger based on the one provided as input.
func (s projectTriggerService) Update(resource model.ProjectTrigger) (*model.ProjectTrigger, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

var _ ServiceInterface = &projectTriggerService{}
