package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// NuGetFeed represents a NuGet feed.
type NuGetFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	EnhancedMode                bool   `json:"EnhancedMode"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	feed
}

// NewNuGetFeed creates and initializes a NuGet feed.
func NewNuGetFeed(name string, feedURI string) (*NuGetFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	feed := NuGetFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		EnhancedMode:                false,
		FeedURI:                     feedURI,
		feed:                        *newFeed(name, FeedTypeNuGet),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
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
