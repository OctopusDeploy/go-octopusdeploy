package runbooks

type RunbookRunsQuery struct {
	Environments []string `uri:"environments,omitempty" url:"environments,omitempty"`
	IDs          []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName  string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Projects     []string `uri:"projects,omitempty" url:"projects,omitempty"`
	Runbooks     []string `uri:"runbooks,omitempty" url:"runbooks,omitempty"`
	Skip         int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take         int      `uri:"take,omitempty" url:"take,omitempty"`
	TaskState    string   `uri:"taskState,omitempty" url:"taskState,omitempty"`
	Tenants      []string `uri:"tenants,omitempty" url:"tenants,omitempty"`
}
