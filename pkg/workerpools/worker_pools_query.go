package workerpools

type WorkerPoolsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Name        string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
