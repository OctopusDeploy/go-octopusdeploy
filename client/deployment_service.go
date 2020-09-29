package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// deploymentService handles communication for any operations in the Octopus
// API that pertain to deployments.
type deploymentService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newDeploymentService returns a deploymentService with a preconfigured
// client.
func newDeploymentService(sling *sling.Sling, uriTemplate string) *deploymentService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &deploymentService{
		name:        serviceDeploymentService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s deploymentService) getClient() *sling.Sling {
	return s.sling
}

func (s deploymentService) getName() string {
	return s.name
}

func (s deploymentService) getPagedResponse(path string) ([]model.Deployment, error) {
	resources := []model.Deployment{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Deployments), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Deployments)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s deploymentService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new deployment.
func (s deploymentService) Add(resource *model.Deployment) (*model.Deployment, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Deployment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Deployment), nil
}

// DeleteByID deletes the deployment that matches the input ID.
func (s deploymentService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetByID gets a deployment that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s deploymentService) GetByID(id string) (*model.Deployment, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Deployment), path)
	if err != nil {
		return nil, createResourceNotFoundError("deployment", "ID", id)
	}

	return resp.(*model.Deployment), nil
}

// GetByIDs gets a list of deployments that match the input IDs.
func (s deploymentService) GetByIDs(ids []string) ([]model.Deployment, error) {
	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []model.Deployment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns instances of a Deployment with a matching partial name.
func (s deploymentService) GetByName(name string) ([]model.Deployment, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []model.Deployment{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a Deployment based on the one provided as input.
func (s deploymentService) Update(resource model.Deployment) (*model.Deployment, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Deployment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Deployment), nil
}

var _ ServiceInterface = &deploymentService{}
