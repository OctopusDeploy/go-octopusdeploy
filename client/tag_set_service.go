package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TagSetService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewTagSetService(sling *sling.Sling, uriTemplate string) *TagSetService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &TagSetService{
		sling: sling,
		path:  path,
	}
}

func (s *TagSetService) Get(id string) (*model.TagSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("TagSetService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

// GetAll returns all instances of a TagSet.
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

// GetByName performs a lookup and returns the TagSet with a matching name.
func (s *TagSetService) GetByName(name string) (*model.TagSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("TagSetService: invalid parameter, name")
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

// Add creates a new TagSet.
func (s *TagSetService) Add(tagSet *model.TagSet) (*model.TagSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if tagSet == nil {
		return nil, errors.New("TagSetService: invalid parameter, tagSet")
	}

	err = tagSet.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, tagSet, new(model.TagSet), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

func (s *TagSetService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("TagSetService: invalid parameter, id")
	}

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
