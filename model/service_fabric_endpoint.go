package model

import "github.com/go-playground/validator/v10"

type ServiceFabricEndpoint struct {
	AadClientCredentialSecret   string         `json:"AadClientCredentialSecret,omitempty"`
	AadCredentialType           string         `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadUserCredentialUsername   string         `json:"AadUserCredentialUsername,omitempty"`
	AadUserCredentialPassword   SensitiveValue `json:"AadUserCredentialPassword,omitempty"`
	CertificateStoreLocation    string         `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName        string         `json:"CertificateStoreName,omitempty"`
	ClientCertificateVariable   string         `json:"ClientCertVariable,omitempty"`
	ConnectionEndpoint          string         `json:"ConnectionEndpoint,omitempty"`
	SecurityMode                string         `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	ServerCertificateThumbprint string         `json:"ServerCertThumbprint,omitempty"`

	endpoint
}

func NewServiceFabricEndpoint() *ServiceFabricEndpoint {
	return &ServiceFabricEndpoint{
		endpoint: *newEndpoint("AzureServiceFabricCluster"),
	}
}

// Validate checks the state of the service fabric endpoint and returns an
// error if invalid.
func (s *ServiceFabricEndpoint) Validate() error {
	return validator.New().Struct(s)
}
