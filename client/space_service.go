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

func NewSpaceService(sling *sling.Sling) *SpaceService {
	if sling == nil {
		return nil
	}

	return &SpaceService{
		sling: sling,
		path:  "spaces",
	}
}

func (s *SpaceService) Get(id string) (*model.Space, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("SpaceService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

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

func (s *SpaceService) GetByName(name string) (*model.Space, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(name, " ")) == 0 {
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

func (s *SpaceService) Add(space *model.Space) (*model.Space, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if space == nil {
		return nil, errors.New("SpaceService: invalid parameter, space")
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

	if len(strings.Trim(id, " ")) == 0 {
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
