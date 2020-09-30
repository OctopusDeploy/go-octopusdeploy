package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type environmentService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newEnvironmentService(sling *sling.Sling, uriTemplate string) *environmentService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &environmentService{
		name:        serviceEnvironmentService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s environmentService) getClient() *sling.Sling {
	return s.sling
}

func (s environmentService) getName() string {
	return s.name
}

func (s environmentService) getPagedResponse(path string) ([]model.Environment, error) {
	resources := []model.Environment{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Environments), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Environments)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s environmentService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new environment.
func (s environmentService) Add(resource *model.Environment) (*model.Environment, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

// DeleteByID deletes the environment that matches the input ID.
func (s environmentService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("environment", "ID", id)
	}

	return err
}

// GetAll returns all environments. If none can be found or an error occurs, it
// returns an empty collection.
func (s environmentService) GetAll() ([]model.Environment, error) {
	items := []model.Environment{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s environmentService) GetByID(id string) (*model.Environment, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Environment), path)
	if err != nil {
		return nil, createResourceNotFoundError("environment", "ID", id)
	}

	return resp.(*model.Environment), nil
}

// GetByIDs returns the environments that match the input IDs.
func (s environmentService) GetByIDs(ids []string) ([]model.Environment, error) {
	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []model.Environment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the environments with a matching partial name.
func (s environmentService) GetByName(name string) ([]model.Environment, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []model.Environment{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an environment based on the one provided as input.
func (s environmentService) Update(resource model.Environment) (*model.Environment, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

var _ ServiceInterface = &environmentService{}
