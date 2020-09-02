package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMachine(t *testing.T) {
	machineValid := &Machine{}
	assert.Error(t, machineValid.Validate())
}

func TestMissingDeploymentModeAndEmptyEndpoint(t *testing.T) {
	machineValid := &Machine{
		Endpoint: MachineEndpoint{},
	}
	assert.Error(t, machineValid.Validate())
}

func TestEmptyEndpoint(t *testing.T) {
	machineValid := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint:       MachineEndpoint{},
	}
	assert.Error(t, machineValid.Validate())
}

func TestValidEndpointAndInvalidDeploymentMode(t *testing.T) {
	machineValid := &Machine{
		DeploymentMode: "invalid",
		Endpoint: MachineEndpoint{
			CommunicationStyle: "None",
		},
	}
	assert.Error(t, machineValid.Validate())
}

func TestValidDeploymentModeAndEndpoint(t *testing.T) {
	machineValid := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint: MachineEndpoint{
			CommunicationStyle: "None",
		},
	}
	assert.Nil(t, machineValid.Validate())
}

func TestValidateMachineValues(t *testing.T) {

	machineValid := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint: MachineEndpoint{
			CommunicationStyle: "None",
			Thumbprint:         "1",
			URI:                "x",
		},
		Status:     "Unknown",
		Thumbprint: "1",
		URI:        "x",
	}

	assert.Nil(t, machineValid.Validate())

	machineInvalidBadURL := &Machine{
		URI: "x",
		Endpoint: MachineEndpoint{
			URI: "y",
		},
		DeploymentMode: "Untenanted",
		Status:         "Unknown",
	}

	assert.Error(t, machineInvalidBadURL.Validate())

	machineInvalidBadTenantedDeployment := &Machine{
		URI: "x",
		Endpoint: MachineEndpoint{
			URI: "y",
		},
		DeploymentMode: "invalid mode",
		Status:         "Unknown",
	}

	assert.Error(t, machineInvalidBadTenantedDeployment.Validate())

	machineInvalidNonMatchingThumbprint := &Machine{
		URI: "x",
		Endpoint: MachineEndpoint{
			URI:        "y",
			Thumbprint: "1",
		},
		Status:     "Unknown",
		Thumbprint: "2",
	}

	assert.Error(t, machineInvalidNonMatchingThumbprint.Validate())
}
