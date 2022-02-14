package resources

import (
	"github.com/go-playground/validator/v10"
)

type FeedResource struct {
	AccessKey                         string          `json:"AccessKey,omitempty"`
	APIVersion                        string          `json:"ApiVersion,omitempty"`
	DeleteUnreleasedPackagesAfterDays *int            `json:"DeleteUnreleasedPackagesAfterDays,omitempty"`
	DownloadAttempts                  int             `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds       int             `json:"DownloadRetryBackoffSeconds"`
	EnhancedMode                      bool            `json:"EnhancedMode"`
	FeedType                          FeedType        `json:"FeedType,omitempty"`
	FeedURI                           string          `json:"FeedUri,omitempty"`
	IsBuiltInRepoSyncEnabled          bool            `json:"IsBuiltInRepoSyncEnabled,omitempty"`
	Name                              string          `json:"Name"`
	Password                          *SensitiveValue `json:"Password,omitempty"`
	PackageAcquisitionLocationOptions []string        `json:"PackageAcquisitionLocationOptions,omitempty"`
	Region                            string          `json:"Region,omitempty"`
	RegistryPath                      string          `json:"RegistryPath,omitempty"`
	SecretKey                         *SensitiveValue `json:"SecretKey,omitempty"`
	SpaceID                           string          `json:"SpaceId,omitempty"`
	Username                          string          `json:"Username,omitempty"`

	Resource
}

func NewFeedResource(name string, feedType FeedType) *FeedResource {
	return &FeedResource{
		FeedType: feedType,
		Name:     name,
		Resource: *NewResource(),
	}
}

func (f *FeedResource) GetFeedType() FeedType {
	return f.FeedType
}

func (f *FeedResource) GetName() string {
	return f.Name
}

func (f *FeedResource) SetName(name string) {
	f.Name = name
}

// Validate checks the state of the feed Resource and returns an error if
// invalid.
func (f FeedResource) Validate() error {
	return validator.New().Struct(f)
}

var _ IFeed = &FeedResource{}
