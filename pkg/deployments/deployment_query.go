package deployments

type DeploymentQuery struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}
