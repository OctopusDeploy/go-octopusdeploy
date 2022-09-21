package issuetrackers

type Commit struct {
	Comment string `json:"Comment,omitempty"`
	ID      string `json:"Id,omitempty"`
}

func NewCommit() *Commit {
	return &Commit{}
}
