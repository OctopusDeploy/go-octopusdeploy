package variables

type LibraryVariablesQuery struct {
	ContentType string   `uri:"contentType,omitempty" url:"contentType,omitempty"`
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
