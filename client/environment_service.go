package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type EnvironmentService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewEnvironmentService(sling *sling.Sling) *EnvironmentService {
	if sling == nil {
		return nil
	}

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

	if isEmpty(id) {
		return nil, errors.New("EnvironmentService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

// GetAll returns all instances of an Environment.
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

// GetByName performs a lookup and returns the Environment with a matching name.
func (s *EnvironmentService) GetByName(name string) (*model.Environment, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("EnvironmentService: invalid parameter, name")
	}

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

// Add creates a new Environment.
func (s *EnvironmentService) Add(environment *model.Environment) (*model.Environment, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if environment == nil {
		return nil, errors.New("EnvironmentService: invalid parameter, environment")
	}

	err = environment.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, environment, new(model.Environment), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s *EnvironmentService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("EnvironmentService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *EnvironmentService) Update(environment *model.Environment) (*model.Environment, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = environment.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", environment.ID)
	resp, err := apiUpdate(s.sling, environment, new(model.Environment), path)

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
