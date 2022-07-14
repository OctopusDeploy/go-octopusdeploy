package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// OctopusProjectFeed represents an Octopus project feed.
type OctopusProjectFeed struct {
	feed
}

// NewOctopusProjectFeed creates and initializes a Octopus project feed.
func NewOctopusProjectFeed(name string) (*OctopusProjectFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	feed := OctopusProjectFeed{
		feed: *newFeed(name, FeedTypeOctopusProject),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
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
