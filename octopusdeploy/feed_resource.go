package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// FeedResource is the embedded struct used for all feeds.
type FeedResource struct {
	Name                              string          `json:"Name" validate:"required,notblank"`
	SpaceID                           string          `json:"SpaceId,omitempty"`
	PackageAcquisitionLocationOptions []string        `json:"PackageAcquisitionLocationOptions,omitempty"`
	Password                          *SensitiveValue `json:"Password,omitempty"`
	Username                          string          `json:"Username,omitempty"`

	Resource
}

// newFeedResource creates and initializes a feed resource.
func newFeedResource(name string) *FeedResource {
	return &FeedResource{
		Name:                              name,
		PackageAcquisitionLocationOptions: []string{},
		Resource:                          *newResource(),
	}
}

// GetName returns the name of the feed resource.
func (f *FeedResource) GetName() string {
	return f.Name
}

// SetName sets the name of the account resource.
func (f *FeedResource) SetName(name string) {
	f.Name = name
}

// Validate checks the state of the feed and returns an error if invalid.
func (f FeedResource) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(f)
}

var _ IHasName = &FeedResource{}
