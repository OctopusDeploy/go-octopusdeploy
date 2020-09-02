package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMachineEndpoint(t *testing.T) {
	machineEndpoint := &MachineEndpoint{}
	assert.Error(t, machineEndpoint.Validate())
}

func TestInvalidEndpointURI(t *testing.T) {
	endpointURI := "x"
	endpoint := &MachineEndpoint{
		CommunicationStyle: "None",
		URI:                endpointURI,
	}

	assert.NotNil(t, endpoint)
	assert.Error(t, endpoint.Validate())
}
