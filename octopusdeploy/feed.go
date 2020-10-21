package octopusdeploy

import "github.com/go-playground/validator/v10"

type Feed struct {
	AccessKey                         string          `json:"AccessKey,omitempty"`
	APIVersion                        string          `json:"ApiVersion,omitempty"`
	DeleteUnreleasedPackagesAfterDays *int            `json:"DeleteUnreleasedPackagesAfterDays,omitempty"`
	DownloadAttempts                  int             `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds       int             `json:"DownloadRetryBackoffSeconds"`
	EnhancedMode                      bool            `json:"EnhancedMode"`
	FeedType                          string          `json:"FeedType,omitempty"`
	FeedURI                           *string         `json:"FeedUri,omitempty"`
	IsBuiltInRepoSyncEnabled          bool            `json:"IsBuiltInRepoSyncEnabled,omitempty"`
	Name                              string          `json:"Name"`
	Password                          *SensitiveValue `json:"Password,omitempty"`
	PackageAcquisitionLocationOptions []string        `json:"PackageAcquisitionLocationOptions,omitempty"`
	Region                            string          `json:"Region,omitempty"`
	RegistryPath                      string          `json:"RegistryPath,omitempty"`
	SecretKey                         *SensitiveValue `json:"SecretKey,omitempty"`
	SpaceID                           string          `json:"SpaceId,omitempty"`
	Username                          string          `json:"Username,omitempty"`

	resource
}

type Feeds struct {
	Items []*Feed `json:"Items"`
	PagedResults
}

func NewFeed(name string, feedType string, feedURI string) *Feed {
	return &Feed{
		FeedType: feedType,
		FeedURI:  &feedURI,
		Name:     name,
		resource: *newResource(),
	}
}

func (f *Feed) GetFeedType() string {
	return f.FeedType
}

func (f *Feed) GetName() string {
	return f.Name
}

func (f *Feed) SetName(name string) {
	f.Name = name
}

// Validate checks the state of the feed and returns an error if invalid.
func (f Feed) Validate() error {
	return validator.New().Struct(f)
}

var _ IFeed = &Feed{}
