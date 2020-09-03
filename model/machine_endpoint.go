package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type MachineEndpoint struct {
	AccountID                     string                        `json:"AccountId,omitempty"`
	CommunicationStyle            string                        `json:"CommunicationStyle" validate:"required,oneof=None TentaclePassive TentacleActive Ssh OfflineDrop AzureWebApp Ftp AzureCloudService AzureServiceFabricCluster Kubernetes"`
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

func NewMachineEndpoint() (*MachineEndpoint, error) {
	return &MachineEndpoint{}, nil
}

func (t *MachineEndpoint) Validate() error {

	validate := validator.New()
	err := validate.Struct(t)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}

		return err
	}

	return nil
}
