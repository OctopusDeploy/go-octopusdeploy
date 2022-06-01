package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// OctopusProjectFeed represents an Octopus project feed.
type OctopusProjectFeed struct {
	feed
}

// NewOctopusProjectFeed creates and initializes a Octopus project feed.
func NewOctopusProjectFeed(name string) (*OctopusProjectFeed, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	feed := OctopusProjectFeed{
		feed: *newFeed(name, FeedTypeOctopusProject),
	}

	// validate to ensure that all expectations are met
	err := feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
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
