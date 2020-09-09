package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewProjectService(sling *sling.Sling) *ProjectService {
	if sling == nil {
		return nil
	}

	return &ProjectService{
		sling: sling,
		path:  "projects",
	}
}

// Get returns a single project by its ID in Octopus Deploy
func (s *ProjectService) Get(id string) (*model.Project, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("ProjectService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// GetAll returns all projects in Octopus Deploy
func (s *ProjectService) GetAll() (*[]model.Project, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Project), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Project), nil
}

// GetByName gets an existing project by its project name in Octopus Deploy
func (s *ProjectService) GetByName(name string) (*model.Project, error) {
	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

// Add adds an new project in Octopus Deploy
func (s *ProjectService) Add(resource *model.Project) (*model.Project, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("ProjectService: invalid parameter, resource")
	}

	err = resource.Validate()
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Project), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// Delete deletes an existing project in Octopus Deploy
func (s *ProjectService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return errors.New("ProjectService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing project in Octopus Deploy
func (s *ProjectService) Update(resource *model.Project) (*model.Project, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("ProjectService: invalid parameter, resource")
	}

	err = resource.Validate()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

func (s *ProjectService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("ProjectService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("ProjectService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &ProjectService{}
