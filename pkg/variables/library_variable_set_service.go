package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
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

// Add creates a new library variable set.
//
// Deprecated: Use libraryvariableset.Add
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
func (s *LibraryVariableSetService) Get(libraryVariablesQuery LibraryVariablesQuery) (*resources.Resources[*LibraryVariableSet], error) {
	path, err := s.GetURITemplate().Expand(libraryVariablesQuery)
	if err != nil {
		return &resources.Resources[*LibraryVariableSet]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*LibraryVariableSet]), path)
	if err != nil {
		return &resources.Resources[*LibraryVariableSet]{}, err
	}

	return response.(*resources.Resources[*LibraryVariableSet]), nil
}

// GetAll returns all library variable sets. If none can be found or an error
// occurs, it returns an empty collection.
func (s *LibraryVariableSetService) GetAll() ([]*LibraryVariableSet, error) {
	items := []*LibraryVariableSet{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the library variable set that matches the input ID. If one
// cannot be found, it returns nil and an error.
//
// Deprecated: Use libraryvariableset.GetByID
func (s *LibraryVariableSetService) GetByID(id string) (*LibraryVariableSet, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(LibraryVariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}

// GetByPartialName performs a lookup and returns a list of library variable sets with a matching partial name.
func (s *LibraryVariableSetService) GetByPartialName(partialName string) ([]*LibraryVariableSet, error) {
	if internal.IsEmpty(partialName) {
		return []*LibraryVariableSet{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*LibraryVariableSet{}, err
	}

	return services.GetPagedResponse[LibraryVariableSet](s, path)
}

// Update modifies a library variable set based on the one provided as input.
//
// Deprecated: Use libraryvariableset.Update
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
