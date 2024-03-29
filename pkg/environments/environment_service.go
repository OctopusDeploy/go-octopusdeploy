package environments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type EnvironmentService struct {
	sortOrderPath string
	summaryPath   string

	services.CanDeleteService
}

// NewEnvironmentService returns an EnvironmentService with a preconfigured
// client.
func NewEnvironmentService(sling *sling.Sling, uriTemplate string, sortOrderPath string, summaryPath string) *EnvironmentService {
	return &EnvironmentService{
		sortOrderPath: sortOrderPath,
		summaryPath:   summaryPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceEnvironmentService, sling, uriTemplate),
		},
	}
}

// Add creates a new environment.
//
// Deprecated: use environments.Add
func (s *EnvironmentService) Add(environment *Environment) (*Environment, error) {
	if IsNil(environment) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterEnvironment)
	}

	path, err := services.GetAddPath(s, environment)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), environment, new(Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
//
// Deprecated: use environments.Get
func (s *EnvironmentService) Get(environmentsQuery EnvironmentsQuery) (*resources.Resources[*Environment], error) {
	path, err := s.GetURITemplate().Expand(environmentsQuery)
	if err != nil {
		return &resources.Resources[*Environment]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Environment]), path)
	if err != nil {
		return &resources.Resources[*Environment]{}, err
	}

	return response.(*resources.Resources[*Environment]), nil
}

// GetAll returns all environments. If none can be found or an error occurs, it
// returns an empty collection.
//
// Deprecated: use environments.GetAll
func (s *EnvironmentService) GetAll() ([]*Environment, error) {
	items := []*Environment{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: use environments.GetByID
func (s *EnvironmentService) GetByID(id string) (*Environment, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

// GetByIDs returns the environments that match the input IDs.
func (s *EnvironmentService) GetByIDs(ids []string) ([]*Environment, error) {
	if len(ids) == 0 {
		return []*Environment{}, nil
	}

	path, err := services.GetByIDsPath(s, ids)
	if err != nil {
		return []*Environment{}, err
	}

	return services.GetPagedResponse[Environment](s, path)
}

// GetByName returns the environments with a matching partial name.
func (s *EnvironmentService) GetByName(name string) ([]*Environment, error) {
	if internal.IsEmpty(name) {
		return []*Environment{}, internal.CreateInvalidParameterError("GetByName", "name")
	}

	path, err := services.GetByNamePath(s, name)
	if err != nil {
		return []*Environment{}, err
	}

	return services.GetPagedResponse[Environment](s, path)
}

// GetByPartialName performs a lookup and returns enironments with a matching
// partial name.
func (s *EnvironmentService) GetByPartialName(partialName string) ([]*Environment, error) {
	if internal.IsEmpty(partialName) {
		return []*Environment{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Environment{}, err
	}

	return services.GetPagedResponse[Environment](s, path)
}

// Update modifies an environment based on the one provided as input.
//
// Deprecated: use environments.Update
func (s *EnvironmentService) Update(environment *Environment) (*Environment, error) {
	if environment == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterEnvironment)
	}

	path, err := services.GetUpdatePath(s, environment)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), environment, new(Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

// --- new ---

const template = "/api/{spaceId}/environments{/id}{?name,skip,ids,take,partialName}"

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func Get(client newclient.Client, spaceID string, environmentsQuery EnvironmentsQuery) (*resources.Resources[*Environment], error) {
	return newclient.GetByQuery[Environment](client, template, spaceID, environmentsQuery)
}

// Add creates a new environment.
func Add(client newclient.Client, environment *Environment) (*Environment, error) {
	return newclient.Add[Environment](client, template, environment.SpaceID, environment)
}

// DeleteById deletes the environment based on the ID provided as input.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetByID returns the environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*Environment, error) {
	return newclient.GetByID[Environment](client, template, spaceID, ID)
}

// Update modifies an environment based on the one provided as input.
func Update(client newclient.Client, environment *Environment) (*Environment, error) {
	return newclient.Update[Environment](client, template, environment.SpaceID, environment.ID, environment)
}

// GetAll returns all environments. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*Environment, error) {
	return newclient.GetAll[Environment](client, template, spaceID)
}
