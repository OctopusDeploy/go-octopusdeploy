package resources

type UserPermissionSet struct {
	ID                    string                           `json:"Id"`
	IsPermissionsComplete bool                             `json:"IsPermissionsComplete,omitempty"`
	IsTeamsComplete       bool                             `json:"IsTeamsComplete,omitempty"`
	Links             map[string]string                `json:"Links,omitempty"`
	SpacePermissions  SpacePermissions                 `json:"SpacePermissions,omitempty"`
	SystemPermissions []string                         `json:"SystemPermissions"`
	Teams             []ProjectedTeamReferenceDataItem `json:"Teams"`
}
