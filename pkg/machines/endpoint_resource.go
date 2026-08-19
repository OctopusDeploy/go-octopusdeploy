package machines

import (
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type EndpointResource struct {
	AadClientCredentialSecret            string                                 `json:"AadClientCredentialSecret,omitempty"`
	AadCredentialType                    string                                 `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadUserCredentialUsername            string                                 `json:"AadUserCredentialUsername,omitempty"`
	AadUserCredentialPassword            *core.SensitiveValue                   `json:"AadUserCredentialPassword,omitempty"`
	AccountID                            string                                 `json:"AccountId"`
	ApplicationsDirectory                string                                 `json:"ApplicationsDirectory,omitempty"`
	AssumedRoleARN                       string                                 `json:"AssumedRoleArn,omitempty"`
	AssumedRoleSession                   string                                 `json:"AssumedRoleSession,omitempty"`
	AssumeRole                           bool                                   `json:"AssumeRole"`
	AssumeRoleExternalID                 string                                 `json:"AssumeRoleExternalId,omitempty"`
	AssumeRoleSessionDurationSeconds     int                                    `json:"AssumeRoleSessionDurationSeconds,omitempty"`
	Authentication                       IKubernetesAuthentication              `json:"Authentication,omitempty"`
	CertificateSignatureAlgorithm        string                                 `json:"CertificateSignatureAlgorithm,omitempty"`
	CertificateStoreLocation             string                                 `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName                 string                                 `json:"CertificateStoreName,omitempty"`
	ClientCertificateVariable            string                                 `json:"ClientCertVariable,omitempty"`
	CloudServiceName                     string                                 `json:"CloudServiceName"`
	ClusterCertificate                   string                                 `json:"ClusterCertificate,omitempty"`
	ClusterCertificatePath               string                                 `json:"ClusterCertificatePath,omitempty"`
	ClusterName                          string                                 `json:"ClusterName,omitempty"`
	ClusterURL                           *url.URL                               `json:"ClusterUrl"`
	CommunicationStyle                   string                                 `json:"CommunicationStyle" validate:"required,oneof=AzureCloudService AzureServiceFabricCluster AzureWebApp Ftp Kubernetes None OfflineDrop Ssh TentacleActive TentaclePassive KubernetesTentacle AwsEcsCluster"`
	ConnectionEndpoint                   string                                 `json:"ConnectionEndpoint,omitempty"`
	Container                            *deployments.DeploymentActionContainer `json:"Container,omitempty"`
	ContainerOptions                     string                                 `json:"ContainerOptions,omitempty"`
	DefaultWorkerPoolID                  string                                 `json:"DefaultWorkerPoolId"`
	Destination                          *OfflinePackageDropDestination         `json:"Destination"`
	DotNetCorePlatform                   string                                 `json:"DotNetCorePlatform,omitempty"`
	Fingerprint                          string                                 `json:"Fingerprint,omitempty"`
	Host                                 string                                 `json:"Host,omitempty"`
	Namespace                            string                                 `json:"Namespace,omitempty"`
	Port                                 int                                    `json:"Port,omitempty"`
	ProxyID                              string                                 `json:"ProxyId,omitempty"`
	Region                               string                                 `json:"Region,omitempty"`
	ResourceGroupName                    string                                 `json:"ResourceGroupName,omitempty"`
	RunningInContainer                   bool                                   `json:"RunningInContainer"`
	SecurityMode                         string                                 `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	SensitiveVariablesEncryptionPassword *core.SensitiveValue                   `json:"SensitiveVariablesEncryptionPassword"`
	ServerCertificateThumbprint          string                                 `json:"ServerCertThumbprint,omitempty"`
	SkipTLSVerification                  bool                                   `json:"SkipTlsVerification"`
	Slot                                 string                                 `json:"Slot"`
	StorageAccountName                   string                                 `json:"StorageAccountName"`
	SwapIfPossible                       bool                                   `json:"SwapIfPossible"`
	TentacleVersionDetails               *TentacleVersionDetails                `json:"TentacleVersionDetails,omitempty"`
	Thumbprint                           string                                 `json:"Thumbprint"`
	WorkingDirectory                     string                                 `json:"OctopusWorkingDirectory,omitempty"`
	UseCurrentInstanceCount              bool                                   `json:"UseCurrentInstanceCount"`
	UseInstanceRole                      bool                                   `json:"UseInstanceRole"`
	URI                                  *url.URL                               `json:"Uri"`
	WebAppName                           string                                 `json:"WebAppName,omitempty"`
	WebAppSlotName                       string                                 `json:"WebAppSlotName"`

	resources.Resource
}

// NewEndpoint creates and initializes an account resource with a name and type.
func NewEndpointResource(communicationStyle string) *EndpointResource {
	return &EndpointResource{
		CommunicationStyle: communicationStyle,
		Resource:           *resources.NewResource(),
	}
}

// GetCommunicationStyle returns the communication style of this endpoint.
func (e *EndpointResource) GetCommunicationStyle() string {
	return e.CommunicationStyle
}

// Validate checks the state of the endpoint resource and returns an error if
// invalid.
func (e EndpointResource) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(validateEndpointResource, EndpointResource{})

	return validate.Struct(e)
}

// validateEndpointResource requires the fields that the communication style in
// hand actually uses. A blanket "required" on ClusterUrl, Thumbprint and Uri
// asked every style for all three, which no single endpoint has.
func validateEndpointResource(sl validator.StructLevel) {
	resource := sl.Current().Interface().(EndpointResource)

	switch resource.CommunicationStyle {
	case "Kubernetes":
		if resource.ClusterURL == nil {
			sl.ReportError(resource.ClusterURL, "ClusterUrl", "ClusterURL", "required", "")
		}
	case "TentacleActive", "TentaclePassive", "KubernetesTentacle":
		if resource.URI == nil {
			sl.ReportError(resource.URI, "Uri", "URI", "required", "")
		}
		if resource.Thumbprint == "" {
			sl.ReportError(resource.Thumbprint, "Thumbprint", "Thumbprint", "required", "")
		}
	}
}

var _ IEndpoint = &EndpointResource{}
