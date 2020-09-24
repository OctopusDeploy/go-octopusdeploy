package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type tenantService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newTenantService(sling *sling.Sling, uriTemplate string) *tenantService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &tenantService{
		name:        serviceTenantService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s tenantService) getClient() *sling.Sling {
	return s.sling
}

func (s tenantService) getName() string {
	return s.name
}

func (s tenantService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns a single tenant by its tenantid in Octopus Deploy
func (s tenantService) GetByID(id string) (*model.Tenant, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// GetAll returns all instances of a Tenant. If none can be found or an error occurs, it returns an empty collection.
func (s tenantService) GetAll() ([]model.Tenant, error) {
	items := new([]model.Tenant)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

// GetByName performs a lookup and returns the Tenant with a matching name.
func (s tenantService) GetByName(name string) (*model.Tenant, error) {
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

// Add creates a new Tenant.
func (s tenantService) Add(tenant *model.Tenant) (*model.Tenant, error) {
	if tenant == nil {
		return nil, createInvalidParameterError(operationAdd, "tenant")
	}

	err := tenant.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), tenant, new(model.Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// Delete deletes an existing tenant in Octopus Deploy
func (s tenantService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// Update updates an existing tenant in Octopus Deploy
func (s tenantService) Update(resource *model.Tenant) (*model.Tenant, error) {
	if resource == nil {
		return nil, createInvalidParameterError(operationUpdate, "resource")
	}

	err := resource.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", resource.ID)

	resp, err := apiUpdate(s.getClient(), resource, new(model.Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

var _ ServiceInterface = &tenantService{}
