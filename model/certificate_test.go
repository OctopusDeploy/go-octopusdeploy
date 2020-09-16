package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var certificateName = "fake-certificate-name"
var sensitiveValueTestValue = "fake-sensitive-value-name"

func TestEmptyCertificate(t *testing.T) {
	certificate := &Certificate{}

	assert.NotNil(t, certificate)
	assert.Error(t, certificate.Validate())
}

func TestCertificateWithOnlyName(t *testing.T) {
	certificate := &Certificate{Name: certificateName}

	assert.NotNil(t, certificate)
	assert.Error(t, certificate.Validate())
}

func TestCertificateWithNameAndCertificateData(t *testing.T) {
	sensitiveValue := NewSensitiveValue(sensitiveValueTestValue)

	assert.NotNil(t, sensitiveValue)

	certificate := Certificate{Name: certificateName, CertificateData: &sensitiveValue}

	assert.NotNil(t, certificate)
	assert.NoError(t, certificate.Validate())
}
