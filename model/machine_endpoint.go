package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type MachineEndpoint struct {
	AadCredentialType                    string                         `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadClientCredentialSecret            string                         `json:"AadClientCredentialSecret,omitempty"`
	AadUserCredentialUsername            string                         `json:"AadUserCredentialUsername,omitempty"`
	AadUserCredentialPassword            SensitiveValue                 `json:"AadUserCredentialPassword,omitempty"`
	AccountID                            string                         `json:"AccountId,omitempty"`
	ApplicationsDirectory                string                         `json:"ApplicationsDirectory,omitempty"`
	Authentication                       *MachineEndpointAuthentication `json:"Authentication,omitempty"`
	CertificateSignatureAlgorithm        string                         `json:"CertificateSignatureAlgorithm,omitempty"`
	CertificateStoreLocation             string                         `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName                 string                         `json:"CertificateStoreName,omitempty"`
	ClientCertVariable                   string                         `json:"ClientCertVariable,omitempty"`
	CloudServiceName                     string                         `json:"CloudServiceName,omitempty"`
	ClusterCertificate                   string                         `json:"ClusterCertificate,omitempty"`
	ClusterURL                           string                         `json:"ClusterUrl,omitempty" validate:"omitempty,url"`
	CommunicationStyle                   enum.CommunicationStyle        `json:"CommunicationStyle" validate:"required"`
	ConnectionEndpoint                   string                         `json:"ConnectionEndpoint,omitempty"`
	Container                            DeploymentActionContainer      `json:"Container"`
	DefaultWorkerPoolID                  string                         `json:"DefaultWorkerPoolId"`
	Destination                          OfflineDropDestination         `json:"Destination"`
	DotNetCorePlatform                   string                         `json:"DotNetCorePlatform,omitempty"`
	Fingerprint                          string                         `json:"Fingerprint,omitempty"`
	Host                                 string                         `json:"Host,omitempty"`
	Namespace                            string                         `json:"Namespace,omitempty"`
	Port                                 *uint16                        `json:"Port,omitempty"`
	ProxyID                              *string                        `json:"ProxyId"`
	ResourceGroupName                    string                         `json:"ResourceGroupName,omitempty"`
	RunningInContainer                   bool                           `json:"RunningInContainer"`
	SecurityMode                         string                         `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	SensitiveVariablesEncryptionPassword SensitiveValue                 `json:"SensitiveVariablesEncryptionPassword"`
	ServerCertThumbprint                 string                         `json:"ServerCertThumbprint,omitempty"`
	SkipTLSVerification                  string                         `json:"SkipTlsVerification,omitempty"`
	Slot                                 string                         `json:"Slot,omitempty"`
	StorageAccountName                   string                         `json:"StorageAccountName,omitempty"`
	SwapIfPossible                       bool                           `json:"SwapIfPossible"`
	TentacleVersionDetails               MachineTentacleVersionDetails  `json:"TentacleVersionDetails"`
	Thumbprint                           string                         `json:"Thumbprint"`
	URI                                  string                         `json:"Uri,omitempty" validate:"omitempty,uri"` // This is not in the spec doc, but it shows up and needs to be kept in sync
	UseCurrentInstanceCount              bool                           `json:"UseCurrentInstanceCount"`
	WebAppName                           string                         `json:"WebAppName,omitempty"`
	WebAppSlotName                       int                            `json:"WebAppSlotName"`
	WorkingDirectory                     string                         `json:"OctopusWorkingDirectory,omitempty"`

	Resource
}

// NewMachineEndpoint initializes a MachineEndpoint.
func NewMachineEndpoint(uri string, thumbprint string, communicationStyle enum.CommunicationStyle, proxyID string, defaultWorkerPoolID string) (*MachineEndpoint, error) {
	return &MachineEndpoint{
		CommunicationStyle:  communicationStyle,
		DefaultWorkerPoolID: defaultWorkerPoolID,
		ProxyID:             &proxyID,
		Thumbprint:          thumbprint,
		URI:                 uri,
	}, nil
}

// GetID returns the ID value of the MachineEndpoint.
func (resource MachineEndpoint) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this MachineEndpoint.
func (resource MachineEndpoint) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this MachineEndpoint was changed.
func (resource MachineEndpoint) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this MachineEndpoint.
func (resource MachineEndpoint) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the MachineEndpoint and returns an error if invalid.
func (resource MachineEndpoint) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return nil
}

var _ ResourceInterface = &MachineEndpoint{}
