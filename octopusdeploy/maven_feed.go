package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// MavenFeed represents a Maven feed.
type MavenFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	FeedType                    string `json:"FeedType" validate:"required,eq=Maven"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	FeedResource
}

// NewMavenFeed creates and initializes a Maven feed.
func NewMavenFeed(name string, feedURI string) *MavenFeed {
	return &MavenFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedType:                    feedMaven,
		FeedURI:                     feedURI,
		FeedResource:                *newFeedResource(name),
	}
}

// GetFeedType returns the feed type of this Maven feed.
func (m *MavenFeed) GetFeedType() string {
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
