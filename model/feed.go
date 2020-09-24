package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type Feeds struct {
	Items []Feed `json:"Items"`
	PagedResults
}

type Feed struct {
	AccessKey                         string         `json:"AccessKey,omitempty"`
	APIVersion                        string         `json:"ApiVersion,omitempty"`
	DeleteUnreleasedPackagesAfterDays int            `json:"DeleteUnreleasedPackagesAfterDays,omitempty"`
	DownloadAttempts                  int            `json:"DownloadAttempts,omitempty"`
	DownloadRetryBackoffSeconds       int            `json:"DownloadRetryBackoffSeconds,omitempty"`
	EnhancedMode                      bool           `json:"EnhancedMode,omitempty"`
	FeedType                          enum.FeedType  `json:"FeedType,omitempty"`
	FeedURI                           string         `json:"FeedUri,omitempty"`
	IsBuiltInRepoSyncEnabled          bool           `json:"IsBuiltInRepoSyncEnabled,omitempty"`
	Name                              string         `json:"Name,omitempty"`
	Password                          SensitiveValue `json:"Password,omitempty"`
	PackageAcquisitionLocationOptions []string       `json:"PackageAcquisitionLocationOptions,omitempty"`
	Region                            string         `json:"Region,omitempty"`
	RegistryPath                      string         `json:"RegistryPath,omitempty"`
	SecretKey                         SensitiveValue `json:"SecretKey,omitempty"`
	SpaceID                           string         `json:"SpaceId,omitempty"`
	Username                          string         `json:"Username,omitempty"`

	Resource
}

// GetID returns the ID value of the Feed.
func (resource Feed) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Feed.
func (resource Feed) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Feed was changed.
func (resource Feed) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Feed.
func (resource Feed) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Feed and returns an error if invalid.
func (resource Feed) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

func NewFeed(name string, feedType enum.FeedType, feedURI string) *Feed {
	return &Feed{
		Name:     name,
		FeedType: feedType,
		FeedURI:  feedURI,
	}
}

var _ ResourceInterface = &Feed{}
