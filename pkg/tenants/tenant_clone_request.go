package tenants

type TenantCloneRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type TenantCloneQuery struct {
	CloneTenantID string `uri:"clone" url:"clone"`
}
