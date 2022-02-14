package resources

type ConvertToVcs struct {
	CommitMessage          string
	VersionControlSettings *VersionControlSettings
}

func NewConvertToVcs(commitMessage string, versionControlSettings *VersionControlSettings) *ConvertToVcs {
	return &ConvertToVcs{
		CommitMessage:          commitMessage,
		VersionControlSettings: versionControlSettings,
	}
}
