package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// HelmFeed represents a Helm feed.
type HelmFeed struct {
	FeedType string `json:"FeedType" validate:"required,eq=Helm"`
	FeedURI  string `json:"FeedUri,omitempty"`

	FeedResource
}

// NewHelmFeed creates and initializes a Helm feed.
func NewHelmFeed(name string, feedURI string) *HelmFeed {
	return &HelmFeed{
		FeedType:     feedHelm,
		FeedURI:      feedURI,
		FeedResource: *newFeedResource(name),
	}
}

// GetFeedType returns the feed type of this Helm feed.
func (h *HelmFeed) GetFeedType() string {
	return h.FeedType
}

// Validate checks the state of this Helm feed and returns an error if invalid.
func (h *HelmFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(h)
}

var _ IFeed = &HelmFeed{}
