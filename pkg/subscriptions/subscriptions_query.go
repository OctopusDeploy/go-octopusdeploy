package subscriptions

type SubscriptionsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces      []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
