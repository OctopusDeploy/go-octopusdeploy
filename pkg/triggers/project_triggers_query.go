package triggers

type ProjectTriggersQuery struct {
	IDs      []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Runbooks []string `uri:"runbooks,omitempty" url:"runbooks,omitempty"`
	Skip     int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take     int      `uri:"take,omitempty" url:"take,omitempty"`
}
