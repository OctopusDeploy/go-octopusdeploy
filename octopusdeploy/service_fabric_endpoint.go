package octopusdeploy

import "github.com/go-playground/validator/v10"

type AzureServiceFabricEndpoint struct {
	AadClientCredentialSecret   string          `json:"AadClientCredentialSecret,omitempty"`
	AadCredentialType           string          `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadUserCredentialPassword   *SensitiveValue `json:"AadUserCredentialPassword,omitempty"`
	AadUserCredentialUsername   string          `json:"AadUserCredentialUsername,omitempty"`
	CertificateStoreLocation    string          `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName        string          `json:"CertificateStoreName,omitempty"`
	ClientCertificateVariable   string          `json:"ClientCertVariable,omitempty"`
	ConnectionEndpoint          string          `json:"ConnectionEndpoint,omitempty"`
	SecurityMode                string          `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	ServerCertificateThumbprint string          `json:"ServerCertThumbprint,omitempty"`

	endpoint
}

func NewAzureServiceFabricEndpoint() *AzureServiceFabricEndpoint {
	return &AzureServiceFabricEndpoint{
		endpoint: *newEndpoint("AzureServiceFabricCluster"),
	}
}

// Validate checks the state of the service fabric endpoint and returns an
// error if invalid.
func (s *AzureServiceFabricEndpoint) Validate() error {
	return validator.New().Struct(s)
}
