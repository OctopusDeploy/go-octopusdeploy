package octopusdeploy

type UserPermissionRestriction struct {
	RestrictedToEnvironmentIds  []string `json:"RestrictedToEnvironmentIds"`
	RestrictedToProjectGroupIds []string `json:"RestrictedToProjectGroupIds"`
	RestrictedToProjectIds      []string `json:"RestrictedToProjectIds"`
	RestrictedToTenantIds       []string `json:"RestrictedToTenantIds"`
	SpaceID                     string   `json:"SpaceId,omitempty"`
}
