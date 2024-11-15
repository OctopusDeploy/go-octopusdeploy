package deploymentfreezes

type DeploymentFreezeQuery struct {
	IncludeComplete bool     `uri:"includeComplete,omitempty"`
	ProjectIds      []string `uri:"projectIds,omitempty"`
	EnvironmentIds  []string `uri:"environmentIds,omitempty"`
	IDs             []string `uri:"ids,omitempty"`
	Skip            int      `uri:"skip,omitempty"`
	Take            int      `uri:"take,omitempty"`
}
