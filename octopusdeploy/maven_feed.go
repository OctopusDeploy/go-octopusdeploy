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

	Feed
}

// NewMavenFeed creates and initializes a Maven feed.
func NewMavenFeed(name string) *MavenFeed {
	mavenFeed := MavenFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		Feed:                        *newFeed(name, FeedTypeMaven),
	}
	mavenFeed.FeedURI = "https://repo.maven.apache.org/maven2/"

	return &mavenFeed
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
