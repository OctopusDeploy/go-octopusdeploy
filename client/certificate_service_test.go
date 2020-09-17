package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestCertificateServiceURITemplate = "certificate-service"
)

func TestNewCertificateService(t *testing.T) {
	service := NewCertificateService(nil, "")
	assert.Nil(t, service)
	createCertificateService(t)
}

func TestCertificateServiceGetWithEmptyID(t *testing.T) {
	service := createCertificateService(t)

	certificate, err := service.Get("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, certificate)

	certificate, err = service.Get(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, certificate)
}

func TestCertificateServiceGetWithEmptyName(t *testing.T) {
	service := createCertificateService(t)

	certificate, err := service.GetByName("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("GetByName", "name"))
	assert.Nil(t, certificate)

	certificate, err = service.GetByName(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("GetByName", "name"))
	assert.Nil(t, certificate)
}

func TestCertificateServiceAddWithEmptyCertificate(t *testing.T) {
	service := createCertificateService(t)

	certificate, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError("Add", "certificate"))

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceAddWithInvalidCertificate(t *testing.T) {
	service := createCertificateService(t)

	certificate, err := service.Add(&model.Certificate{})

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceDeleteWithEmptyID(t *testing.T) {
	service := createCertificateService(t)

	err := service.Delete("")
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))

	assert.Error(t, err)

	err = service.Delete(" ")
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))

	assert.Error(t, err)
}

func TestCertificateServiceUpdateWithEmptyCertificate(t *testing.T) {
	service := createCertificateService(t)

	certificate, err := service.Update(model.Certificate{})

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func TestCertificateServiceReplaceWithInvalidArguments(t *testing.T) {
	service := createCertificateService(t)

	certificate, err := service.Replace("", nil)

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Replace", "certificateID"))
	assert.Nil(t, certificate)

	certificate, err = service.Replace(" ", nil)

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Replace", "certificateID"))
	assert.Nil(t, certificate)

	certificate, err = service.Replace("fake-id-string", nil)

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Replace", "replacementCertificate"))
	assert.Nil(t, certificate)

	replacementCertificate := model.NewReplacementCertificate("fake-name-string", "fake-password-string")

	assert.NotNil(t, replacementCertificate)

	certificate, err = service.Replace(" ", replacementCertificate)

	assert.Error(t, err)
	assert.Nil(t, certificate)
}

func createCertificateService(t *testing.T) *CertificateService {
	service := NewCertificateService(&sling.Sling{}, TestCertificateServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestCertificateServiceURITemplate)
	assert.Equal(t, service.name, "CertificateService")

	return service
}
