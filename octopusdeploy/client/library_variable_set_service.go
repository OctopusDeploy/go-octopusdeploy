package client

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type LibraryVariableSetService struct {
	sling *sling.Sling
	path  string
}

func NewLibraryVariableSetService(sling *sling.Sling) *LibraryVariableSetService {
	return &LibraryVariableSetService{
		sling: sling,
		path:  "libraryvariablesets",
	}
}

// Get returns a single LibraryVariableSet by its Id in Octopus Deploy
func (s *LibraryVariableSetService) Get(id string) (*model.LibraryVariableSet, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.LibraryVariableSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// GetAll returns all libraryVariableSets in Octopus Deploy
func (s *LibraryVariableSetService) GetAll() (*[]model.LibraryVariableSet, error) {
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

// GetByName gets an existing Library Variable Set by its name in Octopus Deploy
func (s *LibraryVariableSetService) GetByName(name string) (*model.LibraryVariableSet, error) {
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

// Add adds an new libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Add(resource *model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	err := model.ValidateLibraryVariableSetValues(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.LibraryVariableSet), "libraryVariableSets")

	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// Delete deletes an existing libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Update(resource *model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	err := model.ValidateLibraryVariableSetValues(resource)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("libraryVariableSets/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.LibraryVariableSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}
