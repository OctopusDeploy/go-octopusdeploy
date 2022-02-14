package resources

type VersionControlReference struct {
	GitRef    string `json:"GitRef"`
	GitCommit string `json:"GitCommit"`
}
