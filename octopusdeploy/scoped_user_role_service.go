package octopusdeploy

import "github.com/dghubble/sling"

type scopedUserRoleService struct {
	canDeleteService
}

func newScopedUserRoleService(sling *sling.Sling, uriTemplate string) *scopedUserRoleService {
	scopedUserRoleService := &scopedUserRoleService{}
	scopedUserRoleService.service = newService(ServiceScopedUserRoleService, sling, uriTemplate)

	return scopedUserRoleService
}

func (s scopedUserRoleService) Add(scopedUserRole *ScopedUserRole) (*ScopedUserRole, error) {
	if scopedUserRole == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterScopedUserRole)
	}

	path, err := getAddPath(s, scopedUserRole)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), scopedUserRole, new(ScopedUserRole), path)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopedUserRole), nil
}

// Currently no known query params, not even take and skip
// Query params could exist, but are undocumented in the swagger
func (s scopedUserRoleService) Get() (*ScopedUserRoles, error) {
	path := s.BasePath

	resp, err := apiGet(s.getClient(), new(ScopedUserRoles), path)
	if err != nil {
		return &ScopedUserRoles{}, err
	}
	return resp.(*ScopedUserRoles), nil
}

func (s scopedUserRoleService) GetByID(id string) (*ScopedUserRole, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(ScopedUserRole), path)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopedUserRole), nil
}

func (s scopedUserRoleService) Update(scopedUserRole *ScopedUserRole) (*ScopedUserRole, error) {
	if scopedUserRole == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterScopedUserRole)
	}

	path, err := getUpdatePath(s, scopedUserRole)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), scopedUserRole, new(ScopedUserRole), path)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopedUserRole), nil
}
