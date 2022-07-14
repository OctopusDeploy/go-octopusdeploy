package userroles

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type ScopedUserRole struct {
	EnvironmentIDs  []string `json:"EnvironmentIds,omitempty"`
	ProjectIDs      []string `json:"ProjectIds,omitempty"`
	ProjectGroupIDs []string `json:"ProjectGroupIds,omitempty"`
	TeamID          string   `json:"TeamId" validate:"required"`
	TenantIDs       []string `json:"TenantIds,omitempty"`
	SpaceID         string   `json:"SpaceId"`
	UserRoleID      string   `json:"UserRoleId" validate:"required"`

	resources.Resource
}

type ScopedUserRoles struct {
	Items []*ScopedUserRole `json:"Items"`
	resources.PagedResults
}

func NewScopedUserRole(userRoleId string) *ScopedUserRole {
	return &ScopedUserRole{
		UserRoleID: userRoleId,
		Resource:   *resources.NewResource(),
	}
}

// Validate checks the state of the scoped user role and returns an error if invalid.
func (r *ScopedUserRole) Validate() error {
	return validator.New().Struct(r)
}
