package feeds

type SearchPackagesQuery struct {
	Skip int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int    `uri:"take,omitempty" url:"take,omitempty"`
	Term string `uri:"term,omitempty" url:"term,omitempty"`
}
