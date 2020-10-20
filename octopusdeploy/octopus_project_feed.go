package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// OctopusProjectFeed represents an Octopus project feed.
type OctopusProjectFeed struct {
	FeedType string `json:"FeedType" validate:"required,eq=OctopusProject"`

	FeedResource
}

// NewOctopusProjectFeed creates and initializes a Octopus project feed.
func NewOctopusProjectFeed(name string, feedURI string) *OctopusProjectFeed {
	return &OctopusProjectFeed{
		FeedType:     feedOctopusProject,
		FeedResource: *newFeedResource(name),
	}
}

// GetFeedType returns the feed type of this Octopus project feed.
func (o *OctopusProjectFeed) GetFeedType() string {
	return o.FeedType
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

var _ IFeed = &OctopusProjectFeed{}
