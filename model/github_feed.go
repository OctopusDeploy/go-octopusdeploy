package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GitHubRepositoryFeed represents a GitHub repository feed.
type GitHubRepositoryFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	FeedType                    string `json:"FeedType" validate:"required,eq=GitHub"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	FeedResource
}

// NewGitHubRepositoryFeed creates and initializes a GitHub repository feed.
func NewGitHubRepositoryFeed(name string, feedURI string) *GitHubRepositoryFeed {
	return &GitHubRepositoryFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedType:                    feedGitHub,
		FeedURI:                     feedURI,
		FeedResource:                *newFeedResource(name),
	}
}

// GetFeedType returns the feed type of this GitHub repository feed.
func (g *GitHubRepositoryFeed) GetFeedType() string {
	return g.FeedType
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

var _ IFeed = &GitHubRepositoryFeed{}
