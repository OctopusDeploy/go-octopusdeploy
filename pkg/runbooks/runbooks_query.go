package runbooks

type RunbooksQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsClone     bool     `uri:"clone,omitempty" url:"clone,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ProjectIDs  []string `uri:"projectIds,omitempty" url:"projectIds,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
