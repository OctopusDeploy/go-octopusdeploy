package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type IFeed interface {
	GetFeedType() FeedType
	GetName() string
	GetPackageAcquisitionLocationOptions() []string
	GetPassword() *core.SensitiveValue
	GetSpaceID() string
	GetUsername() string
	SetFeedType(FeedType)
	SetName(string)
	SetPackageAcquisitionLocationOptions([]string)
	SetPassword(*core.SensitiveValue)
	SetSpaceID(string)
	SetUsername(string)

	resources.IResource
}

// feed is the embedded struct used for all feeds.
type feed struct {
	FeedType                          FeedType             `json:"FeedType" validate:"required,notblank"`
	Name                              string               `json:"Name" validate:"required,notblank"`
	SpaceID                           string               `json:"SpaceId,omitempty"`
	PackageAcquisitionLocationOptions []string             `json:"PackageAcquisitionLocationOptions,omitempty"`
	Password                          *core.SensitiveValue `json:"Password,omitempty"`
	Username                          string               `json:"Username,omitempty"`

	resources.Resource
}

type Feeds struct {
	Items []IFeed `json:"Items"`
	resources.PagedResults
}

// newFeed creates and initializes a feed resource.
func newFeed(name string, feedType FeedType) *feed {
	return &feed{
		Name:                              name,
		FeedType:                          feedType,
		PackageAcquisitionLocationOptions: []string{},
		Resource:                          *resources.NewResource(),
	}
}

// GetFeedType returns the type of this feed.
func (f *feed) GetFeedType() FeedType {
	return f.FeedType
}

// GetName returns the name of the feed.
func (f *feed) GetName() string {
	return f.Name
}

// GetSpaceID returns the space ID of the feed.
func (f *feed) GetSpaceID() string {
	return f.SpaceID
}

// GetPackageAcquisitionLocationOptions returns the package acquisition location options of the feed.
func (f *feed) GetPackageAcquisitionLocationOptions() []string {
	return f.PackageAcquisitionLocationOptions
}

// GetPassword returns the password of the feed.
func (f *feed) GetPassword() *core.SensitiveValue {
	return f.Password
}

// GetUsername returns the username of the feed.
func (f *feed) GetUsername() string {
	return f.Username
}

// SetFeedType returns the type of this feed.
func (f *feed) SetFeedType(feedType FeedType) {
	f.FeedType = feedType
}

// SetName sets the name of the feed.
func (f *feed) SetName(name string) {
	f.Name = name
}

// SetPackageAcquisitionLocationOptions sets the package acquisition location options of the feed.
func (f *feed) SetPackageAcquisitionLocationOptions(packageAcquisitionLocationOptions []string) {
	f.PackageAcquisitionLocationOptions = packageAcquisitionLocationOptions
}

// SetPassword sets the password of the feed.
func (f *feed) SetPassword(password *core.SensitiveValue) {
	f.Password = password
}

// SetSpaceID sets the space ID of the feed.
func (f *feed) SetSpaceID(spaceID string) {
	f.SpaceID = spaceID
}

// SetUsername sets the username of the feed.
func (f *feed) SetUsername(username string) {
	f.Username = username
}

// Validate checks the state of the feed and returns an error if invalid.
func (f feed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(f)
}

var _ IFeed = &feed{}
