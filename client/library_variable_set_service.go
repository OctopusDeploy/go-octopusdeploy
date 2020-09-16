package client

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type LibraryVariableSetService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewLibraryVariableSetService(sling *sling.Sling, uriTemplate string) *LibraryVariableSetService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &LibraryVariableSetService{
		sling: sling,
		path:  path,
	}
}

// Get returns a single LibraryVariableSet by its Id in Octopus Deploy
func (s *LibraryVariableSetService) Get(id string) (*model.LibraryVariableSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("LibraryVariableSetService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.LibraryVariableSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// GetAll returns all instances of a LibraryVariableSet.
func (s *LibraryVariableSetService) GetAll() (*[]model.LibraryVariableSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	return s.get("")
}

func (s *LibraryVariableSetService) get(query string) (*[]model.LibraryVariableSet, error) {
	var p []model.LibraryVariableSet

	path := s.path + "?take=2147483647"
	loadNextPage := true

	if query != "" {
		path = fmt.Sprintf("%s&%s", path, query)
	}

	for loadNextPage { // Older Octopus Servers do not accept the take parameter, so the only choice is to page through them
		resp, err := apiGet(s.sling, new(model.LibraryVariableSets), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.LibraryVariableSets)

		p = append(p, r.Items...)

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName performs a lookup and returns the LibraryVariableSet with a matching name.
func (s *LibraryVariableSetService) GetByName(name string) (*model.LibraryVariableSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("LibraryVariableSetService: invalid parameter, name")
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

// Add creates a new LibraryVariableSet.
func (s *LibraryVariableSetService) Add(libraryVariableSet *model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if libraryVariableSet == nil {
		return nil, errors.New("LibraryVariableSetService: invalid parameter, libraryVariableSet")
	}

	err = model.ValidateLibraryVariableSetValues(libraryVariableSet)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, libraryVariableSet, new(model.LibraryVariableSet), "libraryVariableSets")

	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// Delete deletes an existing libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("LibraryVariableSetService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Update(resource *model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	err := model.ValidateLibraryVariableSetValues(resource)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.LibraryVariableSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

func (s *LibraryVariableSetService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("LibraryVariableSetService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("LibraryVariableSetService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &LibraryVariableSetService{}
