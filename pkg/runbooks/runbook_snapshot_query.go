package runbooks

type RunbookSnapshotsQuery struct {
	IDs     []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Publish bool     `uri:"publish,omitempty" url:"publish,omitempty"`
	Skip    int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take    int      `uri:"take,omitempty" url:"take,omitempty"`
}
