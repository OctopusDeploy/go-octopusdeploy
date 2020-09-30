package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewCloudRegionMachineEndpoint(uri string, thumbprint string, proxyID string, defaultWorkerPoolID string) (*MachineEndpoint, error) {
	return &MachineEndpoint{
		CommunicationStyle:  enum.AzureCloudService,
		DefaultWorkerPoolID: defaultWorkerPoolID,
		ProxyID:             &proxyID,
		Thumbprint:          thumbprint,
		URI:                 uri,
	}, nil
}
