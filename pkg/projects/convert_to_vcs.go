package projects

type ConvertToVcs struct {
	CommitMessage          string
	VersionControlSettings GitPersistenceSettings
}

func NewConvertToVcs(commitMessage string, gitPersistenceSettings GitPersistenceSettings) *ConvertToVcs {
	return &ConvertToVcs{
		CommitMessage:          commitMessage,
		VersionControlSettings: gitPersistenceSettings,
	}
}
