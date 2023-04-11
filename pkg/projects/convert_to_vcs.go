package projects

type ConvertToVcs struct {
	CommitMessage           string                 `json:"CommitMessage"`
	VersionControlSettings  GitPersistenceSettings `json:"VersionControlSettings"`
	InitialCommitBranchName string                 `json:"InitialCommitBranchName,omitempty"`
}

// NewConvertToVcs returns the new structure to send to Octopus to convert a project to VCS.
// Will return error if initialCommitBranchName not explicitly specified and
// the default branch is listed in the protected branch patterns.
func NewConvertToVcs(commitMessage string, initialCommitBranchName string, gitPersistenceSettings GitPersistenceSettings) *ConvertToVcs {
	c := &ConvertToVcs{
		CommitMessage:           commitMessage,
		VersionControlSettings:  gitPersistenceSettings,
		InitialCommitBranchName: initialCommitBranchName,
	}

	return c
}
