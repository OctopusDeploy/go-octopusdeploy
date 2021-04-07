package octopusdeploy

type ScopedUserRole struct {
	EnvironmentIDs  []string `json:"EnvironmentIds,omitempty"`
	ProjectIDs      []string `json:"ProjectIds,omitempty"`
	ProjectGroupIDs []string `json:"ProjectGroupIds,omitempty"`
	TeamID          string   `json:"TeamId"`
	TenantIDs       []string `json:"TenantIds,omitempty"`
	SpaceID         string   `json:"SpaceId"`
	UserRoleID      string   `json:"UserRoleId"`

	resource
}

type ScopedUserRoles struct {
	Items []*ScopedUserRole `json:"Items"`
	PagedResults
}

func NewScopedUserRole(userRoleId string) *ScopedUserRole {
	return &ScopedUserRole{
		UserRoleID: userRoleId,
		resource:   *newResource(),
	}
}
