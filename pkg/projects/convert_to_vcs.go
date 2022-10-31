package projects

type ConvertToVcs struct {
	CommitMessage          string
	GitPersistenceSettings GitPersistenceSettings
}

func NewConvertToVcs(commitMessage string, gitPersistenceSettings GitPersistenceSettings) *ConvertToVcs {
	return &ConvertToVcs{
		CommitMessage:          commitMessage,
		GitPersistenceSettings: gitPersistenceSettings,
	}
}
