package model

import (
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
	Region                            string         `json:"Region,omitempty"`
	RegistryPath                      string         `json:"RegistryPath,omitempty"`
	SecretKey                         SensitiveValue `json:"SecretKey,omitempty"`
	Username                          string         `json:"Username,omitempty"`

	Resource
}

func (f *Feed) GetID() string {
	return f.ID
}

func (f *Feed) Validate() error {
	validate := validator.New()
	err := validate.Struct(f)

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
