package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PyPiFeed represents a PyPI feed.
type PyPiFeed struct {
	DownloadAttempts            int                  `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int                  `json:"DownloadRetryBackoffSeconds"`
	FeedURI                     string               `json:"FeedUri,omitempty"`
	Username                    string               `json:"Username,omitempty"`
	Password                    *core.SensitiveValue `json:"Password,omitempty"`

	feed
}

// NewPyPiFeed creates and initializes a PyPI feed.
func NewPyPiFeed(name string, feedURI string) (*PyPiFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	feed := PyPiFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedURI:                     feedURI,
		feed:                        *newFeed(name, FeedTypePyPI),
	}

	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this PyPI feed and returns an error if invalid.
func (p *PyPiFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(p)
}
