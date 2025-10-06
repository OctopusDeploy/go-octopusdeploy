package environments

type EnvironmentV2Query struct {
	Skip int    `uri:"skip" url:"skip"`
	Take int    `uri:"take" url:"take"`
	Type string `uri:"type,omitempty" url:"type,omitempty"`
}
