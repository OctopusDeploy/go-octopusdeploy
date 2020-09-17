package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TagSetService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewTagSetService(sling *sling.Sling, uriTemplate string) *TagSetService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &TagSetService{
		name:  "TagSetService",
		path:  path,
		sling: sling,
	}
}

func (s *TagSetService) Get(id string) (*model.TagSet, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

// GetAll returns all instances of a TagSet.
func (s *TagSetService) GetAll() ([]model.TagSet, error) {
	err := s.validateInternalState()

	items := new([]model.TagSet)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the TagSet with a matching name.
func (s *TagSetService) GetByName(name string) (*model.TagSet, error) {
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

// Add creates a new TagSet.
func (s *TagSetService) Add(tagSet *model.TagSet) (*model.TagSet, error) {
	if tagSet == nil {
		return nil, createInvalidParameterError("Add", "tagSet")
	}

	err := tagSet.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

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
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
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
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &TagSetService{}
