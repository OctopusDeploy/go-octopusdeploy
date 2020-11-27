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
