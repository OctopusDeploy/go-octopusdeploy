package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type MachineEndpoint struct {
	AccountID                     string                        `json:"AccountId,omitempty"`
	CommunicationStyle            enum.CommunicationStyle       `json:"CommunicationStyle" validate:"required,oneof=None TentaclePassive TentacleActive Ssh OfflineDrop AzureWebApp Ftp AzureCloudService AzureServiceFabricCluster Kubernetes"`
	DefaultWorkerPoolID           string                        `json:"DefaultWorkerPoolId,omitempty"`
	ProxyID                       *string                       `json:"ProxyId"`
	Thumbprint                    string                        `json:"Thumbprint,omitempty"`
	TentacleVersionDetails        MachineTentacleVersionDetails `json:"TentacleVersionDetails"`
	CertificateSignatureAlgorithm string                        `json:"CertificateSignatureAlgorithm,omitempty"`
	URI                           string                        `json:"Uri,omitempty" validate:"omitempty,uri"` // This is not in the spec doc, but it shows up and needs to be kept in sync
	AzureWebAppMachineEndpoint
	CloudRegionMachineEndpoint
	CloudServiceMachineEndpoint
	KubernetesMachineEndpoint
	ListeningTentacleMachineEndpoint
	OfflineDropMachineEndpoint
	PollingTentacleMachineEndpoint
	ServiceFabricMachineEndpoint
	SshMachineEndpoint

	Resource
}

// NewMachineEndpoint initializes a MachineEndpoint.
func NewMachineEndpoint(uri string, thumbprint string, communicationStyle enum.CommunicationStyle, proxyID string, defaultWorkerPoolID string) (*MachineEndpoint, error) {
	return &MachineEndpoint{
		URI:                 uri,
		Thumbprint:          thumbprint,
		CommunicationStyle:  communicationStyle,
		ProxyID:             &proxyID,
		DefaultWorkerPoolID: defaultWorkerPoolID,
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
