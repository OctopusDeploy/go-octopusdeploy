package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// Feed is the embedded struct used for all feeds.
type Feed struct {
	Name                              string          `json:"Name" validate:"required,notblank"`
	FeedType                          FeedType        `json:"FeedType"`
	SpaceID                           string          `json:"SpaceId,omitempty"`
	PackageAcquisitionLocationOptions []string        `json:"PackageAcquisitionLocationOptions,omitempty"`
	Password                          *SensitiveValue `json:"Password,omitempty"`
	Username                          string          `json:"Username,omitempty"`

	Resource
}

type Feeds struct {
	Items []IFeed `json:"Items"`
	PagedResults
}

// newFeed creates and initializes a feed Resource.
func newFeed(name string, feedType FeedType) *Feed {
	return &Feed{
		FeedType:                          feedType,
		Name:                              name,
		PackageAcquisitionLocationOptions: []string{},
		Resource:                          *newResource(),
	}
}

// GetFeedType returns the type of this feed.
func (f *Feed) GetFeedType() FeedType {
	return f.FeedType
}

// GetName returns the name of the feed.
func (f *Feed) GetName() string {
	return f.Name
}

// SetName sets the name of the feed.
func (f *Feed) SetName(name string) {
	f.Name = name
}

// Validate checks the state of the feed and returns an error if invalid.
func (f Feed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(f)
}

var _ IHasName = &Feed{}
