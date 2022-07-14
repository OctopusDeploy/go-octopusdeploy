package tenants

type TenantsQuery struct {
	ClonedFromTenantID string   `uri:"clonedFromTenantId,omitempty" url:"clonedFromTenantId,omitempty"`
	IDs                []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsClone            bool     `uri:"clone,omitempty" url:"clone,omitempty"`
	Name               string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName        string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ProjectID          string   `uri:"projectId,omitempty" url:"projectId,omitempty"`
	Skip               int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Tags               []string `uri:"tags,omitempty" url:"tags,omitempty"`
	Take               int      `uri:"take,omitempty" url:"take,omitempty"`
}
