package tenants

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
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

// Add creates a new Tenant.
//
// Deprecated: use tenants.Add
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

func (s *TenantService) Clone(sourceTenant *Tenant, request TenantCloneRequest) (*Tenant, error) {
	path, err := s.GetURITemplate().Expand(&TenantCloneQuery{CloneTenantID: sourceTenant.GetID()})
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiPost(s.GetClient(), request, new(Tenant), path)
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
//
// Deprecated: use tenants.Get
func (s *TenantService) Get(tenantsQuery TenantsQuery) (*resources.Resources[*Tenant], error) {
	path, err := s.GetURITemplate().Expand(tenantsQuery)
	if err != nil {
		return &resources.Resources[*Tenant]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Tenant]), path)
	if err != nil {
		return &resources.Resources[*Tenant]{}, err
	}

	return response.(*resources.Resources[*Tenant]), nil
}

// GetAll returns all tenants. If none can be found or an error occurs, it
// returns an empty collection.
//
// Deprecated: use tenants.GetAll
func (s *TenantService) GetAll() ([]*Tenant, error) {
	items := []*Tenant{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the tenant that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: use tenants.GetByID
func (s *TenantService) GetByID(id string) (*Tenant, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Tenant), path)
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

	return services.GetPagedResponse[Tenant](s, path)
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

	response, err := api.ApiGet(s.GetClient(), new([]variables.TenantsMissingVariables), path)
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

	return services.GetPagedResponse[Tenant](s, path)
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

	return services.GetPagedResponse[Tenant](s, path)
}

func (s *TenantService) GetByName(name string) (*Tenant, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	tenants, err := s.Get(TenantsQuery{
		PartialName: name,
	})
	if err != nil {
		return nil, err
	}

	for _, tenant := range tenants.Items {
		if strings.EqualFold(tenant.Name, name) {
			return tenant, nil
		}
	}

	return nil, services.ErrItemNotFound
}

func (s *TenantService) GetByIdentifier(identifier string) (*Tenant, error) {
	tenant, err := s.GetByID(identifier)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if tenant != nil {
			return tenant, nil
		}
	}

	return s.GetByName(identifier)
}

func (s *TenantService) GetVariables(tenant *Tenant) (*variables.TenantVariables, error) {
	resp, err := api.ApiGet(s.GetClient(), new(variables.TenantVariables), tenant.Links["Variables"])
	if err != nil {
		return nil, err
	}

	return resp.(*variables.TenantVariables), nil
}

// Update modifies a tenant based on the one provided as input.
//
// Deprecated: use tenant.Update
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

// --- new ---

const template = "/api/{spaceId}/tenants{/id}{?skip,projectId,name,tags,take,ids,clone,partialName,isDisabled,clonedFromTenantId}"

// Get returns a collection of tenants based on the criteria defined by its
// input query parameter.
func Get(client newclient.Client, spaceID string, tenantsQuery TenantsQuery) (*resources.Resources[*Tenant], error) {
	return newclient.GetByQuery[Tenant](client, template, spaceID, tenantsQuery)
}

// Update modifies a tenant based on the one provided as input.
func Update(client newclient.Client, resource *Tenant) (*Tenant, error) {
	return newclient.Update[Tenant](client, template, resource.SpaceID, resource.ID, resource)
}

// Add creates a new Tenant.
func Add(client newclient.Client, tenant *Tenant) (*Tenant, error) {
	return newclient.Add[Tenant](client, template, tenant.SpaceID, tenant)
}

// GetByID returns the tenant that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*Tenant, error) {
	return newclient.GetByID[Tenant](client, template, spaceID, ID)
}

// DeleteByID deletes the tenant that matches the input ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetAll returns all tenants. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*Tenant, error) {
	return newclient.GetAll[Tenant](client, template, spaceID)
}

const tenantProjectVariableTemplate = "/api/{spaceId}/tenants/{id}/projectvariables?IncludeMissingVariables={includeMissingVariables}"
const tenantCommonVariableTemplate = "/api/{spaceId}/tenants/{id}/commonvariables?IncludeMissingVariables={includeMissingVariables}"

// GetProjectVariables returns all tenant project variables. If an error occurs, it returns nil.
func GetProjectVariables(client newclient.Client, query variables.GetTenantProjectVariablesQuery) (*variables.GetTenantProjectVariablesResponse, error) {
	return newclient.GetResourceByQuery[variables.GetTenantProjectVariablesResponse](client, tenantProjectVariableTemplate, query)
}

// GetCommonVariables returns all tenant pagination variables. If an error occurs, it returns nil.
func GetCommonVariables(client newclient.Client, query variables.GetTenantCommonVariablesQuery) (*variables.GetTenantCommonVariablesResponse, error) {
	return newclient.GetResourceByQuery[variables.GetTenantCommonVariablesResponse](client, tenantCommonVariableTemplate, query)
}

// UpdateProjectVariables modifies tenant project variables based on the ones provided as input.
func UpdateProjectVariables(client newclient.Client, spaceID string, tenantID string, projectVariables *variables.ModifyTenantProjectVariablesCommand) (*variables.ModifyTenantProjectVariablesResponse, error) {
	return newclient.Update[variables.ModifyTenantProjectVariablesResponse](client, tenantProjectVariableTemplate, spaceID, tenantID, projectVariables)
}

// UpdateCommonVariables modifies tenant pagination variables based on the ones provided as input.
func UpdateCommonVariables(client newclient.Client, spaceID string, tenantID string, commonVariables *variables.ModifyTenantCommonVariablesCommand) (*variables.ModifyTenantCommonVariablesResponse, error) {
	return newclient.Update[variables.ModifyTenantCommonVariablesResponse](client, tenantCommonVariableTemplate, spaceID, tenantID, commonVariables)
}
