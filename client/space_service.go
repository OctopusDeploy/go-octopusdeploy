package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type spaceService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newSpaceService(sling *sling.Sling, uriTemplate string) *spaceService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &spaceService{
		name:        serviceSpaceService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s spaceService) getClient() *sling.Sling {
	return s.sling
}

func (s spaceService) getName() string {
	return s.name
}

func (s spaceService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new space.
func (s spaceService) Add(resource *model.Space) (*model.Space, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

// DeleteByID deletes the space that matches the input ID.
func (s spaceService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetByID returns the space that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s spaceService) GetByID(id string) (*model.Space, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Space), path)
	if err != nil {
		return nil, createResourceNotFoundError("space", "ID", id)
	}

	return resp.(*model.Space), nil
}

// GetAll returns all spaces. If none can be found or an error occurs, it
// returns an empty collection.
func (s spaceService) GetAll() ([]model.Space, error) {
	items := []model.Space{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Space with a matching name.
func (s spaceService) GetByName(name string) (*model.Space, error) {
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

// Update modifies a space based on the one provided as input.
func (s spaceService) Update(resource model.Space) (*model.Space, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

var _ ServiceInterface = &spaceService{}
