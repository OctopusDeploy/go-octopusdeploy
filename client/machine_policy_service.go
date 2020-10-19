package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type machinePolicyService struct {
	templatePath string

	service
}

func newMachinePolicyService(sling *sling.Sling, uriTemplate string, templatePath string) *machinePolicyService {
	machinePolicyService := &machinePolicyService{
		templatePath: templatePath,
	}
	machinePolicyService.service = newService(serviceMachinePolicyService, sling, uriTemplate, new(model.MachinePolicy))

	return machinePolicyService
}

func (s machinePolicyService) getPagedResponse(path string) ([]*model.MachinePolicy, error) {
	resources := []*model.MachinePolicy{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.MachinePolicies), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.MachinePolicies)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new machine policy.
func (s machinePolicyService) Add(resource *model.MachinePolicy) (*model.MachinePolicy, error) {
	if resource == nil {
		return nil, createInvalidParameterError(operationAdd, parameterMachinePolicy)
	}

	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}

// GetAll returns all machine policies. If none can be found or an error
// occurs, it returns an empty collection.
func (s machinePolicyService) GetAll() ([]*model.MachinePolicy, error) {
	items := []*model.MachinePolicy{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the machine policy that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s machinePolicyService) GetByID(id string) (*model.MachinePolicy, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.MachinePolicy), nil
}

// GetByPartialName performs a lookup and returns machine policies with a
// matching partial name.
func (s machinePolicyService) GetByPartialName(name string) ([]*model.MachinePolicy, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.MachinePolicy{}, err
	}

	return s.getPagedResponse(path)
}

func (s machinePolicyService) GetTemplate() (*model.MachinePolicy, error) {
	resp, err := apiGet(s.getClient(), s.itemType, s.templatePath)
	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}

// Update modifies a machine policy based on the one provided as input.
func (s machinePolicyService) Update(machinePolicy *model.MachinePolicy) (*model.MachinePolicy, error) {
	path, err := getUpdatePath(s, machinePolicy)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), machinePolicy, new(model.MachinePolicy), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}
