package projects

type ConvertToVcs struct {
	CommitMessage           string                 `json:"CommitMessage"`
	VersionControlSettings  GitPersistenceSettings `json:"VersionControlSettings"`
	InitialCommitBranchName string                 `json:"InitialCommitBranchName"`
}

func NewConvertToVcs(commitMessage string, initialCommitBranchName string, gitPersistenceSettings GitPersistenceSettings) *ConvertToVcs {
	return &ConvertToVcs{
		CommitMessage:           commitMessage,
		VersionControlSettings:  gitPersistenceSettings,
		InitialCommitBranchName: initialCommitBranchName,
	}
}
