package machines

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type MachineService struct {
	discoverMachinePath  string
	operatingSystemsPath string
	shellsPath           string

	services.CanDeleteService
}

func NewMachineService(sling *sling.Sling, uriTemplate string, discoverMachinePath string, operatingSystemsPath string, shellsPath string) *MachineService {
	return &MachineService{
		discoverMachinePath:  discoverMachinePath,
		operatingSystemsPath: operatingSystemsPath,
		shellsPath:           shellsPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceMachineService, sling, uriTemplate),
		},
	}
}

// Add creates a new machine.
func (s *MachineService) Add(deploymentTarget *DeploymentTarget) (*DeploymentTarget, error) {
	if IsNil(deploymentTarget) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterDeploymentTarget)
	}

	path, err := services.GetAddPath(s, deploymentTarget)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), deploymentTarget, new(DeploymentTarget), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentTarget), nil
}

// Get returns a collection of machines based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *MachineService) Get(machinesQuery MachinesQuery) (*resources.Resources[DeploymentTarget], error) {
	path, err := s.GetURITemplate().Expand(machinesQuery)
	if err != nil {
		return &resources.Resources[DeploymentTarget]{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(resources.Resources[DeploymentTarget]), path)
	if err != nil {
		return &resources.Resources[DeploymentTarget]{}, err
	}

	return response.(*resources.Resources[DeploymentTarget]), nil
}

// GetByID returns the machine that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *MachineService) GetByID(id string) (*DeploymentTarget, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(DeploymentTarget), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentTarget), nil
}

// GetAll returns all machines. If none can be found or an error occurs, it
// returns an empty collection.
func (s *MachineService) GetAll() ([]*DeploymentTarget, error) {
	items := []*DeploymentTarget{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Machine with a matching name.
func (s *MachineService) GetByName(name string) ([]*DeploymentTarget, error) {
	if internal.IsEmpty(name) {
		return []*DeploymentTarget{}, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	path, err := services.GetByNamePath(s, name)
	if err != nil {
		return []*DeploymentTarget{}, err
	}

	return services.GetPagedResponse[DeploymentTarget](s, path)
}

// GetByPartialName performs a lookup and returns the machine with a matching
// partial name.
func (s *MachineService) GetByPartialName(partialName string) ([]*DeploymentTarget, error) {
	if internal.IsEmpty(partialName) {
		return []*DeploymentTarget{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*DeploymentTarget{}, err
	}

	return services.GetPagedResponse[DeploymentTarget](s, path)
}

// Update updates an existing machine in Octopus Deploy
func (s *MachineService) Update(resource *DeploymentTarget) (*DeploymentTarget, error) {
	path, err := services.GetUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(DeploymentTarget), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentTarget), nil
}

var _ services.IService = &MachineService{}
