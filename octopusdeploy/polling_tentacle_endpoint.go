package octopusdeploy

import "net/url"

type PollingTentacleEndpoint struct {
	URI *url.URL `json:"Uri" validate:"required,uri"`

	tentacleEndpoint `validate:"required"`
}

func NewPollingTentacleEndpoint(uri *url.URL, thumbprint string) *PollingTentacleEndpoint {
	return &PollingTentacleEndpoint{
		tentacleEndpoint: *newTentacleEndpoint("TentacleActive", thumbprint),
		URI:              uri,
	}
}
