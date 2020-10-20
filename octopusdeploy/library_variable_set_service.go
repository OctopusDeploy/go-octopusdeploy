package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type libraryVariableSetService struct {
	service
}

func newLibraryVariableSetService(sling *sling.Sling, uriTemplate string) *libraryVariableSetService {
	libraryVariableSetService := &libraryVariableSetService{}
	libraryVariableSetService.service = newService(serviceLibraryVariableSetService, sling, uriTemplate, new(LibraryVariableSet))

	return libraryVariableSetService
}

func (s libraryVariableSetService) getPagedResponse(path string) ([]*LibraryVariableSet, error) {
	resources := []*LibraryVariableSet{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(LibraryVariableSets), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*LibraryVariableSets)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the library variable set that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s libraryVariableSetService) GetByID(id string) (*LibraryVariableSet, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(LibraryVariableSet), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*LibraryVariableSet), nil
}

// GetAll returns all library variable sets. If none can be found or an error
// occurs, it returns an empty collection.
func (s libraryVariableSetService) GetAll() ([]*LibraryVariableSet, error) {
	items := []*LibraryVariableSet{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByPartialName performs a lookup and returns a list of library variable sets with a matching partial name.
func (s libraryVariableSetService) GetByPartialName(name string) ([]*LibraryVariableSet, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*LibraryVariableSet{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new library variable set.
func (s libraryVariableSetService) Add(resource *LibraryVariableSet) (*LibraryVariableSet, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}

// Update modifies a library variable set based on the one provided as input.
func (s libraryVariableSetService) Update(libraryVariableSet *LibraryVariableSet) (*LibraryVariableSet, error) {
	if libraryVariableSet == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterLibraryVariableSet)
	}

	path, err := getUpdatePath(s, libraryVariableSet)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), libraryVariableSet, new(LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}