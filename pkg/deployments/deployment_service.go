package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
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

func (s *DeploymentService) getPagedResponse(path string) ([]*Deployment, error) {
	resources := []*Deployment{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(Deployments), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Deployments)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
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

	return s.getPagedResponse(path)
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

	return s.getPagedResponse(path)
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

func (s *DeploymentService) GetDeployments(release *releases.Release, deploymentQuery ...*DeploymentQuery) (*Deployments, error) {
	if release == nil {
		return nil, internal.CreateInvalidParameterError("GetDeployments", "release")
	}

	uriTemplate, err := uritemplates.Parse(release.GetLinks()[constants.LinkDeployments])
	if err != nil {
		return &Deployments{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &Deployments{}, err
	}

	if deploymentQuery != nil {
		path, err = uriTemplate.Expand(deploymentQuery[0])
		if err != nil {
			return &Deployments{}, err
		}
	}

	resp, err := services.ApiGet(s.GetClient(), new(Deployments), path)
	if err != nil {
		return &Deployments{}, err
	}

	return resp.(*Deployments), nil
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
