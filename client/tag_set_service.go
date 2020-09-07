package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TagSetService struct {
	sling *sling.Sling
	path  string
}

func NewTagSetService(sling *sling.Sling) *TagSetService {
	return &TagSetService{
		sling: sling,
		path:  "tagsets",
	}
}

func (s *TagSetService) Get(id string) (*model.TagSet, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("TagSetService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

func (s *TagSetService) GetAll() (*[]model.TagSet, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.TagSet), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.TagSet), nil
}

func (s *TagSetService) GetByName(name string) (*model.TagSet, error) {
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

func (s *TagSetService) Add(resource *model.TagSet) (*model.TagSet, error) {
	resp, err := apiAdd(s.sling, resource, new(model.TagSet), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

func (s *TagSetService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *TagSetService) Update(resource *model.TagSet) (*model.TagSet, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

func (s *TagSetService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("TagSetService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("TagSetService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &TagSetService{}
