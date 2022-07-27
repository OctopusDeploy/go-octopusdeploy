package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

// DeploymentService handles communication for any operations in the Octopus
// API that pertain to deployments.
type DeploymentService struct {
	services.CanDeleteService
}

// NewDeploymentService returns a deploymentService with a preconfigured
// client.
func NewDeploymentService(sling *sling.Sling, uriTemplate string) *DeploymentService {
	return &DeploymentService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceDeploymentService, sling, uriTemplate),
		},
	}
}

// Add creates a new deployment.
func (s *DeploymentService) Add(deployment *Deployment) (*Deployment, error) {
	if IsNil(deployment) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, "deployment")
	}

	path, err := services.GetAddPath(s, deployment)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), deployment, new(Deployment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Deployment), nil
}

// GetByID gets a deployment that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s *DeploymentService) GetByID(id string) (*Deployment, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(Deployment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Deployment), nil
}

// GetByIDs gets a list of deployments that match the input IDs.
func (s *DeploymentService) GetByIDs(ids []string) ([]*Deployment, error) {
	if len(ids) == 0 {
		return []*Deployment{}, nil
	}

	path, err := services.GetByIDsPath(s, ids)
	if err != nil {
		return []*Deployment{}, err
	}

	return services.GetPagedResponse[Deployment](s, path)
}

// GetByName performs a lookup and returns instances of a Deployment with a matching partial name.
func (s *DeploymentService) GetByName(name string) ([]*Deployment, error) {
	if internal.IsEmpty(name) {
		return []*Deployment{}, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	path, err := services.GetByNamePath(s, name)
	if err != nil {
		return []*Deployment{}, err
	}

	return services.GetPagedResponse[Deployment](s, path)
}

// Update modifies a Deployment based on the one provided as input.
func (s *DeploymentService) Update(resource Deployment) (*Deployment, error) {
	path, err := services.GetUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(Deployment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Deployment), nil
}

func (s *DeploymentService) GetDeployments(release *releases.Release, deploymentQuery ...*DeploymentQuery) (*resources.Resources[Deployment], error) {
	if release == nil {
		return nil, internal.CreateInvalidParameterError("GetDeployments", "release")
	}

	uriTemplate, err := uritemplates.Parse(release.GetLinks()[constants.LinkDeployments])
	if err != nil {
		return &resources.Resources[Deployment]{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &resources.Resources[Deployment]{}, err
	}

	if deploymentQuery != nil {
		path, err = uriTemplate.Expand(deploymentQuery[0])
		if err != nil {
			return &resources.Resources[Deployment]{}, err
		}
	}

	resp, err := services.ApiGet(s.GetClient(), new(resources.Resources[Deployment]), path)
	if err != nil {
		return &resources.Resources[Deployment]{}, err
	}

	return resp.(*resources.Resources[Deployment]), nil
}

func (s *DeploymentService) GetProgression(release *releases.Release) (*releases.Progression, error) {
	if release == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetDeployments, constants.ParameterRelease)
	}

	path := release.GetLinks()[constants.LinkProgression]
	resp, err := services.ApiGet(s.GetClient(), new(releases.Progression), path)
	if err != nil {
		return nil, err
	}

	return resp.(*releases.Progression), nil
}
