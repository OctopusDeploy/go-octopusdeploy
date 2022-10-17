package credentials

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
	"strings"
)

type Service struct {
	services.CanDeleteService
}

// NewService returns a service with a preconfigured client.
func NewService(sling *sling.Sling, uriTemplate string) *Service {
	return &Service{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceGitCredentialService, sling, uriTemplate),
		},
	}
}

// Add creates a new resource.
func (s *Service) Add(resource *Resource) (*Resource, error) {
	if IsNil(resource) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterGitCredential)
	}

	path, err := services.GetAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), resource, new(Resource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Resource), nil
}

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s *Service) Get(query Query) (*resources.Resources[*Resource], error) {
	path, err := s.GetURITemplate().Expand(query)
	if err != nil {
		return &resources.Resources[*Resource]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Resource]), path)
	if err != nil {
		return &resources.Resources[*Resource]{}, err
	}

	return response.(*resources.Resources[*Resource]), nil
}

// GetByID returns the Git credential that matches the input ID. If one cannot be found, it returns nil and an error.
func (s *Service) GetByID(id string) (*Resource, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Resource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Resource), nil
}

func (s *Service) GetByPartialName(partialName string) ([]*Resource, error) {
	if internal.IsEmpty(partialName) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Resource{}, err
	}

	return services.GetPagedResponse[Resource](s, path)
}

func (s *Service) GetByName(name string) (*Resource, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	credentials, err := s.GetByPartialName(name)
	if err != nil {
		return nil, err
	}
	for _, creds := range credentials {
		if strings.EqualFold(creds.Name, name) {
			return creds, nil
		}
	}

	return nil, services.ErrItemNotFound
}

func (s *Service) GetByIDOrName(idOrName string) (*Resource, error) {
	creds, err := s.GetByID(idOrName)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if creds != nil {
			return creds, nil
		}
	}

	return s.GetByName(idOrName)
}

// Update modifies a Git credential based on the one provided as input.
func (s *Service) Update(gitCredential *Resource) (*Resource, error) {
	if gitCredential == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterGitCredential)
	}

	path, err := services.GetUpdatePath(s, gitCredential)
	if err != nil {
		return nil, err
	}

	_, err = services.ApiUpdate(s.GetClient(), gitCredential, new(Resource), path)
	if err != nil {
		return nil, err
	}

	// TODO: remove this once the API is fixed
	return s.GetByID(gitCredential.GetID())
}
