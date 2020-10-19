package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type machineService struct {
	discoverMachinePath  string
	operatingSystemsPath string
	shellsPath           string

	service
}

func newMachineService(sling *sling.Sling, uriTemplate string, discoverMachinePath string, operatingSystemsPath string, shellsPath string) *machineService {
	machineService := &machineService{
		discoverMachinePath:  discoverMachinePath,
		operatingSystemsPath: operatingSystemsPath,
		shellsPath:           shellsPath,
	}
	machineService.service = newService(serviceMachineService, sling, uriTemplate, new(model.DeploymentTarget))

	return machineService
}

func (s machineService) getPagedResponse(path string) ([]*model.DeploymentTarget, error) {
	resources := []*model.DeploymentTarget{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.DeploymentTargets), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.DeploymentTargets)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new machine.
func (s machineService) Add(resource *model.DeploymentTarget) (*model.DeploymentTarget, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentTarget), nil
}

// GetByID returns the machine that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s machineService) GetByID(id string) (*model.DeploymentTarget, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.DeploymentTarget), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.DeploymentTarget), nil
}

// GetAll returns all machines. If none can be found or an error occurs, it
// returns an empty collection.
func (s machineService) GetAll() ([]*model.DeploymentTarget, error) {
	items := []*model.DeploymentTarget{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Machine with a matching name.
func (s machineService) GetByName(name string) ([]*model.DeploymentTarget, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*model.DeploymentTarget{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns the machine with a matching
// partial name.
func (s machineService) GetByPartialName(name string) ([]*model.DeploymentTarget, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.DeploymentTarget{}, err
	}

	return s.getPagedResponse(path)
}

// Update updates an existing machine in Octopus Deploy
func (s machineService) Update(resource model.DeploymentTarget) (*model.DeploymentTarget, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentTarget), nil
}

var _ IService = &machineService{}
