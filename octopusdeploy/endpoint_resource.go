package octopusdeploy

import (
	"net/url"

	"github.com/go-playground/validator/v10"
)

type EndpointResource struct {
	AadClientCredentialSecret            string                         `json:"AadClientCredentialSecret,omitempty"`
	AadCredentialType                    string                         `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadUserCredentialUsername            string                         `json:"AadUserCredentialUsername,omitempty"`
	AadUserCredentialPassword            SensitiveValue                 `json:"AadUserCredentialPassword,omitempty"`
	AccountID                            string                         `json:"AccountId"`
	ApplicationsDirectory                string                         `json:"ApplicationsDirectory,omitempty"`
	Authentication                       IKubernetesAuthentication      `json:"Authentication,omitempty"`
	CertificateSignatureAlgorithm        string                         `json:"CertificateSignatureAlgorithm,omitempty"`
	CertificateStoreLocation             string                         `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName                 string                         `json:"CertificateStoreName,omitempty"`
	ClientCertificateVariable            string                         `json:"ClientCertVariable,omitempty"`
	CloudServiceName                     string                         `json:"CloudServiceName"`
	ClusterCertificate                   string                         `json:"ClusterCertificate,omitempty"`
	ClusterURL                           *url.URL                       `json:"ClusterUrl" validate:"required,url"`
	CommunicationStyle                   string                         `json:"CommunicationStyle" validate:"required,oneof=AzureCloudService AzureServiceFabricCluster Ftp Kubernetes None OfflineDrop Ssh TentacleActive TentaclePassive"`
	ConnectionEndpoint                   string                         `json:"ConnectionEndpoint,omitempty"`
	Container                            DeploymentActionContainer      `json:"Container,omitempty"`
	DefaultWorkerPoolID                  string                         `json:"DefaultWorkerPoolId"`
	Destination                          *OfflinePackageDropDestination `json:"Destination"`
	DotNetCorePlatform                   string                         `json:"DotNetCorePlatform,omitempty"`
	Fingerprint                          string                         `json:"Fingerprint,omitempty"`
	Host                                 string                         `json:"Host,omitempty"`
	Namespace                            string                         `json:"Namespace,omitempty"`
	Port                                 int                            `json:"Port,omitempty"`
	ProxyID                              string                         `json:"ProxyId,omitempty"`
	ResourceGroupName                    string                         `json:"ResourceGroupName,omitempty"`
	RunningInContainer                   bool                           `json:"RunningInContainer"`
	SecurityMode                         string                         `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	SensitiveVariablesEncryptionPassword SensitiveValue                 `json:"SensitiveVariablesEncryptionPassword"`
	ServerCertificateThumbprint          string                         `json:"ServerCertThumbprint,omitempty"`
	SkipTLSVerification                  bool                           `json:"SkipTlsVerification"`
	Slot                                 string                         `json:"Slot"`
	StorageAccountName                   string                         `json:"StorageAccountName"`
	SwapIfPossible                       bool                           `json:"SwapIfPossible"`
	TentacleVersionDetails               *TentacleVersionDetails        `json:"TentacleVersionDetails,omitempty"`
	Thumbprint                           string                         `json:"Thumbprint" validate:"required"`
	WorkingDirectory                     string                         `json:"OctopusWorkingDirectory,omitempty"`
	UseCurrentInstanceCount              bool                           `json:"UseCurrentInstanceCount"`
	URI                                  *url.URL                       `json:"Uri" validate:"required,uri"`
	WebAppName                           string                         `json:"WebAppName,omitempty"`
	WebAppSlotName                       string                         `json:"WebAppSlotName"`

	resource
}

type EndpointResources struct {
	Items []*EndpointResource `json:"Items"`
	PagedResults
}

// NewEndpoint creates and initializes an account resource with a name and type.
func NewEndpointResource(communicationStyle string) *EndpointResource {
	return &EndpointResource{
		CommunicationStyle: communicationStyle,
		resource:           *newResource(),
	}
}

// GetCommunicationStyle returns the communication style of this endpoint.
func (e *EndpointResource) GetCommunicationStyle() string {
	return e.CommunicationStyle
}

// Validate checks the state of the endpoint resource and returns an error if
// invalid.
func (e EndpointResource) Validate() error {
	return validator.New().Struct(e)
}

var _ IEndpoint = &EndpointResource{}
