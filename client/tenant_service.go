package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TenantService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewTenantService(sling *sling.Sling) *TenantService {
	if sling == nil {
		return nil
	}

	return &TenantService{
		sling: sling,
		path:  "tenants",
	}
}

// Get returns a single tenant by its tenantid in Octopus Deploy
func (s *TenantService) Get(id string) (*model.Tenant, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("TenantService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// GetAll returns all instances of a Tenant.
func (s *TenantService) GetAll() (*[]model.Tenant, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Tenant), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Tenant), nil
}

// GetByName performs a lookup and returns the Tenant with a matching name.
func (s *TenantService) GetByName(name string) (*model.Tenant, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("TenantService: invalid parameter, name")
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

// Add creates a new Tenant.
func (s *TenantService) Add(tenant *model.Tenant) (*model.Tenant, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if tenant == nil {
		return nil, errors.New("TenantService: invalid parameter, tenant")
	}

	err = tenant.Validate()

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
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("TenantService: invalid parameter, ID")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing tenant in Octopus Deploy
func (s *TenantService) Update(resource *model.Tenant) (*model.Tenant, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("TenantService: invalid parameter, resource")
	}

	err = resource.Validate()

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
		return fmt.Errorf("TenantService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("TenantService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &TenantService{}
