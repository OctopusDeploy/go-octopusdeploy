package teams

type TeamMembershipQuery struct {
	IncludeSystem bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	Spaces        []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	UserID        string   `uri:"userId,omitempty" url:"userId,omitempty"`
}
