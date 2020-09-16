package client

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type LifecycleService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewLifecycleService(sling *sling.Sling, uriTemplate string) *LifecycleService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &LifecycleService{
		sling: sling,
		path:  path,
	}
}

// Get returns a single lifecycle by its lifecycleid in Octopus Deploy
func (s *LifecycleService) Get(id string) (*model.Lifecycle, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("LifecycleService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Lifecycle), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

// GetAll returns all instances of a Lifecycle.
func (s *LifecycleService) GetAll() (*[]model.Lifecycle, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

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

// GetByName performs a lookup and returns the Lifecycle with a matching name.
func (s *LifecycleService) GetByName(name string) (*model.Lifecycle, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("LifecycleService: invalid parameter, name")
	}

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

// Add creates a new Lifecycle.
func (s *LifecycleService) Add(lifecycle *model.Lifecycle) (*model.Lifecycle, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if lifecycle == nil {
		return nil, errors.New("LifecycleService: invalid parameter, lifecycle")
	}

	err = model.ValidateLifecycleValues(lifecycle)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, lifecycle, new(model.Lifecycle), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Lifecycle), nil
}

// Delete deletes an existing lifecycle in Octopus Deploy
func (s *LifecycleService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("LifecycleService: invalid parameter, id")
	}

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

func (s *LifecycleService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("LifecycleService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("LifecycleService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &LifecycleService{}
