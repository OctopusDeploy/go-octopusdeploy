package gitdependencies

type GitDependency struct {
	Name                         string   `json:"Name" validate:"required,notblank"`
	RepositoryUri                string   `json:"RepositoryUri" validate:"required,notblank"`
	DefaultBranch                string   `json:"DefaultBranch" validate:"required,notblank"`
	GitCredentialType            string   `json:"GitCredentialType" validate:"required,notblank"`
	FilePathFilters              []string `json:"FilePathFilters,omitempty"`
	GitCredentialId              string   `json:"GitCredentialId,omitempty"`
	StepPackageInputsReferenceId string   `json:"StepPackageInputsReferenceId,omitempty"`
	GitHubConnectionId           string   `json:"GitHubConnectionId,omitempty"`
}
