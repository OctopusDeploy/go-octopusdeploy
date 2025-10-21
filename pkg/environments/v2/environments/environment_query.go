package environments

type EnvironmentQuery struct {
	Ids         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip"`
	Take        int      `uri:"take" url:"take"`
	Type        []string `uri:"type,omitempty" url:"type,omitempty"`
}
