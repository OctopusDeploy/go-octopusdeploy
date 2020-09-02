package model

type MachineEndpointAuthentication struct {
	AccountID          string `json:"AccountId,omitempty"`
	ClientCertificate  string `json:"ClientCertificate,omitempty"`
	AuthenticationType string `json:"AuthenticationType,omitempty"`
}
