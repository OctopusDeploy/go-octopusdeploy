package releases

type ReleaseQuery struct {
	SearchByVersion string `uri:"searchByVersion" url:"searchByVersion,omitempty"`
	Skip            int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take            int    `uri:"take,omitempty" url:"take,omitempty"`
}
