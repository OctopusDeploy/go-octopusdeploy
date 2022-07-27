package machines

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type MachinePolicyService struct {
	templatePath string

	services.CanDeleteService
}

func NewMachinePolicyService(sling *sling.Sling, uriTemplate string, templatePath string) *MachinePolicyService {
	return &MachinePolicyService{
		templatePath: templatePath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceMachinePolicyService, sling, uriTemplate),
		},
	}
}

// Add creates a new machine policy.
func (s *MachinePolicyService) Add(machinePolicy *MachinePolicy) (*MachinePolicy, error) {
	if IsNil(machinePolicy) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterMachinePolicy)
	}

	path, err := services.GetAddPath(s, machinePolicy)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), machinePolicy, new(MachinePolicy), path)
	if err != nil {
		return nil, err
	}

	return resp.(*MachinePolicy), nil
}

// Get returns a collection of machine policies based on the criteria defined
// by its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s *MachinePolicyService) Get(machinePoliciesQuery MachinePoliciesQuery) (*resources.Resources[MachinePolicy], error) {
	path, err := s.GetURITemplate().Expand(machinePoliciesQuery)
	if err != nil {
		return &resources.Resources[MachinePolicy]{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(resources.Resources[MachinePolicy]), path)
	if err != nil {
		return &resources.Resources[MachinePolicy]{}, err
	}

	return response.(*resources.Resources[MachinePolicy]), nil
}

// GetAll returns all machine policies. If none can be found or an error
// occurs, it returns an empty collection.
func (s *MachinePolicyService) GetAll() ([]*MachinePolicy, error) {
	items := []*MachinePolicy{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the machine policy that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s *MachinePolicyService) GetByID(id string) (*MachinePolicy, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(MachinePolicy), path)
	if err != nil {
		return nil, err
	}

	return resp.(*MachinePolicy), nil
}

// GetByPartialName performs a lookup and returns machine policies with a
// matching partial name.
func (s *MachinePolicyService) GetByPartialName(partialName string) ([]*MachinePolicy, error) {
	if internal.IsEmpty(partialName) {
		return []*MachinePolicy{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*MachinePolicy{}, err
	}

	return services.GetPagedResponse[MachinePolicy](s, path)
}

func (s *MachinePolicyService) GetTemplate() (*MachinePolicy, error) {
	resp, err := services.ApiGet(s.GetClient(), new(MachinePolicy), s.templatePath)
	if err != nil {
		return nil, err
	}

	return resp.(*MachinePolicy), nil
}

// Update modifies a machine policy based on the one provided as input.
func (s *MachinePolicyService) Update(machinePolicy *MachinePolicy) (*MachinePolicy, error) {
	path, err := services.GetUpdatePath(s, machinePolicy)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), machinePolicy, new(MachinePolicy), path)
	if err != nil {
		return nil, err
	}

	return resp.(*MachinePolicy), nil
}
