package deploymentfreezes

type DeploymentFreezeQuery struct {
	IncludeComplete bool     `uri:"includeComplete,omitempty" url:"includeComplete,omitempty"`
	PartialName     string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Status          string   `uri:"status,omitempty" url:"status,omitempty"`
	ProjectIds      []string `uri:"projectIds,omitempty" url:"projectIds,omitempty"`
	TenantIds       []string `uri:"tenantIds,omitempty" url:"tenantIds,omitempty"`
	EnvironmentIds  []string `uri:"environmentIds,omitempty" url:"environmentIds,omitempty"`
	IDs             []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip            int      `uri:"skip" url:"skip"`
	Take            int      `uri:"take,omitempty" url:"take,omitempty"`
}
