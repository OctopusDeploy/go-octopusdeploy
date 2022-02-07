package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type deploymentProcessService struct {
	services.service
}

func newDeploymentProcessService(sling *sling.Sling, uriTemplate string) *deploymentProcessService {
	return &deploymentProcessService{
		service: services.newService(ServiceDeploymentProcessesService, sling, uriTemplate),
	}
}

// Get returns the deployment process that matches the input project and
// a git reference. If one cannot be found, it returns nil and an error.
func (s deploymentProcessService) Get(project *Project, gitRef string) (*DeploymentProcess, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationGet, ParameterProject)
	}

	if project.PersistenceSettings == nil || project.PersistenceSettings.GetType() != "VersionControlled" {
		return s.GetByID(project.DeploymentProcessID)
	}

	gitPersistenceSettings := project.PersistenceSettings.(*GitPersistenceSettings)

	if len(gitRef) <= 0 {
		gitRef = gitPersistenceSettings.DefaultBranch
	}

	template, _ := uritemplates.Parse(project.Links["DeploymentProcess"])
	path, _ := template.Expand(map[string]interface{}{"gitRef": gitRef})

	resp, err := apiGet(s.getClient(), new(DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	deploymentProcess := resp.(*DeploymentProcess)
	deploymentProcess.Branch = gitRef

	return deploymentProcess, err
}

// GetAll returns all deployment processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s deploymentProcessService) GetAll() ([]*DeploymentProcess, error) {
	path, err := getPath(s)
	if err != nil {
		return []*DeploymentProcess{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the deployment process that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s deploymentProcessService) GetByID(id string) (*DeploymentProcess, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

// Update modifies a deployment process based on the one provided as input.
func (s deploymentProcessService) Update(deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	if deploymentProcess == nil {
		return nil, createInvalidParameterError(OperationUpdate, "deploymentProcess")
	}

	resp, err := apiUpdate(s.getClient(), deploymentProcess, new(DeploymentProcess), deploymentProcess.Links["Self"])
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

func (s deploymentProcessService) getPagedResponse(path string) ([]*DeploymentProcess, error) {
	resources := []*DeploymentProcess{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(DeploymentProcesses), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*DeploymentProcesses)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}
