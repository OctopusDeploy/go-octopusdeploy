package dashboard

type DashboardQuery struct {
	IncludeLatest   bool     `url:"highestLatestVersionPerProjectAndEnvironment"`
	ProjectID       string   `uri:"projectId,omitempty" url:"projectId,omitempty"`
	SelectedTags    []string `uri:"selectedTags,omitempty" url:"selectedTags,omitempty"`
	SelectedTenants []string `uri:"selectedTenants,omitempty" url:"selectedTenants,omitempty"`
	ShowAll         bool     `uri:"showAll,omitempty" url:"showAll,omitempty"`
	ReleaseID       string   `uri:"releaseId,omitempty" url:"releaseId,omitempty"`
}
