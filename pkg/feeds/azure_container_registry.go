package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AzureContainerRegistry represents an Azure Container Registry (ACR).
type AzureContainerRegistry struct {
	OidcAuthentication *AzureContainerRegistryOidcAuthentication `json:"OidcAuthentication,omitempty"`

	DockerContainerRegistry
}

type AzureContainerRegistryOidcAuthentication struct {
	ClientId    string   `json:"ClientId,omitempty"`
	TenantId    string   `json:"TenantId,omitempty"`
	Audience    string   `json:"Audience,omitempty"`
	SubjectKeys []string `json:"SubjectKeys,omitempty"`
}

// NewAzureContainerRegistry creates and initializes an Azure Container Registry (ACR).
func NewAzureContainerRegistry(name string, username string, password *core.SensitiveValue, oidcAuthentication *AzureContainerRegistryOidcAuthentication) (*AzureContainerRegistry, error) {
	if oidcAuthentication == nil {
		if internal.IsEmpty(username) && password.HasValue {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("username")
		}
		if !internal.IsEmpty(username) && !password.HasValue {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("password")
		}
	}

	dockerContainerRegistry, err := NewDockerContainerRegistryWithFeedType(name, FeedTypeAzureContainerRegistry)

	if err != nil {
		return nil, err
	}

	feed := AzureContainerRegistry{
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
func (a *AzureContainerRegistry) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
