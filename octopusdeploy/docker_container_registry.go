package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// DockerContainerRegistry represents a Docker container registry.
type DockerContainerRegistry struct {
	APIVersion   string  `json:"ApiVersion,omitempty"`
	FeedURI      *string `json:"FeedUri,omitempty"`
	RegistryPath string  `json:"RegistryPath,omitempty"`

	Feed
}

// NewDockerContainerRegistry creates and initializes a Docker container
// registry.
func NewDockerContainerRegistry(name string, feedURI string) *DockerContainerRegistry {
	return &DockerContainerRegistry{
		Feed: *newFeed(name, FeedTypeDocker),
	}
}

// GetFeedType returns the feed type of this Docker container registry.
func (d *DockerContainerRegistry) GetFeedType() FeedType {
	return d.FeedType
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

var _ IFeed = &DockerContainerRegistry{}
