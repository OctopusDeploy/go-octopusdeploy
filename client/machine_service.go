package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type MachineService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewMachineService(sling *sling.Sling) *MachineService {
	if sling == nil {
		return nil
	}

	return &MachineService{
		sling: sling,
		path:  "machines",
	}
}

// Get returns a single machine with a given ID.
func (s *MachineService) Get(id string) (*model.Machine, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("MachineService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// GetAll returns all instances of a Machine.
func (s *MachineService) GetAll() (*[]model.Machine, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	var p []model.Machine
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Machines), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Machines)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}
	return &p, nil
}

// GetByName performs a lookup and returns the Machine with a matching name.
func (s *MachineService) GetByName(name string) (*model.Machine, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("MachineService: invalid parameter, name")
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

// Add creates a new Machine.
func (s *MachineService) Add(machine *model.Machine) (*model.Machine, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if machine == nil {
		return nil, errors.New("MachineService: invalid parameter, machine")
	}

	err = machine.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, machine, new(model.Machine), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// Delete deletes an existing machine in Octopus Deploy
func (s *MachineService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("MachineService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing machine in Octopus Deploy
func (s *MachineService) Update(machine *model.Machine) (*model.Machine, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = machine.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", machine.ID)
	resp, err := apiUpdate(s.sling, machine, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

func (s *MachineService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("MachineService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("MachineService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &MachineService{}
