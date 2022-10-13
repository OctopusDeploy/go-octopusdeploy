package credentials

type Query struct {
	Name string `uri:"name,omitempty" url:"name,omitempty"`
	Skip int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int    `uri:"take,omitempty" url:"take,omitempty"`
}
