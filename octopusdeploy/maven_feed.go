package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// MavenFeed represents a Maven feed.
type MavenFeed struct {
	DownloadAttempts            int    `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int    `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string `json:"FeedUri,omitempty"`

	feed
}

// NewMavenFeed creates and initializes a Maven feed.
func NewMavenFeed(name string) (*MavenFeed, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	feed := MavenFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedURI:                     "https://repo.maven.apache.org/maven2/",
		feed:                        *newFeed(name, FeedTypeMaven),
	}

	// validate to ensure that all expectations are met
	err := feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this Maven feed and returns an error if invalid.
func (m *MavenFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(m)
}
