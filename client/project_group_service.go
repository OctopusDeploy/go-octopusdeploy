package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectGroupService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewProjectGroupService(sling *sling.Sling, uriTemplate string) *ProjectGroupService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ProjectGroupService{
		name:  "ProjectGroupService",
		path:  path,
		sling: sling,
	}
}

func (s *ProjectGroupService) Get(id string) (*model.ProjectGroup, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

// GetAll returns all instances of a ProjectGroup.
func (s *ProjectGroupService) GetAll() ([]model.ProjectGroup, error) {
	err := s.validateInternalState()

	items := new([]model.ProjectGroup)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// Add creates a new ProjectGroup.
func (s *ProjectGroupService) Add(projectGroup *model.ProjectGroup) (*model.ProjectGroup, error) {
	if projectGroup == nil {
		return nil, createInvalidParameterError("Add", "projectGroup")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = projectGroup.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, projectGroup, new(model.ProjectGroup), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

func (s *ProjectGroupService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectGroupService) Update(projectGroup *model.ProjectGroup) (*model.ProjectGroup, error) {
	if projectGroup == nil {
		return nil, createInvalidParameterError("Update", "projectGroup")
	}

	err := projectGroup.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", projectGroup.ID)
	resp, err := apiUpdate(s.sling, projectGroup, new(model.ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

func (s *ProjectGroupService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ProjectGroupService{}
