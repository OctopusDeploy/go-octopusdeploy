package actions

type CommunityActionTemplatesQuery struct {
	IDs  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int      `uri:"take,omitempty" url:"take,omitempty"`
}
