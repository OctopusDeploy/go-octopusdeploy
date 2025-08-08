package observability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterKubernetesMonitorCommand_Validate(t *testing.T) {
	// Test valid command
	command := &RegisterKubernetesMonitorCommand{
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		SpaceID:        "Spaces-1",
		MachineID:      "Machines-1",
	}
	err := command.Validate()
	assert.NoError(t, err)

	// Test invalid command (missing InstallationID)
	command2 := &RegisterKubernetesMonitorCommand{
		SpaceID:   "Spaces-1",
		MachineID: "Machines-1",
	}
	err = command2.Validate()
	assert.Error(t, err)

	// Test invalid command (missing SpaceID)
	command3 := &RegisterKubernetesMonitorCommand{
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		MachineID:      "Machines-1",
	}
	err = command3.Validate()
	assert.Error(t, err)

	// Test invalid command (missing MachineID)
	command4 := &RegisterKubernetesMonitorCommand{
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		SpaceID:        "Spaces-1",
	}
	err = command4.Validate()
	assert.Error(t, err)
}

func TestRegisterKubernetesMonitorResponse_Validate(t *testing.T) {
	// Test valid response
	resource := &KubernetesMonitorResource{
		ID:             KubernetesMonitorId("km-123"),
		SpaceID:        "Spaces-1",
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		MachineID:      "Machines-1",
	}
	response := &RegisterKubernetesMonitorResponse{
		Resource:              resource,
		AuthenticationToken:   "auth-token-123",
		CertificateThumbprint: "cert-thumbprint-123",
	}
	err := response.Validate()
	assert.NoError(t, err)

	// Test invalid response (missing Resource)
	response2 := &RegisterKubernetesMonitorResponse{
		AuthenticationToken:   "auth-token-123",
		CertificateThumbprint: "cert-thumbprint-123",
	}
	err = response2.Validate()
	assert.Error(t, err)

	// Test invalid response (missing AuthenticationToken)
	response3 := &RegisterKubernetesMonitorResponse{
		Resource:              resource,
		CertificateThumbprint: "cert-thumbprint-123",
	}
	err = response3.Validate()
	assert.Error(t, err)

	// Test invalid response (missing CertificateThumbprint)  
	response4 := &RegisterKubernetesMonitorResponse{
		Resource:            resource,
		AuthenticationToken: "auth-token-123",
	}
	err = response4.Validate()
	assert.Error(t, err)
}

func TestKubernetesMonitorResource_Validate(t *testing.T) {
	// Test valid resource
	resource := &KubernetesMonitorResource{
		ID:             KubernetesMonitorId("km-123"),
		SpaceID:        "Spaces-1",
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		MachineID:      "Machines-1",
	}
	err := resource.Validate()
	assert.NoError(t, err)

	// Test invalid resource (missing ID)
	resource2 := &KubernetesMonitorResource{
		SpaceID:        "Spaces-1",
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		MachineID:      "Machines-1",
	}
	err = resource2.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing SpaceID)
	resource3 := &KubernetesMonitorResource{
		ID:             KubernetesMonitorId("km-123"),
		InstallationID: KubernetesMonitorInstallationId("install-123"),
		MachineID:      "Machines-1",
	}
	err = resource3.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing InstallationID)
	resource4 := &KubernetesMonitorResource{
		ID:        KubernetesMonitorId("km-123"),
		SpaceID:   "Spaces-1",
		MachineID: "Machines-1",
	}
	err = resource4.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing MachineID)
	resource5 := &KubernetesMonitorResource{
		ID:             KubernetesMonitorId("km-123"),
		SpaceID:        "Spaces-1",
		InstallationID: KubernetesMonitorInstallationId("install-123"),
	}
	err = resource5.Validate()
	assert.Error(t, err)
}