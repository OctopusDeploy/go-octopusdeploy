package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GitHubRepositoryFeed represents a GitHub repository feed.
type GitHubRepositoryFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	feed
}

// NewGitHubRepositoryFeed creates and initializes a GitHub repository feed.
func NewGitHubRepositoryFeed(name string) (*GitHubRepositoryFeed, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	feed := GitHubRepositoryFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedURI:                     "https://api.github.com",
		feed:                        *newFeed(name, FeedTypeGitHub),
	}

	// validate to ensure that all expectations are met
	err := feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
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
