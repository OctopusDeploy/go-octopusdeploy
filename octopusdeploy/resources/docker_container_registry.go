package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// DockerContainerRegistry represents a Docker container registry.
type DockerContainerRegistry struct {
	APIVersion   string `json:"ApiVersion,omitempty"`
	FeedURI      string `json:"FeedUri,omitempty"`
	RegistryPath string `json:"RegistryPath,omitempty"`

	Feed
}

// NewDockerContainerRegistry creates and initializes a Docker container
// registry.
func NewDockerContainerRegistry(name string) *DockerContainerRegistry {
	dockerContainerRegistry := DockerContainerRegistry{
		Feed: *newFeed(name, FeedTypeDocker),
	}
	dockerContainerRegistry.FeedURI = "https://index.docker.io"

	return &dockerContainerRegistry
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
