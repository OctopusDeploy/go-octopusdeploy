package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GitHubRepositoryFeed represents a GitHub repository feed.
type GitHubRepositoryFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	Feed
}

// NewGitHubRepositoryFeed creates and initializes a GitHub repository feed.
func NewGitHubRepositoryFeed(name string) *GitHubRepositoryFeed {
	gitHubRepositoryFeed := GitHubRepositoryFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		Feed:                        *newFeed(name, FeedTypeGitHub),
	}
	gitHubRepositoryFeed.FeedURI = "https://api.github.com"

	return &gitHubRepositoryFeed
}

// Validate checks the state of this GitHub repository feed and returns an
// error if invalid.
func (g *GitHubRepositoryFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(g)
}
