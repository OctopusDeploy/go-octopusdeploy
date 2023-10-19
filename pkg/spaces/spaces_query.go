package spaces

type SpacesQuery struct {
	IDs         []string `uri:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty"`
	Take        int      `uri:"take,omitempty"`
}
