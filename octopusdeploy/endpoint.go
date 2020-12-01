package octopusdeploy

import "github.com/go-playground/validator/v10"

// endpoint is the base definition of an endpoint and describes its
// communication style (SSH, Kubernetes, etc.)
type endpoint struct {
	CommunicationStyle string `json:"CommunicationStyle" validate:"required,oneof=AzureCloudService AzureServiceFabricCluster AzureWebApp Ftp Kubernetes None OfflineDrop Ssh TentacleActive TentaclePassive"`

	resource
}

// newEndpoint creates and initializes a new endpoint.
func newEndpoint(communicationStyle string) *endpoint {
	endpoint := &endpoint{
		CommunicationStyle: communicationStyle,
		resource:           *newResource(),
	}
	return endpoint
}

// GetCommunicationStyle returns the communication style of this endpoint.
func (e endpoint) GetCommunicationStyle() string {
	return e.CommunicationStyle
}

// Validate checks the state of the endpoint and returns an error if invalid.
func (e endpoint) Validate() error {
	return validator.New().Struct(e)
}

var _ IEndpoint = &endpoint{}
