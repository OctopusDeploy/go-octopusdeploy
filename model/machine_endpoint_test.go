package model

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMachineEndpoint(t *testing.T) {
	machineEndpoint := &MachineEndpoint{}
	assert.Error(t, machineEndpoint.Validate())
}

func TestInvalidEndpointURI(t *testing.T) {
	communicationStyle, _ := enum.ParseCommunicationStyle("None")
	endpointURI := "x"
	endpoint := &MachineEndpoint{
		CommunicationStyle: communicationStyle,
		URI:                endpointURI,
	}

	assert.NotNil(t, endpoint)
	assert.Error(t, endpoint.Validate())
}

func TestValidSshEndpointPort(t *testing.T) {
	communicationStyle, _ := enum.ParseCommunicationStyle("None")
	endpoint := &MachineEndpoint{
		CommunicationStyle: communicationStyle,
	}

	var port uint16 = 22
	endpoint.Port = &port

	assert.NotNil(t, endpoint)
	assert.NoError(t, endpoint.Validate())
}
