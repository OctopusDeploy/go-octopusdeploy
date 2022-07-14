package artifacts

type Query struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Order       string   `uri:"order" url:"order,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Regarding   string   `uri:"regarding,omitempty" url:"regarding,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}
