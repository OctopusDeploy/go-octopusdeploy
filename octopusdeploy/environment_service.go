package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type environmentService struct {
	sortOrderPath string
	summaryPath   string

	canDeleteService
}

func newEnvironmentService(sling *sling.Sling, uriTemplate string, sortOrderPath string, summaryPath string) *environmentService {
	environmentService := &environmentService{
		sortOrderPath: sortOrderPath,
		summaryPath:   summaryPath,
	}
	environmentService.service = newService(ServiceEnvironmentService, sling, uriTemplate)

	return environmentService
}

func (s environmentService) getPagedResponse(path string) ([]*Environment, error) {
	resources := []*Environment{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Environments), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Environments)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new environment.
func (s environmentService) Add(environment *Environment) (*Environment, error) {
	if environment == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterEnvironment)
	}

	path, err := getAddPath(s, environment)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), environment, new(Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s environmentService) Get(environmentsQuery EnvironmentsQuery) (*Environments, error) {
	path, err := s.getURITemplate().Expand(environmentsQuery)
	if err != nil {
		return &Environments{}, err
	}

	response, err := apiGet(s.getClient(), new(Environments), path)
	if err != nil {
		return &Environments{}, err
	}

	return response.(*Environments), nil
}

// GetAll returns all environments. If none can be found or an error occurs, it
// returns an empty collection.
func (s environmentService) GetAll() ([]*Environment, error) {
	items := []*Environment{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s environmentService) GetByID(id string) (*Environment, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

// GetByIDs returns the environments that match the input IDs.
func (s environmentService) GetByIDs(ids []string) ([]*Environment, error) {
	if len(ids) == 0 {
		return []*Environment{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*Environment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the environments with a matching partial name.
func (s environmentService) GetByName(name string) ([]*Environment, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*Environment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns enironments with a matching
// partial name.
func (s environmentService) GetByPartialName(name string) ([]*Environment, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Environment{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an environment based on the one provided as input.
func (s environmentService) Update(environment *Environment) (*Environment, error) {
	if environment == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterEnvironment)
	}

	path, err := getUpdatePath(s, environment)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), environment, new(Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}
