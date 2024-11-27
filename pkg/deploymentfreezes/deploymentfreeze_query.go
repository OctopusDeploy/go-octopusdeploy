package deploymentfreezes

type DeploymentFreezeQuery struct {
	IncludeComplete bool     `uri:"includeComplete,omitempty" url:"includeComplete,omitempty"`
	ProjectIds      []string `uri:"projectIds,omitempty" url:"projectIds,omitempty"`
	EnvironmentIds  []string `uri:"environmentIds,omitempty" url:"environmentIds,omitempty"`
	IDs             []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip            int      `uri:"skip" url:"skip"`
	Take            int      `uri:"take,omitempty" url:"take,omitempty"`
}
