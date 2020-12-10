package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type tenantService struct {
	missingVariablesPath string
	statusPath           string
	tagTestPath          string

	canDeleteService
}

func newTenantService(sling *sling.Sling, uriTemplate string, missingVariablesPath string, statusPath string, tagTestPath string) *tenantService {
	tenantService := &tenantService{
		missingVariablesPath: missingVariablesPath,
		statusPath:           statusPath,
		tagTestPath:          tagTestPath,
	}
	tenantService.service = newService(ServiceTenantService, sling, uriTemplate)

	return tenantService
}

func (s tenantService) getByProjectIDPath(id string) (string, error) {
	if isEmpty(id) {
		return emptyString, createInvalidParameterError(OperationGetByProjectID, ParameterID)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[ParameterProjectID] = id

	return s.getURITemplate().Expand(values)
}

func (s tenantService) getPagedResponse(path string) ([]*Tenant, error) {
	resources := []*Tenant{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Tenants), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Tenants)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new Tenant.
func (s tenantService) Add(resource *Tenant) (*Tenant, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

// Get returns a collection of tenants based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s tenantService) Get(tenantsQuery TenantsQuery) (*Tenants, error) {
	path, err := s.getURITemplate().Expand(tenantsQuery)
	if err != nil {
		return &Tenants{}, err
	}

	response, err := apiGet(s.getClient(), new(Tenants), path)
	if err != nil {
		return &Tenants{}, err
	}

	return response.(*Tenants), nil
}

// GetAll returns all tenants. If none can be found or an error occurs, it
// returns an empty collection.
func (s tenantService) GetAll() ([]*Tenant, error) {
	items := []*Tenant{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the tenant that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s tenantService) GetByID(id string) (*Tenant, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s tenantService) GetByIDs(ids []string) ([]*Tenant, error) {
	if len(ids) == 0 {
		return []*Tenant{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*Tenant{}, err
	}

	return s.getPagedResponse(path)
}

// GetByProjectID performs a lookup and returns all tenants with a matching
// project ID.
func (s tenantService) GetByProjectID(id string) ([]*Tenant, error) {
	path, err := s.getByProjectIDPath(id)
	if err != nil {
		return []*Tenant{}, nil
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns all tenants with a matching
// partial name.
func (s tenantService) GetByPartialName(name string) ([]*Tenant, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Tenant{}, nil
	}

	return s.getPagedResponse(path)
}

// Update modifies a tenant based on the one provided as input.
func (s tenantService) Update(resource *Tenant) (*Tenant, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}
