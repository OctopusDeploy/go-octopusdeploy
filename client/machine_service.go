package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type MachineService struct {
	sling *sling.Sling
	path  string
}

func NewMachineService(sling *sling.Sling) *MachineService {
	return &MachineService{
		sling: sling,
		path:  "machines",
	}
}

// Get returns a single machine with a given ID.
func (s *MachineService) Get(id string) (*model.Machine, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// GetAll returns all registered machines
func (s *MachineService) GetAll() (*[]model.Machine, error) {
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

// GetByName gets an existing machine by its name in Octopus Deploy
func (s *MachineService) GetByName(name string) (*model.Machine, error) {
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

// Add creates a new machine in Octopus Deploy
func (s *MachineService) Add(resource *model.Machine) (*model.Machine, error) {
	err := resource.Validate()
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Machine), s.path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// Delete deletes an existing machine in Octopus Deploy
func (s *MachineService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing machine in Octopus Deploy
func (s *MachineService) Update(resource *model.Machine) (*model.Machine, error) {
	err := resource.Validate()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}
