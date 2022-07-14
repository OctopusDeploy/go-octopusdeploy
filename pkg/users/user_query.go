package users

type UserQuery struct {
	IncludeSystem bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	Spaces        []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
}
