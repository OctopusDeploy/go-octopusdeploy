package tenants

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type TenantService struct {
	missingVariablesPath string
	statusPath           string
	tagTestPath          string

	services.CanDeleteService
}

func NewTenantService(sling *sling.Sling, uriTemplate string, missingVariablesPath string, statusPath string, tagTestPath string) *TenantService {
	return &TenantService{
		missingVariablesPath: missingVariablesPath,
		statusPath:           statusPath,
		tagTestPath:          tagTestPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceTenantService, sling, uriTemplate),
		},
	}
}

func (s *TenantService) getByProjectIDPath(id string) (string, error) {
	if internal.IsEmpty(id) {
		return "", internal.CreateInvalidParameterError(constants.OperationGetByProjectID, "id")
	}

	err := services.ValidateInternalState(s)
	if err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	values[constants.ParameterProjectID] = id

	return s.GetURITemplate().Expand(values)
}

func (s *TenantService) getPagedResponse(path string) ([]*Tenant, error) {
	resources := []*Tenant{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(Tenants), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Tenants)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new Tenant.
func (s *TenantService) Add(tenant *Tenant) (*Tenant, error) {
	if IsNil(tenant) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterTenant)
	}

	path, err := services.GetAddPath(s, tenant)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), tenant, new(Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

func (s *TenantService) CreateVariables(tenant *Tenant, tenantVariable *variables.TenantVariables) (*variables.TenantVariables, error) {
	resp, err := services.ApiAdd(s.GetClient(), tenantVariable, new(variables.TenantVariables), tenant.Links["Variables"])
	if err != nil {
		return nil, err
	}

	return resp.(*variables.TenantVariables), nil
}

// Get returns a collection of tenants based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *TenantService) Get(tenantsQuery TenantsQuery) (*Tenants, error) {
	path, err := s.GetURITemplate().Expand(tenantsQuery)
	if err != nil {
		return &Tenants{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(Tenants), path)
	if err != nil {
		return &Tenants{}, err
	}

	return response.(*Tenants), nil
}

// GetAll returns all tenants. If none can be found or an error occurs, it
// returns an empty collection.
func (s *TenantService) GetAll() ([]*Tenant, error) {
	items := []*Tenant{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the tenant that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *TenantService) GetByID(id string) (*Tenant, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s *TenantService) GetByIDs(ids []string) ([]*Tenant, error) {
	if len(ids) == 0 {
		return []*Tenant{}, nil
	}

	path, err := services.GetByIDsPath(s, ids)
	if err != nil {
		return []*Tenant{}, err
	}

	return s.getPagedResponse(path)
}

func (s *TenantService) GetMissingVariables(missibleVariablesQuery variables.MissingVariablesQuery) (*[]variables.TenantsMissingVariables, error) {
	template, err := uritemplates.Parse(s.missingVariablesPath)
	if err != nil {
		return &[]variables.TenantsMissingVariables{}, err
	}

	path, err := template.Expand(missibleVariablesQuery)
	if err != nil {
		return &[]variables.TenantsMissingVariables{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new([]variables.TenantsMissingVariables), path)
	if err != nil {
		return &[]variables.TenantsMissingVariables{}, err
	}

	return response.(*[]variables.TenantsMissingVariables), nil
}

// GetByProjectID performs a lookup and returns all tenants with a matching
// project ID.
func (s *TenantService) GetByProjectID(id string) ([]*Tenant, error) {
	path, err := s.getByProjectIDPath(id)
	if err != nil {
		return []*Tenant{}, nil
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns all tenants with a matching
// partial name.
func (s *TenantService) GetByPartialName(partialName string) ([]*Tenant, error) {
	if internal.IsEmpty(partialName) {
		return []*Tenant{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Tenant{}, nil
	}

	return s.getPagedResponse(path)
}

func (s *TenantService) GetVariables(tenant *Tenant) (*variables.TenantVariables, error) {
	resp, err := services.ApiGet(s.GetClient(), new(variables.TenantVariables), tenant.Links["Variables"])
	if err != nil {
		return nil, err
	}

	return resp.(*variables.TenantVariables), nil
}

// Update modifies a tenant based on the one provided as input.
func (s *TenantService) Update(resource *Tenant) (*Tenant, error) {
	path, err := services.GetUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

func (s *TenantService) UpdateVariables(tenant *Tenant, tenantVariables *variables.TenantVariables) (*variables.TenantVariables, error) {
	resp, err := services.ApiPost(s.GetClient(), tenantVariables, new(variables.TenantVariables), tenant.Links["Variables"])
	if err != nil {
		return nil, err
	}

	return resp.(*variables.TenantVariables), nil
}
