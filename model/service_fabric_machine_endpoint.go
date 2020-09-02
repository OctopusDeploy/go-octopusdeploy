package model

type ServiceFabricMachineEndpoint struct {
	ConnectionEndpoint        string         `json:ConnectionEndpoint`
	SecurityMode              string         `json:SecurityMode validate:"oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	ServerCertThumbprint      string         `json:ServerCertThumbprint`
	ClientCertVariable        string         `json:ClientCertVariable`
	CertificateStoreLocation  string         `json:CertificateStoreLocation`
	CertificateStoreName      string         `json:CertificateStoreName`
	AadCredentialType         string         `json:AadCredentialType validate:"oneof=ClientCredential UserCredential"`
	AadClientCredentialSecret string         `json:AadClientCredentialSecret`
	AadUserCredentialUsername string         `json:AadUserCredentialUsername`
	AadUserCredentialPassword SensitiveValue `json:AadUserCredentialPassword`
}
