package channels

type Query struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type QueryByProjectID struct {
	ProjectID   string `uri:"projectId,omitempty" url:"projectId,omitempty"`
	PartialName string `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int    `uri:"take,omitempty" url:"take,omitempty"`
}
