package lifecycles

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type LifecycleService struct {
	services.CanDeleteService
}

func NewLifecycleService(sling *sling.Sling, uriTemplate string) *LifecycleService {
	return &LifecycleService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceLifecycleService, sling, uriTemplate),
		},
	}
}

// Add creates a new lifecycle.
func (s *LifecycleService) Add(lifecycle *Lifecycle) (*Lifecycle, error) {
	if IsNil(lifecycle) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, "lifecycle")
	}

	path, err := services.GetAddPath(s, lifecycle)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), lifecycle, new(Lifecycle), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

// Get returns a collection of lifecycles based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *LifecycleService) Get(lifecyclesQuery Query) (*resources.Resources[*Lifecycle], error) {
	path, err := s.GetURITemplate().Expand(lifecyclesQuery)
	if err != nil {
		return &resources.Resources[*Lifecycle]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Lifecycle]), path)
	if err != nil {
		return &resources.Resources[*Lifecycle]{}, err
	}

	return response.(*resources.Resources[*Lifecycle]), nil
}

// GetAll returns all lifecycles. If none can be found or an error occurs, it
// returns an empty collection.
func (s *LifecycleService) GetAll() ([]*Lifecycle, error) {
	items := []*Lifecycle{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the lifecycle that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *LifecycleService) GetByID(id string) (*Lifecycle, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Lifecycle), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

// GetByPartialName performs a lookup and returns a collection of lifecycles
// with a matching partial name.
func (s *LifecycleService) GetByPartialName(partialName string) ([]*Lifecycle, error) {
	if internal.IsEmpty(partialName) {
		return []*Lifecycle{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Lifecycle{}, err
	}

	return services.GetPagedResponse[Lifecycle](s, path)
}

func (s *LifecycleService) GetProjects(lifecycle *Lifecycle) ([]*projects.Project, error) {
	items := []*projects.Project{}

	if lifecycle == nil {
		return items, internal.CreateInvalidParameterError("GetProjects", "lifecycle")
	}

	path := lifecycle.Links["Projects"]
	resp, err := api.ApiGet(s.GetClient(), new([]*projects.Project), path)
	if err != nil {
		return items, err
	}

	return *resp.(*[]*projects.Project), nil

}

// Update modifies a lifecycle based on the one provided as input.
func (s *LifecycleService) Update(lifecycle *Lifecycle) (*Lifecycle, error) {
	path, err := services.GetUpdatePath(s, lifecycle)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), lifecycle, new(Lifecycle), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

var _ services.IService = &LifecycleService{}
