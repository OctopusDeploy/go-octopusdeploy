package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type LibraryVariableSetService struct {
	services.CanDeleteService
}

func NewLibraryVariableSetService(sling *sling.Sling, uriTemplate string) *LibraryVariableSetService {
	return &LibraryVariableSetService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceLibraryVariableSetService, sling, uriTemplate),
		},
	}
}

func (s *LibraryVariableSetService) getPagedResponse(path string) ([]*LibraryVariableSet, error) {
	resources := []*LibraryVariableSet{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(LibraryVariableSets), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*LibraryVariableSets)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new library variable set.
func (s *LibraryVariableSetService) Add(libraryVariableSet *LibraryVariableSet) (*LibraryVariableSet, error) {
	if IsNil(libraryVariableSet) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterLibraryVariableSet)
	}

	path, err := services.GetAddPath(s, libraryVariableSet)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), libraryVariableSet, new(LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}

// Get returns a collection of library variable sets based on the criteria
// defined by its input query parameter. If an error occurs, an empty
// collection is returned along with the associated error.
func (s *LibraryVariableSetService) Get(libraryVariablesQuery LibraryVariablesQuery) (*LibraryVariableSets, error) {
	path, err := s.GetURITemplate().Expand(libraryVariablesQuery)
	if err != nil {
		return &LibraryVariableSets{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(LibraryVariableSets), path)
	if err != nil {
		return &LibraryVariableSets{}, err
	}

	return response.(*LibraryVariableSets), nil
}

// GetAll returns all library variable sets. If none can be found or an error
// occurs, it returns an empty collection.
func (s *LibraryVariableSetService) GetAll() ([]*LibraryVariableSet, error) {
	items := []*LibraryVariableSet{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the library variable set that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s *LibraryVariableSetService) GetByID(id string) (*LibraryVariableSet, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}

// GetByPartialName performs a lookup and returns a list of library variable sets with a matching partial name.
func (s *LibraryVariableSetService) GetByPartialName(name string) ([]*LibraryVariableSet, error) {
	path, err := services.GetByPartialNamePath(s, name)
	if err != nil {
		return []*LibraryVariableSet{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a library variable set based on the one provided as input.
func (s *LibraryVariableSetService) Update(libraryVariableSet *LibraryVariableSet) (*LibraryVariableSet, error) {
	if libraryVariableSet == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterLibraryVariableSet)
	}

	path, err := services.GetUpdatePath(s, libraryVariableSet)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), libraryVariableSet, new(LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}
