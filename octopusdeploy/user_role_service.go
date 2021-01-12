package octopusdeploy

import "github.com/dghubble/sling"

type userRoleService struct {
	canDeleteService
}

func newUserRoleService(sling *sling.Sling, uriTemplate string) *userRoleService {
	userRoleService := &userRoleService{}
	userRoleService.service = newService(ServiceUserRoleService, sling, uriTemplate)

	return userRoleService
}

// Add creates a new user role.
func (s userRoleService) Add(userRole *UserRole) (*UserRole, error) {
	if userRole == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterUserRole)
	}

	path, err := getAddPath(s, userRole)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), userRole, new(UserRole), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserRole), nil
}

// Get returns a collection of user roles based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s userRoleService) Get(userRolesQuery UserRolesQuery) (*UserRoles, error) {
	path, err := s.getURITemplate().Expand(userRolesQuery)
	if err != nil {
		return &UserRoles{}, err
	}

	response, err := apiGet(s.getClient(), new(UserRoles), path)
	if err != nil {
		return &UserRoles{}, err
	}

	return response.(*UserRoles), nil
}

// GetAll returns all user roles. If none can be found or an error occurs, it
// returns an empty collection.
func (s userRoleService) GetAll() ([]*UserRole, error) {
	items := []*UserRole{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the user role that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s userRoleService) GetByID(id string) (*UserRole, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(UserRole), path)
	if err != nil {
		return nil, createResourceNotFoundError("userRole", "ID", id)
	}

	return resp.(*UserRole), nil
}

// Update modifies a user role based on the one provided as input.
func (s userRoleService) Update(userRole *UserRole) (*UserRole, error) {
	if userRole == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUserRole)
	}

	path, err := getUpdatePath(s, userRole)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), userRole, new(UserRole), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserRole), nil
}
