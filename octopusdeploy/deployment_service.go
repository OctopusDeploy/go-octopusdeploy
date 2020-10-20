package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// deploymentService handles communication for any operations in the Octopus
// API that pertain to deployments.
type deploymentService struct {
	service
}

// newDeploymentService returns a deploymentService with a preconfigured
// client.
func newDeploymentService(sling *sling.Sling, uriTemplate string) *deploymentService {
	deploymentService := &deploymentService{}
	deploymentService.service = newService(serviceDeploymentService, sling, uriTemplate, new(Deployment))

	return deploymentService
}

func (s deploymentService) getPagedResponse(path string) ([]*Deployment, error) {
	resources := []*Deployment{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Deployments), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Deployments)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new deployment.
func (s deploymentService) Add(resource *Deployment) (*Deployment, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Deployment), nil
}

// GetByID gets a deployment that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s deploymentService) GetByID(id string) (*Deployment, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Deployment), nil
}

// GetByIDs gets a list of deployments that match the input IDs.
func (s deploymentService) GetByIDs(ids []string) ([]*Deployment, error) {
	if len(ids) == 0 {
		return []*Deployment{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*Deployment{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns instances of a Deployment with a matching partial name.
func (s deploymentService) GetByName(name string) ([]*Deployment, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*Deployment{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a Deployment based on the one provided as input.
func (s deploymentService) Update(resource Deployment) (*Deployment, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Deployment), nil
}
