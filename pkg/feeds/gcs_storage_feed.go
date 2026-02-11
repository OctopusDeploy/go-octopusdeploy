package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type GcsStorageFeed struct {
	UseServiceAccountKey        bool                                       `json:"UseServiceAccountKey"`
	ServiceAccountJsonKey       *core.SensitiveValue                       `json:"ServiceAccountJsonKey,omitempty"`
	Project                     string                                     `json:"Project,omitempty"`
	OidcAuthentication          *GoogleContainerRegistryOidcAuthentication `json:"OidcAuthentication,omitempty"`
	DownloadAttempts            int                                        `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int                                        `json:"DownloadRetryBackoffSeconds"`

	feed
}

func NewGcsStorageFeed(name string, useServiceAccountKey bool, serviceAccountJsonKey *core.SensitiveValue, project string, oidcAuthentication *GoogleContainerRegistryOidcAuthentication) (*GcsStorageFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	if useServiceAccountKey {
		if serviceAccountJsonKey == nil || !serviceAccountJsonKey.HasValue {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("serviceAccountJsonKey")
		}
	} else {
		if oidcAuthentication == nil {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("oidcAuthentication")
		}
	}

	feed := GcsStorageFeed{
		UseServiceAccountKey:        useServiceAccountKey,
		ServiceAccountJsonKey:       serviceAccountJsonKey,
		Project:                     project,
		OidcAuthentication:          oidcAuthentication,
		feed:                        *newFeed(name, FeedTypeGcsStorage),
	}

	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

func (g *GcsStorageFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(g)
}
