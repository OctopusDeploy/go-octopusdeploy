package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type libraryVariableSetService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newLibraryVariableSetService(sling *sling.Sling, uriTemplate string) *libraryVariableSetService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &libraryVariableSetService{
		name:        serviceLibraryVariableSetService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s libraryVariableSetService) getClient() *sling.Sling {
	return s.sling
}

func (s libraryVariableSetService) getName() string {
	return s.name
}

func (s libraryVariableSetService) getPagedResponse(path string) ([]model.LibraryVariableSet, error) {
	resources := []model.LibraryVariableSet{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.LibraryVariableSets), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.LibraryVariableSets)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s libraryVariableSetService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns the library variable set that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s libraryVariableSetService) GetByID(id string) (*model.LibraryVariableSet, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.LibraryVariableSet), path)
	if err != nil {
		return nil, createResourceNotFoundError("library variable set", "ID", id)
	}

	return resp.(*model.LibraryVariableSet), nil
}

// GetAll returns all library variable sets. If none can be found or an error
// occurs, it returns an empty collection.
func (s libraryVariableSetService) GetAll() ([]model.LibraryVariableSet, error) {
	items := []model.LibraryVariableSet{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByPartialName performs a lookup and returns a list of library variable sets with a matching partial name.
func (s libraryVariableSetService) GetByPartialName(name string) ([]model.LibraryVariableSet, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.LibraryVariableSet{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new library variable set.
func (s libraryVariableSetService) Add(resource *model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// DeleteByID deletes the library variable set that matches the input ID.
func (s libraryVariableSetService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// Update modifies a library variable set based on the one provided as input.
func (s libraryVariableSetService) Update(resource model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

var _ ServiceInterface = &libraryVariableSetService{}
