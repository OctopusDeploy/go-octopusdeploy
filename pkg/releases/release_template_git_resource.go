package releases

type ReleaseTemplateGitResource struct {
	ActionName                     string                      `json:"ActionName,omitempty"`
	RepositoryUri                  string                      `json:"RepositoryUri,omitempty"`
	DefaultBranch                  string                      `json:"DefaultBranch,omitempty"`
	IsResolvable                   bool                        `json:"IsResolvable"`
	Name                           string                      `json:"Name,omitempty"`
	FilePathFilters                []string                    `json:"FilePathFilters,omitempty"`
	GitCredentialId                string                      `json:"NuGetPackageId,omitempty"`
	GitResourceSelectedLastRelease ReleaseTemplateGitReference `json:"GitResourceSelectedLastRelease,omitempty"`
}

type ReleaseTemplateGitReference struct {
	GitRef    string `json:"GitRef,omitempty"`
	GitCommit string `json:"GitCommit,omitempty"`
}
