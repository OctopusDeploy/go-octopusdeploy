package interruptions

type InterruptionsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PendingOnly bool     `uri:"pendingOnly,omitempty" url:"pendingOnly,omitempty"`
	Regarding   string   `uri:"regarding,omitempty" url:"regarding,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
