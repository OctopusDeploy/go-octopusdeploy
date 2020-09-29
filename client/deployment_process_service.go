package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type deploymentProcessService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newDeploymentProcessService(sling *sling.Sling, uriTemplate string) *deploymentProcessService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &deploymentProcessService{
		name:        serviceDeploymentProcessService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s deploymentProcessService) getClient() *sling.Sling {
	return s.sling
}

func (s deploymentProcessService) getName() string {
	return s.name
}

func (s deploymentProcessService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetAll returns all deployment processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s deploymentProcessService) GetAll() ([]model.DeploymentProcess, error) {
	items := []model.DeploymentProcess{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the deployment process that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s deploymentProcessService) GetByID(id string) (*model.DeploymentProcess, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.DeploymentProcess), path)
	if err != nil {
		return nil, createResourceNotFoundError("deployment process", "ID", id)
	}

	return resp.(*model.DeploymentProcess), nil
}

func (s deploymentProcessService) Update(resource model.DeploymentProcess) (*model.DeploymentProcess, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentProcess), nil
}

func (s deploymentProcessService) getPagedResponse(path string) ([]model.DeploymentProcess, error) {
	resources := []model.DeploymentProcess{}
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

var _ ServiceInterface = &deploymentProcessService{}
