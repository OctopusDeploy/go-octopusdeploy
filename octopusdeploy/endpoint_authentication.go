package octopusdeploy

type EndpointAuthentication struct {
	AccountID          string `json:"AccountId,omitempty"`
	ClientCertificate  string `json:"ClientCertificate,omitempty"`
	AuthenticationType string `json:"AuthenticationType,omitempty"`
}
