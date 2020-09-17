package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type MachineService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewMachineService(sling *sling.Sling, uriTemplate string) *MachineService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &MachineService{
		name:  "MachineService",
		path:  path,
		sling: sling,
	}
}

// Get returns a single machine with a given ID.
func (s *MachineService) Get(id string) (*model.Machine, error) {
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
	if isEmpty(name) {
		return nil, createInvalidParameterError("GetByName", "name")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
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

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new Machine.
func (s *MachineService) Add(machine *model.Machine) (*model.Machine, error) {
	if machine == nil {
		return nil, createInvalidParameterError("Add", "machine")
	}

	err := machine.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

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
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
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
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &MachineService{}
