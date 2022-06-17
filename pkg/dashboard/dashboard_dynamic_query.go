package dashboard

type DashboardDynamicQuery struct {
	Environments    []string `uri:"environments,omitempty" url:"environments,omitempty"`
	IncludePrevious bool     `uri:"includePrevious,omitempty" url:"includePrevious,omitempty"`
	Projects        []string `uri:"projects,omitempty" url:"projects,omitempty"`
}
