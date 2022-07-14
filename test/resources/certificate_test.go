package resources

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/certificates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestEmptyCertificateResource(t *testing.T) {
	certificate := &certificates.CertificateResource{}
	require.NotNil(t, certificate)
	require.Error(t, certificate.Validate())
}

func TestCertificateResourceWithOnlyName(t *testing.T) {
	name := internal.GetRandomName()
	certificate := &certificates.CertificateResource{Name: name}
	require.NotNil(t, certificate)
	require.Error(t, certificate.Validate())
}

func TestCertificateResourceWithNameAndCertificateData(t *testing.T) {
	name := internal.GetRandomName()
	newValue := internal.GetRandomName()
	sensitiveValue := core.NewSensitiveValue(newValue)
	require.NotNil(t, sensitiveValue)

	certificateResource := certificates.CertificateResource{
		CertificateData:        sensitiveValue,
		Name:                   name,
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, certificateResource)
	require.NoError(t, certificateResource.Validate())
}
