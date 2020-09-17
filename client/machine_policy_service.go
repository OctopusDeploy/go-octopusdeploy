package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type MachinePolicyService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewMachinePolicyService(sling *sling.Sling, uriTemplate string) *MachinePolicyService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &MachinePolicyService{
		name:  "MachinePolicyService",
		path:  path,
		sling: sling,
	}
}

// Get returns a single machine with a given MachineID
func (s *MachinePolicyService) Get(id string) (*model.MachinePolicy, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}

// GetAll returns all instances of a MachinePolicy.
func (s *MachinePolicyService) GetAll() ([]model.MachinePolicy, error) {
	err := s.validateInternalState()

	items := new([]model.MachinePolicy)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

func (s *MachinePolicyService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &MachinePolicyService{}
