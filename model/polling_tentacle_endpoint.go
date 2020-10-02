package model

type PollingTentacleEndpoint struct {
	URI string `json:"Uri" validate:"uri"`

	tentacleEndpoint
}

func NewPollingTentacleEndpoint(thumbprint string) *PollingTentacleEndpoint {
	resource := &PollingTentacleEndpoint{}
	resource.CommunicationStyle = "TentacleActive"
	resource.Thumbprint = thumbprint

	return resource
}
