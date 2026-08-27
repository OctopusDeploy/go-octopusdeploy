package dashboard

type DashboardDynamicQuery struct {
	// Environments narrows the dashboard to these environments. The server
	// matches on ID only: an environment name matches nothing and yields an
	// empty dashboard rather than an error.
	Environments    []string `uri:"environments,omitempty" url:"environments,omitempty"`
	IncludePrevious bool     `uri:"includePrevious,omitempty" url:"includePrevious,omitempty"`
	// Projects narrows the dashboard to these projects. As with Environments,
	// the server matches on ID only.
	Projects []string `uri:"projects,omitempty" url:"projects,omitempty"`
}
