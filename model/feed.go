package model

import "github.com/go-playground/validator/v10"

type Feeds struct {
	Items []Feed `json:"Items"`
	PagedResults
}

type Feed struct {
	ID                          string         `json:"Id"`
	Name                        string         `json:"Name"`
	FeedType                    string         `json:"FeedType"`
	DownloadAttempts            int            `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int            `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string         `json:"FeedUri"`
	EnhancedMode                bool           `json:"EnhancedMode"`
	Username                    string         `json:"Username"`
	Password                    SensitiveValue `json:"Password"`
}

func (t *Feed) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

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
