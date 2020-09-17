package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TenantService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewTenantService(sling *sling.Sling, uriTemplate string) *TenantService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &TenantService{
		name:  "TenantService",
		path:  path,
		sling: sling,
	}
}

// Get returns a single tenant by its tenantid in Octopus Deploy
func (s *TenantService) Get(id string) (*model.Tenant, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// GetAll returns all instances of a Tenant.
func (s *TenantService) GetAll() ([]model.Tenant, error) {
	err := s.validateInternalState()

	items := new([]model.Tenant)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the Tenant with a matching name.
func (s *TenantService) GetByName(name string) (*model.Tenant, error) {
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

	for _, item := range collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new Tenant.
func (s *TenantService) Add(tenant *model.Tenant) (*model.Tenant, error) {
	if tenant == nil {
		return nil, createInvalidParameterError("Add", "tenant")
	}

	err := tenant.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, tenant, new(model.Tenant), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// Delete deletes an existing tenant in Octopus Deploy
func (s *TenantService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing tenant in Octopus Deploy
func (s *TenantService) Update(resource *model.Tenant) (*model.Tenant, error) {
	if resource == nil {
		return nil, createInvalidParameterError("Update", "resource")
	}

	err := resource.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

func (s *TenantService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &TenantService{}
