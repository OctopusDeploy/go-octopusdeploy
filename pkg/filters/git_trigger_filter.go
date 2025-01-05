package filters

type GitTriggerSource struct {
	DeploymentActionSlug string   `json:"DeploymentActionSlug"`
	GitDependencyName    string   `json:"GitDependencyName"`
	IncludeFilePaths     []string `json:"IncludeFilePaths"`
	ExcludeFilePaths     []string `json:"ExcludeFilePaths"`
}

type GitTriggerFilter struct {
	Sources []GitTriggerSource `json:"Sources"`

	triggerFilter
}

func NewGitTriggerFilter(gitTriggerSources []GitTriggerSource) *GitTriggerFilter {
	return &GitTriggerFilter{
		Sources:       gitTriggerSources,
		triggerFilter: *newTriggerFilter(GitFilter),
	}
}

func (t *GitTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *GitTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &GitTriggerFilter{}
