package octopusdeploy

import (
	"fmt"
	"time"

	"github.com/dghubble/sling"
)

type TenantService struct {
	sling *sling.Sling
}

func NewTenantService(sling *sling.Sling) *TenantService {
	return &TenantService{
		sling: sling,
	}
}

type Tenants struct {
	Items []Tenant `json:"Items"`
	PagedResults
}

type Tenant struct {
	ID                  string              `json:"Id"`
	Name                string              `json:"Name"`
	TenantTags          []string            `json:"TenantTags,omitempty"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments,omitempty"`
	SpaceID             string              `json:"SpaceId"`
	ClonedFromTenantID  string              `json:"ClonedFromTenantId"`
	Description         string              `json:"Description"`
	LastModifiedOn      time.Time           `json:"LastModifiedOn"`
	LastModifiedBy      string              `json:"LastModifiedBy"`
	Links               map[string]string   `json:"Links"`
}

func NewTenant(name, description string) *Tenant {
	return &Tenant{
		Name:        name,
		Description: description,
	}
}

func ValidateTenantValues(tenant *Tenant) error {
	return ValidateMultipleProperties([]error{
		ValidateRequiredPropertyValue("Name", tenant.Name),
	})
}

// Get returns a single tenant by its tenantid in Octopus Deploy
func (s *TenantService) Get(tenantId string) (*Tenant, error) {
	path := fmt.Sprintf("tenants/%s", tenantId)
	resp, err := apiGet(s.sling, new(Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

// GetAll returns all tenants in Octopus Deploy
func (s *TenantService) GetAll() (*[]Tenant, error) {
	var t []Tenant

	path := "tenants"

	loadNextPage := true
	for loadNextPage {
		resp, err := apiGet(s.sling, new(Tenants), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Tenants)

		t = append(t, r.Items...)

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &t, nil
}

// GetByName gets an existing Tenant by its name in Octopus Deploy
func (s *TenantService) GetByName(tenantName string) (*Tenant, error) {
	var foundTenant Tenant
	tenants, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, tenant := range *tenants {
		if tenant.Name == tenantName {
			return &tenant, nil
		}
	}

	return &foundTenant, fmt.Errorf("no tenant found with tenant name %s", tenantName)
}

// Add adds an new tenant in Octopus Deploy
func (s *TenantService) Add(tenant *Tenant) (*Tenant, error) {
	err := ValidateTenantValues(tenant)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, tenant, new(Tenant), "tenants")

	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

// Delete deletes an existing tenant in Octopus Deploy
func (s *TenantService) Delete(tenantId string) error {
	path := fmt.Sprintf("tenants/%s", tenantId)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing tenant in Octopus Deploy
func (s *TenantService) Update(tenant *Tenant) (*Tenant, error) {
	err := ValidateTenantValues(tenant)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("tenants/%s", tenant.ID)
	resp, err := apiUpdate(s.sling, tenant, new(Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}
