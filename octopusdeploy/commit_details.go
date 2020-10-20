package octopusdeploy

type CommitDetails struct {
	Comment string `json:"Comment,omitempty"`
	ID      string `json:"Id,omitempty"`
	LinkURL string `json:"LinkUrl,omitempty"`
}

func NewCommitDetails() *CommitDetails {
	return &CommitDetails{}
}
