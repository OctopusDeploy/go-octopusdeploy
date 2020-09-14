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
	Name                        string         `json:"Name"`
	FeedType                    enum.FeedType  `json:"FeedType"`
	DownloadAttempts            int            `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int            `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string         `json:"FeedUri"`
	EnhancedMode                bool           `json:"EnhancedMode"`
	Username                    string         `json:"Username"`
	Password                    SensitiveValue `json:"Password"`

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

func NewFeed(name, feedType string, feedURI string) *Feed {
	return &Feed{
		Name:     name,
		FeedType: feedType,
		FeedURI:  feedURI,
	}
}

var _ ResourceInterface = &Feed{}
