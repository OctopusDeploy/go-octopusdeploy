package tagsets

type TagSetsQuery struct {
	IDs         []string      `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string        `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Scopes      []TagSetScope `uri:"scopes,omitempty" url:"scopes,omitempty"`
	Skip        int           `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int           `uri:"take,omitempty" url:"take,omitempty"`
}
