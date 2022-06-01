package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// HelmFeed represents a Helm feed.
type HelmFeed struct {
	FeedURI string `json:"FeedUri,omitempty"`

	feed
}

// NewHelmFeed creates and initializes a Helm feed.
func NewHelmFeed(name string) (*HelmFeed, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	feed := HelmFeed{
		FeedURI: "https://kubernetes-charts.storage.googleapis.com",
		feed:    *newFeed(name, FeedTypeHelm),
	}

	// validate to ensure that all expectations are met
	err := feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
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
