package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyCertificateResource(t *testing.T) {
	certificate := &CertificateResource{}
	require.NotNil(t, certificate)
	require.Error(t, certificate.Validate())
}

func TestCertificateResourceWithOnlyName(t *testing.T) {
	name := octopusdeploy.getRandomName()
	certificate := &CertificateResource{Name: name}
	require.NotNil(t, certificate)
	require.Error(t, certificate.Validate())
}

func TestCertificateResourceWithNameAndCertificateData(t *testing.T) {
	name := octopusdeploy.getRandomName()
	newValue := octopusdeploy.getRandomName()
	sensitiveValue := NewSensitiveValue(newValue)
	require.NotNil(t, sensitiveValue)

	certificateResource := CertificateResource{
		CertificateData:        sensitiveValue,
		Name:                   name,
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, certificateResource)
	require.NoError(t, certificateResource.Validate())
}
