package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type SpaceService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewSpaceService(sling *sling.Sling, uriTemplate string) *SpaceService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &SpaceService{
		name:  "SpaceService",
		path:  path,
		sling: sling,
	}
}

func (s *SpaceService) Get(id string) (*model.Space, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

// GetAll returns all instances of a Space.
func (s *SpaceService) GetAll() ([]model.Space, error) {
	err := s.validateInternalState()

	items := new([]model.Space)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the Space with a matching name.
func (s *SpaceService) GetByName(name string) (*model.Space, error) {
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

// Add creates a new Space.
func (s *SpaceService) Add(space *model.Space) (*model.Space, error) {
	if space == nil {
		return nil, createInvalidParameterError("Add", "space")
	}

	err := space.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, space, new(model.Space), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

func (s *SpaceService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *SpaceService) Update(space *model.Space) (*model.Space, error) {
	if space == nil {
		return nil, createInvalidParameterError("Update", "space")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", space.ID)
	resp, err := apiUpdate(s.sling, space, new(model.Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

func (s *SpaceService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &SpaceService{}
