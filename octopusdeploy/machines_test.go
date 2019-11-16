package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMachineValues(t *testing.T) {

	machineValid := &Machine{
		URI: "x",
		Endpoint: &MachineEndpoint{
			URI:        "x",
			Thumbprint: "1",
		},
		TenantedDeploymentParticipation: Untenanted,
		Status:                          "Unknown",
		Thumbprint:                      "1",
	}

	assert.Nil(t, ValidateMachineValues(machineValid))

	machineInvalidBadURL := &Machine{
		URI: "x",
		Endpoint: &MachineEndpoint{
			URI: "y",
		},
		TenantedDeploymentParticipation: Untenanted,
		Status:                          "Unknown",
	}

	assert.Error(t, ValidateMachineValues(machineInvalidBadURL))

	machineInvalidBadTenatedDeployment := &Machine{
		URI: "x",
		Endpoint: &MachineEndpoint{
			URI: "y",
		},
		TenantedDeploymentParticipation: 5,
		Status:                          "Unknown",
	}

	assert.Error(t, ValidateMachineValues(machineInvalidBadTenatedDeployment))

	machineInvalidNonMatchingThumbprint := &Machine{
		URI: "x",
		Endpoint: &MachineEndpoint{
			URI:        "y",
			Thumbprint: "1",
		},
		TenantedDeploymentParticipation: 5,
		Status:                          "Unknown",
		Thumbprint:                      "2",
	}

	assert.Error(t, ValidateMachineValues(machineInvalidNonMatchingThumbprint))
}
