package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type machineService struct {
	discoverMachinePath  string
	operatingSystemsPath string
	shellsPath           string

	canDeleteService
}

func newMachineService(sling *sling.Sling, uriTemplate string, discoverMachinePath string, operatingSystemsPath string, shellsPath string) *machineService {
	machineService := &machineService{
		discoverMachinePath:  discoverMachinePath,
		operatingSystemsPath: operatingSystemsPath,
		shellsPath:           shellsPath,
	}
	machineService.service = newService(ServiceMachineService, sling, uriTemplate)

	return machineService
}

func (s machineService) getPagedResponse(path string) ([]*DeploymentTarget, error) {
	resources := []*DeploymentTarget{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(DeploymentTargets), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*DeploymentTargets)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new machine.
func (s machineService) Add(resource *DeploymentTarget) (*DeploymentTarget, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(DeploymentTarget), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentTarget), nil
}

// Get returns a collection of machines based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s machineService) Get(machinesQuery MachinesQuery) (*DeploymentTargets, error) {
	path, err := s.getURITemplate().Expand(machinesQuery)
	if err != nil {
		return &DeploymentTargets{}, err
	}

	response, err := apiGet(s.getClient(), new(DeploymentTargets), path)
	if err != nil {
		return &DeploymentTargets{}, err
	}

	return response.(*DeploymentTargets), nil
}

// GetByID returns the machine that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s machineService) GetByID(id string) (*DeploymentTarget, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(DeploymentTarget), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*DeploymentTarget), nil
}

// GetAll returns all machines. If none can be found or an error occurs, it
// returns an empty collection.
func (s machineService) GetAll() ([]*DeploymentTarget, error) {
	items := []*DeploymentTarget{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Machine with a matching name.
func (s machineService) GetByName(name string) ([]*DeploymentTarget, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*DeploymentTarget{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns the machine with a matching
// partial name.
func (s machineService) GetByPartialName(name string) ([]*DeploymentTarget, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*DeploymentTarget{}, err
	}

	return s.getPagedResponse(path)
}

// Update updates an existing machine in Octopus Deploy
func (s machineService) Update(resource *DeploymentTarget) (*DeploymentTarget, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(DeploymentTarget), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentTarget), nil
}

var _ IService = &machineService{}
