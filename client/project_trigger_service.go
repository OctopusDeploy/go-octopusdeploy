package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectTriggerService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewProjectTriggerService(sling *sling.Sling, uriTemplate string) *ProjectTriggerService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ProjectTriggerService{
		name:  "ProjectTriggerService",
		path:  path,
		sling: sling,
	}
}

func (s *ProjectTriggerService) Get(id string) (*model.ProjectTrigger, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) GetByProjectID(id string) (*[]model.ProjectTrigger, error) {
	var triggersByProject []model.ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, triggers...)

	return &triggersByProject, nil
}

// GetAll returns all instances of a ProjectTrigger.
func (s *ProjectTriggerService) GetAll() ([]model.ProjectTrigger, error) {
	err := s.validateInternalState()

	items := new([]model.ProjectTrigger)

	if err != nil {
		return *items, err
	}

	var p []model.ProjectTrigger
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.ProjectTriggers), path)

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
func (s *ProjectTriggerService) Add(projectTrigger *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	if projectTrigger == nil {
		return nil, createInvalidParameterError("Add", "projectTrigger")
	}

	err := projectTrigger.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, projectTrigger, new(model.ProjectTrigger), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectTriggerService) Update(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	if resource == nil {
		return nil, createInvalidParameterError("Update", "resource")
	}

	err := resource.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ProjectTriggerService{}
