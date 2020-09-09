package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewCertificateServiceWithNil(t *testing.T) {
	service := NewCertificateService(nil)
	assert.Nil(t, service)
}

func TestCertificateServiceWithEmptyClient(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestCertificateServiceGetWithEmptyID(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})

	certificate, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, certificate)

	certificate, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceGetWithEmptyName(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})

	certificate, err := service.GetByName("")

	assert.Error(t, err)
	assert.Nil(t, certificate)

	certificate, err = service.GetByName(" ")

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceAddWithEmptyCertificate(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})

	certificate, err := service.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceDeleteWithEmptyID(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})

	err := service.Delete("")

	assert.Error(t, err)

	err = service.Delete(" ")

	assert.Error(t, err)
}

func TestCertificateServiceUpdateWithEmptyCertificate(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})

	certificate, err := service.Update(model.Certificate{})

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceReplaceWithInvalidArguments(t *testing.T) {
	service := NewCertificateService(&sling.Sling{})

	certificate, err := service.Replace("fake-id-string", nil)

	assert.Error(t, err)
	assert.Nil(t, certificate)

	replacementCertificate, err := model.NewCertificateReplace("fake-name-string", "fake-password-string")

	assert.NoError(t, err)
	assert.NotNil(t, replacementCertificate)

	certificate, err = service.Replace(" ", replacementCertificate)

	assert.Error(t, err)
	assert.Nil(t, certificate)
}
