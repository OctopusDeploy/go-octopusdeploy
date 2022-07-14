package actiontemplates

// Query represents parameters to query the ActionTemplates service.
type Query struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
