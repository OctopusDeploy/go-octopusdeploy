package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type DeploymentProcessService struct {
	services.Service
}

func NewDeploymentProcessService(sling *sling.Sling, uriTemplate string) *DeploymentProcessService {
	return &DeploymentProcessService{
		Service: services.NewService(constants.ServiceDeploymentProcessesService, sling, uriTemplate),
	}
}

// Get returns the deployment process that matches the input project and
// a git reference. If one cannot be found, it returns nil and an error.
func (s *DeploymentProcessService) Get(project *projects.Project, gitRef string) (*DeploymentProcess, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("Get", "project")
	}

	if project.PersistenceSettings == nil || project.PersistenceSettings.GetType() != "VersionControlled" {
		return s.GetByID(project.DeploymentProcessID)
	}

	gitPersistenceSettings := project.PersistenceSettings.(*projects.GitPersistenceSettings)

	if len(gitRef) <= 0 {
		gitRef = gitPersistenceSettings.DefaultBranch
	}

	template, _ := uritemplates.Parse(project.Links["DeploymentProcess"])
	path, _ := template.Expand(map[string]interface{}{"gitRef": gitRef})

	resp, err := services.ApiGet(s.GetClient(), new(DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	deploymentProcess := resp.(*DeploymentProcess)
	deploymentProcess.Branch = gitRef

	return deploymentProcess, err
}

// GetAll returns all deployment processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s *DeploymentProcessService) GetAll() ([]*DeploymentProcess, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return []*DeploymentProcess{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the deployment process that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s *DeploymentProcessService) GetByID(id string) (*DeploymentProcess, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

// Update modifies a deployment process based on the one provided as input.
func (s *DeploymentProcessService) Update(deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	if deploymentProcess == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "deploymentProcess")
	}

	resp, err := services.ApiUpdate(s.GetClient(), deploymentProcess, new(DeploymentProcess), deploymentProcess.Links["Self"])
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

func (s *DeploymentProcessService) getPagedResponse(path string) ([]*DeploymentProcess, error) {
	resources := []*DeploymentProcess{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(DeploymentProcesses), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*DeploymentProcesses)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}
