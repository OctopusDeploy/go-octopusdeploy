package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// MavenFeed represents a Maven feed.
type MavenFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	Feed
}

// NewMavenFeed creates and initializes a Maven feed.
func NewMavenFeed(name string, feedURI string) *MavenFeed {
	return &MavenFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedURI:                     feedURI,
		Feed:                        *newFeed(name, FeedTypeMaven),
	}
}

// GetFeedType returns the feed type of this Maven feed.
func (m *MavenFeed) GetFeedType() FeedType {
	return m.FeedType
}

// Validate checks the state of this Maven feed and returns an error if invalid.
func (m *MavenFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(m)
}

var _ IFeed = &MavenFeed{}
