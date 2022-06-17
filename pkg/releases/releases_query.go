package releases

type ReleasesQuery struct {
	IDs                []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IgnoreChannelRules bool     `uri:"ignoreChannelRules,omitempty" url:"ignoreChannelRules,omitempty"`
	Skip               int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take               int      `uri:"take,omitempty" url:"take,omitempty"`
}
