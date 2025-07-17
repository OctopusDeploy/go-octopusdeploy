package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GoogleContainerRegistry represents a Google Container Registry (GCR).
type GoogleContainerRegistry struct {
	OidcAuthentication *GoogleContainerRegistryOidcAuthentication `json:"OidcAuthentication,omitempty"`

	DockerContainerRegistry
}

type GoogleContainerRegistryOidcAuthentication struct {
	Audience    string   `json:"Audience,omitempty"`
	SubjectKeys []string `json:"SubjectKeys,omitempty"`
}

// NewGoogleContainerRegistry creates and initializes a Google Container Registry (GCR).
func NewGoogleContainerRegistry(name string, username string, password *core.SensitiveValue, oidcAuthentication *GoogleContainerRegistryOidcAuthentication) (*GoogleContainerRegistry, error) {
	if oidcAuthentication == nil {
		if internal.IsEmpty(username) && password.HasValue {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("username")
		}
		if !internal.IsEmpty(username) && !password.HasValue {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("password")
		}
	}

	dockerContainerRegistry, err := NewDockerContainerRegistryWithFeedType(name, FeedTypeGoogleContainerRegistry)

	if err != nil {
		return nil, err
	}

	feed := GoogleContainerRegistry{
		OidcAuthentication:      oidcAuthentication,
		DockerContainerRegistry: *dockerContainerRegistry,
	}

	err = feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this Azure Container Registry (ACR)
// and returns an error if invalid.
func (a *GoogleContainerRegistry) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
