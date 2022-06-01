package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// DockerContainerRegistry represents a Docker container registry.
type DockerContainerRegistry struct {
	APIVersion   string `json:"ApiVersion,omitempty"`
	FeedURI      string `json:"FeedUri,omitempty"`
	RegistryPath string `json:"RegistryPath,omitempty"`

	feed
}

// NewDockerContainerRegistry creates and initializes a Docker container
// registry.
func NewDockerContainerRegistry(name string) (*DockerContainerRegistry, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	feed := DockerContainerRegistry{
		FeedURI: "https://index.docker.io",
		feed:    *newFeed(name, FeedTypeDocker),
	}

	// validate to ensure that all expectations are met
	err := feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this Docker container registry and returns an
// error if invalid.
func (d *DockerContainerRegistry) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(d)
}
