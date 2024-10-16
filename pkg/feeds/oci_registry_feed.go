package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// OCIRegistryFeed represents a feed of Open Container Initiative Registry.
type OCIRegistryFeed struct {
	FeedURI string `json:"FeedUri,omitempty"`

	feed
}

// NewOCIRegistryFeed creates and initializes a OCIRegistry feed.
func NewOCIRegistryFeed(name string) (*OCIRegistryFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	feed := OCIRegistryFeed{
		FeedURI: "oci://registry-1.docker.io",
		feed:    *newFeed(name, FeedTypeMaven),
	}

	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this feed and returns an error if invalid.
func (m *OCIRegistryFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(m)
}
