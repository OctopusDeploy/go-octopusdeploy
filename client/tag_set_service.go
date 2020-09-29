package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type tagSetService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newTagSetService(sling *sling.Sling, uriTemplate string) *tagSetService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &tagSetService{
		name:        serviceTagSetService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s tagSetService) getClient() *sling.Sling {
	return s.sling
}

func (s tagSetService) getName() string {
	return s.name
}

func (s tagSetService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new tag set.
func (s tagSetService) Add(resource *model.TagSet) (*model.TagSet, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

// DeleteByID deletes the tag set that matches the input ID.
func (s tagSetService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetByID returns the tag set that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s tagSetService) GetByID(id string) (*model.TagSet, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.TagSet), path)
	if err != nil {
		return nil, createResourceNotFoundError("tag set", "ID", id)
	}

	return resp.(*model.TagSet), nil
}

// GetAll returns all tag sets. If none can be found or an error occurs, it
// returns an empty collection.
func (s tagSetService) GetAll() ([]model.TagSet, error) {
	items := []model.TagSet{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the TagSet with a matching name.
func (s tagSetService) GetByName(name string) (*model.TagSet, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError(operationGetByName, parameterName)
	}

	err := validateInternalState(s)
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

	return nil, createItemNotFoundError(s.name, operationGetByName, name)
}

// Update modifies a tag set based on the one provided as input.
func (s tagSetService) Update(resource model.TagSet) (*model.TagSet, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

var _ ServiceInterface = &tagSetService{}
