package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type deploymentProcessService struct {
	service
}

func newDeploymentProcessService(sling *sling.Sling, uriTemplate string) *deploymentProcessService {
	return &deploymentProcessService{
		service: newService(ServiceDeploymentProcessesService, sling, uriTemplate),
	}
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
	path, err := getUpdatePath(s, deploymentProcess)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), deploymentProcess, new(DeploymentProcess), path)
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
