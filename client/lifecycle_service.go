package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type lifecycleService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newLifecycleService(sling *sling.Sling, uriTemplate string) *lifecycleService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &lifecycleService{
		name:        serviceLifecycleService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s lifecycleService) getClient() *sling.Sling {
	return s.sling
}

func (s lifecycleService) getName() string {
	return s.name
}

func (s lifecycleService) getPagedResponse(path string) ([]model.Lifecycle, error) {
	resources := []model.Lifecycle{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Lifecycles), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Lifecycles)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s lifecycleService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns the lifecycle that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s lifecycleService) GetByID(id string) (*model.Lifecycle, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Lifecycle), path)
	if err != nil {
		return nil, createResourceNotFoundError("lifecycle", "ID", id)
	}

	return resp.(*model.Lifecycle), nil
}

// GetAll returns all lifecycles. If none can be found or an error occurs, it
// returns an empty collection.
func (s lifecycleService) GetAll() ([]model.Lifecycle, error) {
	items := []model.Lifecycle{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByPartialName performs a lookup and returns instances of a Lifecycle with a matching partial name.
func (s lifecycleService) GetByPartialName(name string) ([]model.Lifecycle, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.Lifecycle{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new lifecycle.
func (s lifecycleService) Add(resource *model.Lifecycle) (*model.Lifecycle, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Lifecycle), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

// DeleteByID deletes the lifecycle that matches the input ID.
func (s lifecycleService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("lifecycle", "ID", id)
	}

	return err
}

// Update modifies a Lifecycle based on the one provided as input.
func (s lifecycleService) Update(resource model.Lifecycle) (*model.Lifecycle, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Lifecycle), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

var _ ServiceInterface = &lifecycleService{}
