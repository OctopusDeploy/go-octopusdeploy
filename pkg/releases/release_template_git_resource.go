package releases

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actions"

type ReleaseTemplateGitResource struct {
	ActionName                     string                          `json:"ActionName,omitempty"`
	RepositoryUri                  string                          `json:"RepositoryUri,omitempty"`
	DefaultBranch                  string                          `json:"DefaultBranch,omitempty"`
	IsResolvable                   bool                            `json:"IsResolvable"`
	Name                           string                          `json:"Name,omitempty"`
	FilePathFilters                []string                        `json:"FilePathFilters,omitempty"`
	GitCredentialId                string                          `json:"NuGetPackageId,omitempty"`
	GitResourceSelectedLastRelease actions.VersionControlReference `json:"GitResourceSelectedLastRelease,omitempty"`
}
