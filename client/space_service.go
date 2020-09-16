package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type SpaceService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewSpaceService(sling *sling.Sling, uriTemplate string) *SpaceService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &SpaceService{
		sling: sling,
		path:  path,
	}
}

func (s *SpaceService) Get(id string) (*model.Space, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("SpaceService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

// GetAll returns all instances of a Space.
func (s *SpaceService) GetAll() (*[]model.Space, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Space), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Space), nil
}

// GetByName performs a lookup and returns the Space with a matching name.
func (s *SpaceService) GetByName(name string) (*model.Space, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("SpaceService: invalid parameter, name")
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

// Add creates a new Space.
func (s *SpaceService) Add(space *model.Space) (*model.Space, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if space == nil {
		return nil, errors.New("SpaceService: invalid parameter, space")
	}

	err = space.Validate()

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
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("SpaceService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *SpaceService) Update(space *model.Space) (*model.Space, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if space == nil {
		return nil, errors.New("SpaceService: invalid parameter, space")
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
		return fmt.Errorf("SpaceService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("SpaceService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &SpaceService{}
