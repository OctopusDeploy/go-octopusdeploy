package model

type OfflineDropEndpoint struct {
	ApplicationsDirectory                string                  `json:"ApplicationsDirectory,omitempty"`
	CommunicationStyle                   string                  `validate:"required,eq=OfflineDrop"`
	Destination                          *OfflineDropDestination `json:"Destination"`
	SensitiveVariablesEncryptionPassword SensitiveValue          `json:"SensitiveVariablesEncryptionPassword"`
	OctopusWorkingDirectory              string                  `json:"OctopusWorkingDirectory,omitempty"`

	endpoint
}

func NewOfflineDropEndpoint() *OfflineDropEndpoint {
	offlineDropEndpoint := &OfflineDropEndpoint{
		CommunicationStyle: "OfflineDrop",
	}
	offlineDropEndpoint.endpoint.CommunicationStyle = offlineDropEndpoint.CommunicationStyle

	return offlineDropEndpoint
}
