package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type projectTriggerService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
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
		path:        strings.TrimSpace(uriTemplate),
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

func (s projectTriggerService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

func (s projectTriggerService) GetByID(id string) (*model.ProjectTrigger, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ProjectTrigger), path)
	if err != nil {
		return nil, err
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

// GetAll returns all instances of a ProjectTrigger. If none can be found or an error occurs, it returns an empty collection.
func (s projectTriggerService) GetAll() ([]model.ProjectTrigger, error) {
	err := validateInternalState(s)

	items := new([]model.ProjectTrigger)
	if err != nil {
		return *items, err
	}

	var p []model.ProjectTrigger
	path := trimTemplate(s.path)
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.ProjectTriggers), path)

		if err != nil {
			return *items, err
		}

		r := resp.(*model.ProjectTriggers)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// Add creates a new ProjectTrigger.
func (s projectTriggerService) Add(projectTrigger *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	if projectTrigger == nil {
		return nil, createInvalidParameterError(operationAdd, "projectTrigger")
	}

	err := projectTrigger.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), projectTrigger, new(model.ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s projectTriggerService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

func (s projectTriggerService) Update(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	if resource == nil {
		return nil, createInvalidParameterError(operationUpdate, "resource")
	}

	err := resource.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", resource.ID)

	resp, err := apiUpdate(s.getClient(), resource, new(model.ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

var _ ServiceInterface = &projectTriggerService{}
