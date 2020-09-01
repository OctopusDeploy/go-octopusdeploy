package client

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type LifecycleService struct {
	sling *sling.Sling
	path  string
}

func NewLifecycleService(sling *sling.Sling) *LifecycleService {
	return &LifecycleService{
		sling: sling,
		path:  "lifecycles",
	}
}

// Get returns a single lifecycle by its lifecycleid in Octopus Deploy
func (s *LifecycleService) Get(id string) (*model.Lifecycle, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Lifecycle), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

// GetAll returns all lifecycles in Octopus Deploy
func (s *LifecycleService) GetAll() (*[]model.Lifecycle, error) {
	return s.get("")
}

func (s *LifecycleService) get(query string) (*[]model.Lifecycle, error) {
	var p []model.Lifecycle

	path := s.path + "?take=2147483647"
	loadNextPage := true

	if query != "" {
		path = fmt.Sprintf("%s&%s", path, query)
	}

	for loadNextPage { // Older Octopus Servers do not accept the take parameter, so the only choice is to page through them
		resp, err := apiGet(s.sling, new(model.Lifecycles), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Lifecycles)

		p = append(p, r.Items...)

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName gets an existing lifecycle by its lifecycle name in Octopus Deploy
func (s *LifecycleService) GetByName(name string) (*model.Lifecycle, error) {
	collection, err := s.get(fmt.Sprintf("partialName=%s", url.PathEscape(name)))

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

// Add adds an new lifecycle in Octopus Deploy
func (s *LifecycleService) Add(resource *model.Lifecycle) (*model.Lifecycle, error) {
	err := model.ValidateLifecycleValues(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Lifecycle), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

// Delete deletes an existing lifecycle in Octopus Deploy
func (s *LifecycleService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing lifecycle in Octopus Deploy
func (s *LifecycleService) Update(resource *model.Lifecycle) (*model.Lifecycle, error) {
	err := model.ValidateLifecycleValues(resource)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Lifecycle), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}
