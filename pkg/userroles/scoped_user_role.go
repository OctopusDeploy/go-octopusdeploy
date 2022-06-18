package userroles

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

type ScopedUserRole struct {
	EnvironmentIDs  []string `json:"EnvironmentIds,omitempty"`
	ProjectIDs      []string `json:"ProjectIds,omitempty"`
	ProjectGroupIDs []string `json:"ProjectGroupIds,omitempty"`
	TeamID          string   `json:"TeamId"`
	TenantIDs       []string `json:"TenantIds,omitempty"`
	SpaceID         string   `json:"SpaceId"`
	UserRoleID      string   `json:"UserRoleId"`

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
