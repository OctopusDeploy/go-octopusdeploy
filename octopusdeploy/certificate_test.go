package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var certificateName = "fake-certificate-name"
var sensitiveValueTestValue = "fake-sensitive-value-name"

func TestEmptyCertificate(t *testing.T) {
	certificate := &Certificate{}
	require.NotNil(t, certificate)
	require.Error(t, certificate.Validate())
}

func TestCertificateWithOnlyName(t *testing.T) {
	certificate := &Certificate{Name: certificateName}
	require.NotNil(t, certificate)
	require.Error(t, certificate.Validate())
}

func TestCertificateWithNameAndCertificateData(t *testing.T) {
	sensitiveValue := NewSensitiveValue(sensitiveValueTestValue)
	require.NotNil(t, sensitiveValue)

	certificate := Certificate{
		CertificateData:        &sensitiveValue,
		Name:                   certificateName,
		TenantedDeploymentMode: "Untenanted",
	}
	require.NotNil(t, certificate)
	require.NoError(t, certificate.Validate())
}
