package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// BuiltInFeed represents a built-in feed.
type BuiltInFeed struct {
	DeleteUnreleasedPackagesAfterDays *int   `json:"DeleteUnreleasedPackagesAfterDays,omitempty"`
	DownloadAttempts                  int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds       int    `json:"DownloadRetryBackoffSeconds"`
	FeedType                          string `json:"FeedType" validate:"required,eq=BuiltIn"`
	IsBuiltInRepoSyncEnabled          bool   `json:"IsBuiltInRepoSyncEnabled"`

	FeedResource
}

// NewBuiltInFeed creates and initializes a built-in feed.
func NewBuiltInFeed(name string, feedURI string) *BuiltInFeed {
	return &BuiltInFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedType:                    feedBuiltIn,
		IsBuiltInRepoSyncEnabled:    false,
		FeedResource:                *newFeedResource(name),
	}
}

// GetFeedType returns the feed type of this built-in feed.
func (b *BuiltInFeed) GetFeedType() string {
	return b.FeedType
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

var _ IFeed = &BuiltInFeed{}
