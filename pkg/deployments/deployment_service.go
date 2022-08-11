package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
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

	resp, err := api.ApiGet(s.GetClient(), new(Deployment), path)
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

func (s *DeploymentService) GetDeployments(release *releases.Release, deploymentQuery ...*DeploymentQuery) (*resources.Resources[*Deployment], error) {
	if release == nil {
		return nil, internal.CreateInvalidParameterError("GetDeployments", "release")
	}

	uriTemplate, err := uritemplates.Parse(release.GetLinks()[constants.LinkDeployments])
	if err != nil {
		return &resources.Resources[*Deployment]{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &resources.Resources[*Deployment]{}, err
	}

	if deploymentQuery != nil {
		path, err = uriTemplate.Expand(deploymentQuery[0])
		if err != nil {
			return &resources.Resources[*Deployment]{}, err
		}
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Deployment]), path)
	if err != nil {
		return &resources.Resources[*Deployment]{}, err
	}

	return resp.(*resources.Resources[*Deployment]), nil
}

func (s *DeploymentService) GetProgression(release *releases.Release) (*releases.Progression, error) {
	if release == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetDeployments, constants.ParameterRelease)
	}

	path := release.GetLinks()[constants.LinkProgression]
	resp, err := api.ApiGet(s.GetClient(), new(releases.Progression), path)
	if err != nil {
		return nil, err
	}

	return resp.(*releases.Progression), nil
}

// GetDeploymentSettings loads the deployment settings for a project.
// If the project is version controlled you'll need to specify a gitRef such as 'main'
func (s *DeploymentService) GetDeploymentSettings(project *projects.Project, gitRef string) (*DeploymentSettings, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetDeploymentSettings", constants.ParameterProject)
	}
	if s.GetClient() == nil { // don't call ValidateInternalState because it checks the "Deployments" link which we don't use or need
		return nil, internal.CreateInvalidClientStateError(s.GetName())
	}

	template, err := uritemplates.Parse(project.Links[constants.LinkDeploymentSettings])
	if err != nil {
		return nil, err
	}

	var templateParameters map[string]interface{}
	if gitRef != "" {
		templateParameters = map[string]interface{}{"gitRef": gitRef}
	} else { // non-CaC project links don't have templates so this is a no-op, but it would safely remove any {?extra}{junk} that might be on the query string
		templateParameters = map[string]interface{}{}
	}

	path, err := template.Expand(templateParameters)
	if err != nil {
		return nil, err
	}
	resp, err := api.ApiGet(s.GetClient(), new(DeploymentSettings), path)
	if err != nil {
		return nil, err
	}
	return resp.(*DeploymentSettings), nil
}
