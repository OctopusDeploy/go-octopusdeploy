package userroles

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type ScopedUserRoleService struct {
	services.CanDeleteService
}

func NewScopedUserRoleService(sling *sling.Sling, uriTemplate string) *ScopedUserRoleService {
	return &ScopedUserRoleService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceScopedUserRoleService, sling, uriTemplate),
		},
	}
}

func (s *ScopedUserRoleService) Add(scopedUserRole *ScopedUserRole) (*ScopedUserRole, error) {
	if IsNil(scopedUserRole) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterScopedUserRole)
	}

	if err := scopedUserRole.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, scopedUserRole)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), scopedUserRole, new(ScopedUserRole), path)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopedUserRole), nil
}

// Currently no known query params, not even take and skip
// Query params could exist, but are undocumented in the swagger
func (s *ScopedUserRoleService) Get() (*ScopedUserRoles, error) {
	path := s.BasePath

	resp, err := services.ApiGet(s.GetClient(), new(ScopedUserRoles), path)
	if err != nil {
		return &ScopedUserRoles{}, err
	}
	return resp.(*ScopedUserRoles), nil
}

func (s *ScopedUserRoleService) GetByID(id string) (*ScopedUserRole, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(ScopedUserRole), path)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopedUserRole), nil
}

func (s *ScopedUserRoleService) Update(scopedUserRole *ScopedUserRole) (*ScopedUserRole, error) {
	if scopedUserRole == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterScopedUserRole)
	}

	path, err := services.GetUpdatePath(s, scopedUserRole)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), scopedUserRole, new(ScopedUserRole), path)
	if err != nil {
		return nil, err
	}
	return resp.(*ScopedUserRole), nil
}
