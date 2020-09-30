package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewOfflineDropMachineEndpoint(uri string, thumbprint string, communicationStyle enum.CommunicationStyle, proxyID string, defaultWorkerPoolID string) (*MachineEndpoint, error) {
	return &MachineEndpoint{
		CommunicationStyle:  communicationStyle,
		DefaultWorkerPoolID: defaultWorkerPoolID,
		ProxyID:             &proxyID,
		Thumbprint:          thumbprint,
		URI:                 uri,
	}, nil
}
