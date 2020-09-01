package model

type MachineEndpointAuthentication struct {
	AccountID          string `json:"AccountId"`
	ClientCertificate  string `json:"ClientCertificate"`
	AuthenticationType string `json:"AuthenticationType"`
}
