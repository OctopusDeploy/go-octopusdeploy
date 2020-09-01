package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TenantService struct {
	sling *sling.Sling
	path  string
}

func NewTenantService(sling *sling.Sling) *TenantService {
	return &TenantService{
		sling: sling,
		path:  "tenants",
	}
}

// Get returns a single tenant by its tenantid in Octopus Deploy
func (s *TenantService) Get(id string) (*model.Tenant, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// GetAll returns all tenants in Octopus Deploy
func (s *TenantService) GetAll() (*[]model.Tenant, error) {
	var t []model.Tenant
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Tenants), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Tenants)
		t = append(t, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &t, nil
}

// GetByName gets an existing Tenant by its name in Octopus Deploy
func (s *TenantService) GetByName(name string) (*model.Tenant, error) {
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

// Add adds a new Tenant in Octopus Deploy
func (s *TenantService) Add(resource *model.Tenant) (*model.Tenant, error) {
	err := model.ValidateTenantValues(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Tenant), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// Delete deletes an existing tenant in Octopus Deploy
func (s *TenantService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing tenant in Octopus Deploy
func (s *TenantService) Update(resource *model.Tenant) (*model.Tenant, error) {
	err := model.ValidateTenantValues(resource)
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
