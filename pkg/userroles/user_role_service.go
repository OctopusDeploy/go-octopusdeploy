package userroles

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type UserRoleService struct {
	services.CanDeleteService
}

func NewUserRoleService(sling *sling.Sling, uriTemplate string) *UserRoleService {
	return &UserRoleService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceUserRoleService, sling, uriTemplate),
		},
	}
}

// Add creates a new user role.
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
func (s *UserRoleService) Get(userRolesQuery UserRolesQuery) (*UserRoles, error) {
	path, err := s.GetURITemplate().Expand(userRolesQuery)
	if err != nil {
		return &UserRoles{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(UserRoles), path)
	if err != nil {
		return &UserRoles{}, err
	}

	return response.(*UserRoles), nil
}

// GetAll returns all user roles. If none can be found or an error occurs, it
// returns an empty collection.
func (s *UserRoleService) GetAll() ([]*UserRole, error) {
	items := []*UserRole{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the user role that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *UserRoleService) GetByID(id string) (*UserRole, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(UserRole), path)
	if err != nil {
		return nil, internal.CreateResourceNotFoundError("userRole", "ID", id)
	}

	return resp.(*UserRole), nil
}

// Update modifies a user role based on the one provided as input.
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
