package userroles

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type UserRoleService struct {
	services.CanDeleteService
}

const (
	userRolesTemplate = "/api/userroles{/id}{?skip,take,ids,partialName}"
)

func NewUserRoleService(sling *sling.Sling, uriTemplate string) *UserRoleService {
	return &UserRoleService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceUserRoleService, sling, uriTemplate),
		},
	}
}

// Add creates a new user role.
//
// Deprecated: Use userroles.Add
func (s *UserRoleService) Add(userRole *UserRole) (*UserRole, error) {
	if IsNil(userRole) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterUserRole)
	}

	path, err := services.GetAddPath(s, userRole)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), userRole, new(UserRole), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserRole), nil
}

// Get returns a collection of user roles based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
//
// Deprecated: Use userroles.Get
func (s *UserRoleService) Get(userRolesQuery UserRolesQuery) (*resources.Resources[*UserRole], error) {
	path, err := s.GetURITemplate().Expand(userRolesQuery)
	if err != nil {
		return &resources.Resources[*UserRole]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*UserRole]), path)
	if err != nil {
		return &resources.Resources[*UserRole]{}, err
	}

	return response.(*resources.Resources[*UserRole]), nil
}

// GetAll returns all user roles. If none can be found or an error occurs, it
// returns an empty collection.
func (s *UserRoleService) GetAll() ([]*UserRole, error) {
	items := []*UserRole{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the user role that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: Use userroles.GetByID
func (s *UserRoleService) GetByID(id string) (*UserRole, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(UserRole), path)
	if err != nil {
		return nil, internal.CreateResourceNotFoundError("userRole", "ID", id)
	}

	return resp.(*UserRole), nil
}

// Update modifies a user role based on the one provided as input.
//
// Deprecated: Use userroles.Update
func (s *UserRoleService) Update(userRole *UserRole) (*UserRole, error) {
	if userRole == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUserRole)
	}

	path, err := services.GetUpdatePath(s, userRole)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), userRole, new(UserRole), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserRole), nil
}

// ----- new -----

// Add creates a new user role.
func Add(client newclient.Client, userRole *UserRole) (*UserRole, error) {
	if IsNil(userRole) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterUserRole)
	}

	expandedUri, err := client.URITemplateCache().Expand(userRolesTemplate, map[string]any{
		"id": userRole.ID,
	})
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Post[UserRole](client.HttpSession(), expandedUri, userRole)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get returns a collection of user roles based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func Get(client newclient.Client, userRolesQuery UserRolesQuery) (*resources.Resources[*UserRole], error) {
	values, _ := uritemplates.Struct2map(userRolesQuery)
	if values == nil {
		values = map[string]any{}
	}

	expandedUri, err := client.URITemplateCache().Expand(userRolesTemplate, values)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[resources.Resources[*UserRole]](client.HttpSession(), expandedUri)
	if err != nil {
		return &resources.Resources[*UserRole]{}, err
	}

	return resp, nil
}

// GetByID returns the user role that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, id string) (*UserRole, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	expandedUri, err := client.URITemplateCache().Expand(userRolesTemplate, map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[UserRole](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Update modifies a user role based on the one provided as input.
func Update(client newclient.Client, userRole *UserRole) (*UserRole, error) {
	if userRole == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUserRole)
	}

	expandedUri, err := client.URITemplateCache().Expand(userRolesTemplate, map[string]any{
		"id": userRole.ID,
	})
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Put[UserRole](client.HttpSession(), expandedUri, userRole)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteByID deletes the resource that matches the space ID and input ID.
func DeleteByID(client newclient.Client, id string) error {
	if internal.IsEmpty(id) {
		return internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID)
	}

	expandedUri, err := client.URITemplateCache().Expand(userRolesTemplate, map[string]any{
		"id": id,
	})
	if err != nil {
		return err
	}

	return newclient.Delete(client.HttpSession(), expandedUri)
}
