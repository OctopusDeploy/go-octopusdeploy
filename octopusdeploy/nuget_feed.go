package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// NuGetFeed represents a NuGet feed.
type NuGetFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	EnhancedMode                bool   `json:"EnhancedMode"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	Feed
}

// NewNuGetFeed creates and initializes a NuGet feed.
func NewNuGetFeed(name string, feedURI string) *NuGetFeed {
	return &NuGetFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		EnhancedMode:                false,
		FeedURI:                     feedURI,
		Feed:                        *newFeed(name, FeedTypeNuGet),
	}
}

// GetFeedType returns the feed type of this NuGet feed.
func (n *NuGetFeed) GetFeedType() FeedType {
	return n.FeedType
}

// Validate checks the state of this NuGet feed and returns an error if invalid.
func (n *NuGetFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(n)
}

var _ IFeed = &NuGetFeed{}
