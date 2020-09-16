package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type MachinePolicyService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewMachinePolicyService(sling *sling.Sling, uriTemplate string) *MachinePolicyService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &MachinePolicyService{
		sling: sling,
		path:  path,
	}
}

// Get returns a single machine with a given MachineID
func (s *MachinePolicyService) Get(id string) (*model.MachinePolicy, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("LifecycleService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}

// GetAll returns all instances of a MachinePolicy.
func (s *MachinePolicyService) GetAll() (*[]model.MachinePolicy, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.MachinePolicy), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.MachinePolicy), nil
}

func (s *MachinePolicyService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("MachinePolicyService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("MachinePolicyService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &MachinePolicyService{}
