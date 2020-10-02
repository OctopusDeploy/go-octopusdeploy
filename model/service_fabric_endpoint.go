package model

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
	resource := &ServiceFabricEndpoint{}

	return resource
}
