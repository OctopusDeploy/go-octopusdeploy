package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// OctopusProjectFeed represents an Octopus project feed.
type OctopusProjectFeed struct {
	Feed
}

// NewOctopusProjectFeed creates and initializes a Octopus project feed.
func NewOctopusProjectFeed(name string, feedURI string) *OctopusProjectFeed {
	return &OctopusProjectFeed{
		Feed: *newFeed(name, FeedTypeOctopusProject),
	}
}

// Validate checks the state of this Octopus project feed and returns an error
// if invalid.
func (o *OctopusProjectFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(o)
}
