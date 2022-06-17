package users

type UsersQuery struct {
	Filter string   `uri:"filter,omitempty" url:"filter,omitempty"`
	IDs    []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip   int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take   int      `uri:"take,omitempty" url:"take,omitempty"`
}
