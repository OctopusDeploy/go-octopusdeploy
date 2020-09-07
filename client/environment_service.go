package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type EnvironmentService struct {
	sling *sling.Sling
	path  string
}

func NewEnvironmentService(sling *sling.Sling) *EnvironmentService {
	return &EnvironmentService{
		sling: sling,
		path:  "environments",
	}
}

func (s *EnvironmentService) Get(id string) (*model.Environment, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("EnvironmentService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s *EnvironmentService) GetAll() (*[]model.Environment, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Environment), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Environment), nil
}

func (s *EnvironmentService) GetByName(name string) (*model.Environment, error) {
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

func (s *EnvironmentService) Add(resource *model.Environment) (*model.Environment, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Environment), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s *EnvironmentService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *EnvironmentService) Update(resource *model.Environment) (*model.Environment, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s *EnvironmentService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("EnvironmentService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("EnvironmentService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &EnvironmentService{}
