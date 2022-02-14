package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// HelmFeed represents a Helm feed.
type HelmFeed struct {
	FeedURI string `json:"FeedUri,omitempty"`

	Feed
}

// NewHelmFeed creates and initializes a Helm feed.
func NewHelmFeed(name string) *HelmFeed {
	helmFeed := HelmFeed{
		Feed: *newFeed(name, FeedTypeHelm),
	}
	helmFeed.FeedURI = "https://kubernetes-charts.storage.googleapis.com"

	return &helmFeed
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
