package feeds

type FeedsQuery struct {
	FeedType    string   `uri:"feedType,omitempty" url:"feedType,omitempty"`
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
