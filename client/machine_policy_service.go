package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type machinePolicyService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newMachinePolicyService(sling *sling.Sling, uriTemplate string) *machinePolicyService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &machinePolicyService{
		name:        serviceMachinePolicyService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s machinePolicyService) getClient() *sling.Sling {
	return s.sling
}

func (s machinePolicyService) getName() string {
	return s.name
}

func (s machinePolicyService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns a single machine with a given MachineID
func (s machinePolicyService) GetByID(id string) (*model.MachinePolicy, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Machine), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}

// GetAll returns all instances of a MachinePolicy. If none can be found or an error occurs, it returns an empty collection.
func (s machinePolicyService) GetAll() ([]model.MachinePolicy, error) {
	items := new([]model.MachinePolicy)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

var _ ServiceInterface = &machinePolicyService{}
