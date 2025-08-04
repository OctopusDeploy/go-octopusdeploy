package gitdependencies

type SelectedGitResources struct {
	ActionName               string        `json:"ActionName,omitempty"`
	GitReference             *GitReference `json:"GitReferenceResource,omitempty"`
	GitResourceReferenceName string        `json:"GitResourceReferenceName,omitempty"`
}

type GitReference struct {
	GitRef    string `json:"GitRef,omitempty"`
	GitCommit string `json:"GitCommit,omitempty"`
}
