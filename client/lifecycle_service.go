package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type lifecycleService struct {
	service
}

func newLifecycleService(sling *sling.Sling, uriTemplate string) *lifecycleService {
	lifecycleService := &lifecycleService{}
	lifecycleService.service = newService(serviceLifecycleService, sling, uriTemplate, new(model.Lifecycle))

	return lifecycleService
}

func (s lifecycleService) getPagedResponse(path string) ([]*model.Lifecycle, error) {
	resources := []*model.Lifecycle{}
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

// GetByID returns the lifecycle that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s lifecycleService) GetByID(id string) (*model.Lifecycle, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Lifecycle), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.Lifecycle), nil
}

// GetAll returns all lifecycles. If none can be found or an error occurs, it
// returns an empty collection.
func (s lifecycleService) GetAll() ([]*model.Lifecycle, error) {
	items := []*model.Lifecycle{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByPartialName performs a lookup and returns instances of a Lifecycle with a matching partial name.
func (s lifecycleService) GetByPartialName(name string) ([]*model.Lifecycle, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.Lifecycle{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new lifecycle.
func (s lifecycleService) Add(resource *model.Lifecycle) (*model.Lifecycle, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

// Update modifies a Lifecycle based on the one provided as input.
func (s lifecycleService) Update(resource model.Lifecycle) (*model.Lifecycle, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

var _ IService = &lifecycleService{}
