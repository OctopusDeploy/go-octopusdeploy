package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// NpmFeed represents an NPM feed.
type NpmFeed struct {
	DownloadAttempts            int                `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int                `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string             `json:"FeedUri,omitempty"`
	Username                    string             `json:"Username,omitempty"`
	Password                    *core.SensitiveValue `json:"Password,omitempty"`

	feed
}

// NewNpmFeed creates and initializes an NPM feed.
func NewNpmFeed(name string, feedURI string) (*NpmFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	feed := NpmFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedURI:                     feedURI,
		feed:                        *newFeed(name, FeedTypeNpm),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this NPM feed and returns an error if invalid.
func (n *NpmFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(n)
}
