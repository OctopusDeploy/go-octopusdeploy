package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type machineService struct {
	name        string                    `validate:"required"`
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

func (s machineService) getPagedResponse(path string) ([]model.Machine, error) {
	resources := []model.Machine{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Machines), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Machines)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s machineService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new machine.
func (s machineService) Add(resource *model.Machine) (*model.Machine, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Machine), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

// DeleteByID deletes the machine that matches the input ID.
func (s machineService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("machine", "ID", id)
	}

	return err
}

// GetByID returns the machine that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s machineService) GetByID(id string) (*model.Machine, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Machine), path)
	if err != nil {
		return nil, createResourceNotFoundError("machine", "ID", id)
	}

	return resp.(*model.Machine), nil
}

// GetAll returns all machines. If none can be found or an error occurs, it
// returns an empty collection.
func (s machineService) GetAll() ([]model.Machine, error) {
	items := []model.Machine{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Machine with a matching name.
func (s machineService) GetByName(name string) ([]model.Machine, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []model.Machine{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns the machine with a matching
// partial name.
func (s machineService) GetByPartialName(name string) ([]model.Machine, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.Machine{}, err
	}

	return s.getPagedResponse(path)
}

// Update updates an existing machine in Octopus Deploy
func (s machineService) Update(resource model.Machine) (*model.Machine, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Machine), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Machine), nil
}

var _ ServiceInterface = &machineService{}
