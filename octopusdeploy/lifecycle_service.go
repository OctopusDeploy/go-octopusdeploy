package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type lifecycleService struct {
	canDeleteService
}

func newLifecycleService(sling *sling.Sling, uriTemplate string) *lifecycleService {
	lifecycleService := &lifecycleService{}
	lifecycleService.service = newService(serviceLifecycleService, sling, uriTemplate, new(Lifecycle))

	return lifecycleService
}

func (s lifecycleService) getPagedResponse(path string) ([]*Lifecycle, error) {
	resources := []*Lifecycle{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Lifecycles), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Lifecycles)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new lifecycle.
func (s lifecycleService) Add(resource *Lifecycle) (*Lifecycle, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

// GetAll returns all lifecycles. If none can be found or an error occurs, it
// returns an empty collection.
func (s lifecycleService) GetAll() ([]*Lifecycle, error) {
	items := []*Lifecycle{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the lifecycle that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s lifecycleService) GetByID(id string) (*Lifecycle, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Lifecycle), nil
}

// GetByPartialName performs a lookup and returns a collection of lifecycles
// with a matching partial name.
func (s lifecycleService) GetByPartialName(name string) ([]*Lifecycle, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Lifecycle{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a lifecycle based on the one provided as input.
func (s lifecycleService) Update(lifecycle *Lifecycle) (*Lifecycle, error) {
	path, err := getUpdatePath(s, lifecycle)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), lifecycle, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

var _ IService = &lifecycleService{}
