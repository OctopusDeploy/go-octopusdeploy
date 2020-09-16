package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectGroupService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewProjectGroupService(sling *sling.Sling, uriTemplate string) *ProjectGroupService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ProjectGroupService{
		sling: sling,
		path:  path,
	}
}

func (s *ProjectGroupService) Get(id string) (*model.ProjectGroup, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("ProjectGroupService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

// GetAll returns all instances of a ProjectGroup.
func (s *ProjectGroupService) GetAll() (*[]model.ProjectGroup, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.ProjectGroup), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.ProjectGroup), nil
}

// Add creates a new ProjectGroup.
func (s *ProjectGroupService) Add(projectGroup *model.ProjectGroup) (*model.ProjectGroup, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if projectGroup == nil {
		return nil, errors.New("ProjectGroupService: invalid parameter, projectGroup")
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
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("ProjectGroupService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectGroupService) Update(projectGroup *model.ProjectGroup) (*model.ProjectGroup, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if projectGroup == nil {
		return nil, errors.New("ProjectGroupService: invalid parameter, projectGroup")
	}

	err = projectGroup.Validate()

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
		return fmt.Errorf("ProjectGroupService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("ProjectGroupService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &ProjectGroupService{}
