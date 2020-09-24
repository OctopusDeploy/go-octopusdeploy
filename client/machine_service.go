package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type machineService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newMachineService(sling *sling.Sling, uriTemplate string) *machineService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &machineService{
		name:        serviceMachineService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s machineService) getClient() *sling.Sling {
	return s.sling
}

func (s machineService) getName() string {
	return s.name
}

func (s machineService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns a single machine with a given ID.
func (s machineService) GetByID(id string) (*model.Machine, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Machine), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// GetAll returns all instances of a Machine. If none can be found or an error occurs, it returns an empty collection.
func (s machineService) GetAll() ([]model.Machine, error) {
	items := new([]model.Machine)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

// GetByName performs a lookup and returns the Machine with a matching name.
func (s machineService) GetByName(name string) (*model.Machine, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError(operationGetByName, parameterName)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, operationGetByName, name)
}

// Add creates a new Machine.
func (s machineService) Add(machine *model.Machine) (*model.Machine, error) {
	if machine == nil {
		return nil, createInvalidParameterError(operationAdd, "machine")
	}

	err := machine.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), machine, new(model.Machine), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// Delete deletes an existing machine in Octopus Deploy
func (s machineService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// Update updates an existing machine in Octopus Deploy
func (s machineService) Update(machine *model.Machine) (*model.Machine, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	err = machine.Validate()

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", machine.ID)

	resp, err := apiUpdate(s.getClient(), machine, new(model.Machine), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

var _ ServiceInterface = &machineService{}
