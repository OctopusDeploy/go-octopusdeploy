package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type EnvironmentService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewEnvironmentService(sling *sling.Sling, uriTemplate string) *EnvironmentService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &EnvironmentService{
		name:  "EnvironmentService",
		path:  path,
		sling: sling,
	}
}

func (s *EnvironmentService) Get(id string) (*model.Environment, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

// GetAll returns all instances of an Environment.
func (s *EnvironmentService) GetAll() ([]model.Environment, error) {
	err := s.validateInternalState()

	items := new([]model.Environment)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the Environment with a matching name.
func (s *EnvironmentService) GetByName(name string) (*model.Environment, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("GetByName", "name")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new Environment.
func (s *EnvironmentService) Add(environment *model.Environment) (*model.Environment, error) {
	if environment == nil {
		return nil, createInvalidParameterError("Add", "environment")
	}

	err := environment.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

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
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
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
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &EnvironmentService{}
