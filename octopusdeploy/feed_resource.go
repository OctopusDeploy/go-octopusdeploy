package octopusdeploy

import "github.com/go-playground/validator/v10"

type FeedResource struct {
	AccessKey                         string          `json:"AccessKey,omitempty"`
	APIVersion                        string          `json:"ApiVersion,omitempty"`
	DeleteUnreleasedPackagesAfterDays int             `json:"DeleteUnreleasedPackagesAfterDays"`
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

	resource
}

type FeedResources struct {
	Items []*FeedResource `json:"Items"`
	PagedResults
}

func NewFeedResource(name string, feedType FeedType) *FeedResource {
	return &FeedResource{
		FeedType: feedType,
		Name:     name,
		resource: *newResource(),
	}
}

// GetFeedType returns the type of this feed.
func (f *FeedResource) GetFeedType() FeedType {
	return f.FeedType
}

// GetName returns the name of the feed.
func (f *FeedResource) GetName() string {
	return f.Name
}

// GetSpaceID returns the space ID of the feed.
func (f *FeedResource) GetSpaceID() string {
	return f.SpaceID
}

// GetPackageAcquisitionLocationOptions returns the package acquisition location options of the feed.
func (f *FeedResource) GetPackageAcquisitionLocationOptions() []string {
	return f.PackageAcquisitionLocationOptions
}

// GetPassword returns the password of the feed.
func (f *FeedResource) GetPassword() *SensitiveValue {
	return f.Password
}

// GetUsername returns the username of the feed.
func (f *FeedResource) GetUsername() string {
	return f.Username
}

// SetFeedType returns the type of this feed.
func (f *FeedResource) SetFeedType(feedType FeedType) {
	f.FeedType = feedType
}

// SetName sets the name of the feed.
func (f *FeedResource) SetName(name string) {
	f.Name = name
}

// SetPackageAcquisitionLocationOptions sets the package acquisition location options of the feed.
func (f *FeedResource) SetPackageAcquisitionLocationOptions(packageAcquisitionLocationOptions []string) {
	f.PackageAcquisitionLocationOptions = packageAcquisitionLocationOptions
}

// SetPassword sets the password of the feed.
func (f *FeedResource) SetPassword(password *SensitiveValue) {
	f.Password = password
}

// SetSpaceID sets the space ID of the feed.
func (f *FeedResource) SetSpaceID(spaceID string) {
	f.SpaceID = spaceID
}

// SetUsername sets the username of the feed.
func (f *FeedResource) SetUsername(username string) {
	f.Username = username
}

// Validate checks the state of the feed resource and returns an error if
// invalid.
func (f FeedResource) Validate() error {
	return validator.New().Struct(f)
}

var _ IFeed = &FeedResource{}
