package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectGroupService struct {
	sling *sling.Sling
	path  string
}

func NewProjectGroupService(sling *sling.Sling) *ProjectGroupService {
	return &ProjectGroupService{
		sling: sling,
		path:  "projectgroups",
	}
}

func (s *ProjectGroupService) Get(id string) (*model.ProjectGroup, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("ProjectGroupService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

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

func (s *ProjectGroupService) Add(resource *model.ProjectGroup) (*model.ProjectGroup, error) {
	resp, err := apiAdd(s.sling, resource, new(model.ProjectGroup), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

func (s *ProjectGroupService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectGroupService) Update(resource model.ProjectGroup) (*model.ProjectGroup, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.GetID())
	resp, err := apiUpdate(s.sling, resource, new(model.ProjectGroup), path)

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
