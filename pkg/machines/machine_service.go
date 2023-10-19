package machines

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
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
//
// Deprecated: use machines.Add
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
//
// Deprecated: use machines.Get
func (s *MachineService) Get(machinesQuery MachinesQuery) (*resources.Resources[*DeploymentTarget], error) {
	path, err := s.GetURITemplate().Expand(machinesQuery)
	if err != nil {
		return &resources.Resources[*DeploymentTarget]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*DeploymentTarget]), path)
	if err != nil {
		return &resources.Resources[*DeploymentTarget]{}, err
	}

	return response.(*resources.Resources[*DeploymentTarget]), nil
}

// GetByID returns the machine that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: use machines.GetByID
func (s *MachineService) GetByID(id string) (*DeploymentTarget, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(DeploymentTarget), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentTarget), nil
}

// GetAll returns all machines. If none can be found or an error occurs, it
// returns an empty collection.
//
// Deprecated: use machines.GetAll
func (s *MachineService) GetAll() ([]*DeploymentTarget, error) {
	items := []*DeploymentTarget{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
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

func (s *MachineService) GetByIdentifier(identifier string) (*DeploymentTarget, error) {
	// the machines endpoint doesn't currently support slugs
	target, err := s.GetByID(identifier)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if target != nil {
			return target, nil
		}
	}

	possibleTargets, err := s.GetByName(identifier)
	if err != nil {
		return nil, err
	}

	for _, t := range possibleTargets {
		if strings.EqualFold(identifier, t.Name) {
			return t, nil
		}
	}

	return nil, fmt.Errorf("cannot find machine with the name or ID of '%s'", identifier)
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
//
// Deprecated: use machines.Update
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

// --- NEW ---

const template = "/api/{spaceId}/machines{/id}{?skip,take,name,ids,partialName,roles,isDisabled,healthStatuses,commStyles,tenantIds,tenantTags,environmentIds,thumbprint,deploymentId,shellNames}"

// Add creates a new machine.
func Add(client newclient.Client, deploymentTarget *DeploymentTarget) (*DeploymentTarget, error) {
	return newclient.Add[DeploymentTarget](client, template, deploymentTarget.SpaceID, deploymentTarget)
}

// Get returns a collection of machines based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func Get(client newclient.Client, spaceID string, machinesQuery MachinesQuery) (*resources.Resources[*DeploymentTarget], error) {
	return newclient.GetByQuery[DeploymentTarget](client, template, spaceID, machinesQuery)
}

// GetByID returns the machine that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*DeploymentTarget, error) {
	return newclient.GetByID[DeploymentTarget](client, template, spaceID, ID)
}

// Update updates an existing machine in Octopus Deploy
func Update(client newclient.Client, spaceID string, deploymentTarget *DeploymentTarget) (*DeploymentTarget, error) {
	return newclient.Update[DeploymentTarget](client, template, spaceID, deploymentTarget.ID, deploymentTarget)
}

// DeleteById deletes the machine based on the ID provided.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetAll returns all machines. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*DeploymentTarget, error) {
	return newclient.GetAll[DeploymentTarget](client, template, spaceID)
}
