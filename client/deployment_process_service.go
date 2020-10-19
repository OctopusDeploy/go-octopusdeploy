package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type deploymentProcessService struct {
	service
}

func newDeploymentProcessService(sling *sling.Sling, uriTemplate string) *deploymentProcessService {
	return &deploymentProcessService{
		service: newService(serviceDeploymentProcesseService, sling, uriTemplate, new(model.DeploymentProcess)),
	}
}

// GetAll returns all deployment processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s deploymentProcessService) GetAll() ([]*model.DeploymentProcess, error) {
	path, err := getPath(s)
	if err != nil {
		return []*model.DeploymentProcess{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the deployment process that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s deploymentProcessService) GetByID(id string) (*model.DeploymentProcess, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.DeploymentProcess), nil
}

func (s deploymentProcessService) Update(resource model.DeploymentProcess) (*model.DeploymentProcess, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentProcess), nil
}

func (s deploymentProcessService) getPagedResponse(path string) ([]*model.DeploymentProcess, error) {
	resources := []*model.DeploymentProcess{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.DeploymentProcesses), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.DeploymentProcesses)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}
