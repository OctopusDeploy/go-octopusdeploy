package githubconnections

// Repository represents a GitHub repository reachable through a GitHub App connection.
type Repository struct {
	RepositoryID   string `json:"RepositoryId"`
	RepositoryName string `json:"RepositoryName"`
	IsAdmin        bool   `json:"IsAdmin"`
	IsPrivate      bool   `json:"IsPrivate"`
	Visibility     string `json:"Visibility"`
	Language       string `json:"Language,omitempty"`
	GitURL         string `json:"GitUrl"`
	DefaultBranch  string `json:"DefaultBranch"`
}

// UnknownRepository represents a repository configured on a connection that has no
// matching repository returned from GitHub.
type UnknownRepository struct {
	RepositoryID   string `json:"RepositoryId"`
	RepositoryName string `json:"RepositoryName,omitempty"`
}
