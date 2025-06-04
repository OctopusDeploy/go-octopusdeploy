package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// BuiltInFeed represents a built-in feed.
type BuiltInFeed struct {
	DeletePackagesAssociatedWithReleases bool `json:"DeletePackagesAssociatedWithReleases"`
	DeleteUnreleasedPackagesAfterDays    int  `json:"DeleteUnreleasedPackagesAfterDays"`
	DownloadAttempts                     int  `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds          int  `json:"DownloadRetryBackoffSeconds"`
	IsBuiltInRepoSyncEnabled             bool `json:"IsBuiltInRepoSyncEnabled"`

	feed
}

// NewBuiltInFeed creates and initializes a built-in feed.
func NewBuiltInFeed(name string) (*BuiltInFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	feed := BuiltInFeed{
		DeletePackagesAssociatedWithReleases: false,
		DeleteUnreleasedPackagesAfterDays:    30,
		DownloadAttempts:                     5,
		DownloadRetryBackoffSeconds:          10,
		IsBuiltInRepoSyncEnabled:             false,
		feed:                                 *newFeed(name, FeedTypeBuiltIn),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this built-in feed and returns an error if
// invalid.
func (b *BuiltInFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(b)
}
