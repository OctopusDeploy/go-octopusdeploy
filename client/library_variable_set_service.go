package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type libraryVariableSetService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
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
		path:        strings.TrimSpace(uriTemplate),
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

func (s libraryVariableSetService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns a single LibraryVariableSet by its Id in Octopus Deploy
func (s libraryVariableSetService) GetByID(id string) (*model.LibraryVariableSet, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// GetAll returns all instances of a LibraryVariableSet. If none can be found or an error occurs, it returns an empty collection.
func (s libraryVariableSetService) GetAll() ([]model.LibraryVariableSet, error) {
	items := new([]model.LibraryVariableSet)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

// GetByPartialName performs a lookup and returns a list of library variable sets with a matching partial name.
func (s libraryVariableSetService) GetByPartialName(name string) ([]model.LibraryVariableSet, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.LibraryVariableSet{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new LibraryVariableSet.
func (s libraryVariableSetService) Add(libraryVariableSet *model.LibraryVariableSet) (*model.LibraryVariableSet, error) {
	path, err := getAddPath(s, libraryVariableSet)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), libraryVariableSet, new(model.LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.LibraryVariableSet), nil
}

// DeleteByID deletes the LibraryVariableSet that matches the input ID.
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

func (s libraryVariableSetService) getPagedResponse(path string) ([]model.LibraryVariableSet, error) {
	items := []model.LibraryVariableSet{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.LibraryVariableSets), path)
		if err != nil {
			return items, err
		}

		responseList := resp.(*model.LibraryVariableSets)
		items = append(items, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return items, nil
}

var _ ServiceInterface = &libraryVariableSetService{}
