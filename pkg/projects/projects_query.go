package projects

type ProjectsQuery struct {
	ClonedFromProjectID string   `uri:"clonedFromProjectId,omitempty" url:"clonedFromProjectId,omitempty"`
	IDs                 []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsClone             bool     `uri:"clone,omitempty" url:"clone,omitempty"`
	Name                string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName         string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip                int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take                int      `uri:"take,omitempty" url:"take,omitempty"`
}
