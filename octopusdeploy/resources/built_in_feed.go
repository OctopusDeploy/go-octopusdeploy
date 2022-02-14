package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// BuiltInFeed represents a built-in feed.
type BuiltInFeed struct {
	DeleteUnreleasedPackagesAfterDays *int `json:"DeleteUnreleasedPackagesAfterDays,omitempty"`
	DownloadAttempts                  int  `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds       int  `json:"DownloadRetryBackoffSeconds"`
	IsBuiltInRepoSyncEnabled          bool `json:"IsBuiltInRepoSyncEnabled"`

	Feed
}

// NewBuiltInFeed creates and initializes a built-in feed.
func NewBuiltInFeed(name string, feedURI string) *BuiltInFeed {
	return &BuiltInFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		IsBuiltInRepoSyncEnabled:    false,
		Feed:                        *newFeed(name, FeedTypeBuiltIn),
	}
}

// Validate checks the state of this built-in feed and returns an error if
// invalid.
func (b *BuiltInFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(b)
}
