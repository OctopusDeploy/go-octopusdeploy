package model

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/stretchr/testify/assert"
)

func TestEmptyMachine(t *testing.T) {
	machine := &Machine{}

	assert.NotNil(t, machine)
	assert.Error(t, machine.Validate())
}

func TestMissingDeploymentModeAndEmptyEndpoint(t *testing.T) {
	machine := &Machine{
		Endpoint: &MachineEndpoint{},
	}

	assert.NotNil(t, machine)
	assert.Error(t, machine.Validate())
}

func TestEmptyEndpoint(t *testing.T) {
	machine := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint:       &MachineEndpoint{},
	}

	assert.NotNil(t, machine)
	assert.Error(t, machine.Validate())
}

func TestValidEndpointAndInvalidDeploymentMode(t *testing.T) {
	machine := &Machine{
		DeploymentMode: "invalid",
		Endpoint: &MachineEndpoint{
			CommunicationStyle: enum.NoCommunicationStyle,
		},
	}

	assert.NotNil(t, machine)
	assert.Error(t, machine.Validate())
}

func TestValidDeploymentModeAndEndpoint(t *testing.T) {
	machine := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint: &MachineEndpoint{
			CommunicationStyle: enum.NoCommunicationStyle,
		},
	}

	assert.NotNil(t, machine)
	assert.NoError(t, machine.Validate())
}

func TestInvalidMachineURI(t *testing.T) {
	machineURI := "x"
	machine := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint: &MachineEndpoint{
			CommunicationStyle: enum.NoCommunicationStyle,
		},
		URI: machineURI,
	}

	assert.NotNil(t, machine)
	assert.Error(t, machine.Validate())
}

func TestValidMachineURI(t *testing.T) {
	machineURI := "http://localhost"
	machine := &Machine{
		DeploymentMode: "Untenanted",
		Endpoint: &MachineEndpoint{
			CommunicationStyle: enum.NoCommunicationStyle,
		},
		URI: machineURI,
	}

	assert.NotNil(t, machine)
	assert.NoError(t, machine.Validate())
}
