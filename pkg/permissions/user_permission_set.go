package permissions

type UserPermissionSet struct {
	ID                    string                           `json:"Id"`
	IsPermissionsComplete bool                             `json:"IsPermissionsComplete"`
	IsTeamsComplete       bool                             `json:"IsTeamsComplete"`
	Links                 map[string]string                `json:"Links,omitempty"`
	SpacePermissions      SpacePermissions                 `json:"SpacePermissions,omitempty"`
	SystemPermissions     []string                         `json:"SystemPermissions"`
	Teams                 []ProjectedTeamReferenceDataItem `json:"Teams"`
}
