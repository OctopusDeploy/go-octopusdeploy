package serverstatus

// LogsQuery represents parameters to query the recent server logs.
type LogsQuery struct {
	IncludeDetail bool `uri:"includeDetail,omitempty" url:"includeDetail,omitempty"`
	Skip          int  `uri:"skip,omitempty" url:"skip,omitempty"`
	Take          int  `uri:"take,omitempty" url:"take,omitempty"`
}
