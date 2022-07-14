package events

type EventsQuery struct {
	AsCSV             string   `uri:"asCsv,omitempty" url:"asCsv,omitempty"`
	DocumentTypes     []string `uri:"documentTypes,omitempty" url:"documentTypes,omitempty"`
	Environments      []string `uri:"environments,omitempty" url:"environments,omitempty"`
	EventAgents       []string `uri:"eventAgents,omitempty" url:"eventAgents,omitempty"`
	EventCategories   []string `uri:"eventCategories,omitempty" url:"eventCategories,omitempty"`
	EventGroups       []string `uri:"eventGroups,omitempty" url:"eventGroups,omitempty"`
	ExcludeDifference bool     `uri:"excludeDifference,omitempty" url:"excludeDifference,omitempty"`
	From              string   `uri:"from,omitempty" url:"from,omitempty"`
	FromAutoID        string   `uri:"fromAutoId,omitempty" url:"fromAutoId,omitempty"`
	IDs               []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IncludeSystem     bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	Internal          string   `uri:"interal,omitempty" url:"interal,omitempty"`
	Name              string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName       string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ProjectGroups     []string `uri:"projectGroups,omitempty" url:"projectGroups,omitempty"`
	Projects          []string `uri:"projects,omitempty" url:"projects,omitempty"`
	Regarding         string   `uri:"regarding,omitempty" url:"regarding,omitempty"`
	RegardingAny      string   `uri:"regardingAny,omitempty" url:"regardingAny,omitempty"`
	Skip              int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces            []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	Take              int      `uri:"take,omitempty" url:"take,omitempty"`
	Tags              []string `uri:"tags,omitempty" url:"tags,omitempty"`
	Tenants           []string `uri:"tenants,omitempty" url:"tenants,omitempty"`
	To                string   `uri:"to,omitempty" url:"to,omitempty"`
	ToAutoID          string   `uri:"toAutoId,omitempty" url:"toAutoId,omitempty"`
	User              string   `uri:"user,omitempty" url:"user,omitempty"`
	Users             []string `uri:"users,omitempty" url:"users,omitempty"`
}
