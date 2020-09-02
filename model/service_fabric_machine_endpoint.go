package model

type ServiceFabricMachineEndpoint struct {
	ConnectionEndpoint        string         `json:"ConnectionEndpoint,omitempty"`
	SecurityMode              string         `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	ServerCertThumbprint      string         `json:"ServerCertThumbprint,omitempty"`
	ClientCertVariable        string         `json:"ClientCertVariable,omitempty"`
	CertificateStoreLocation  string         `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName      string         `json:"CertificateStoreName,omitempty"`
	AadCredentialType         string         `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadClientCredentialSecret string         `json:"AadClientCredentialSecret,omitempty"`
	AadUserCredentialUsername string         `json:"AadUserCredentialUsername,omitempty"`
	AadUserCredentialPassword SensitiveValue `json:"AadUserCredentialPassword,omitempty"`
}
