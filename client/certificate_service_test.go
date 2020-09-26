package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestCertificateService(t *testing.T) {
	t.Run("New", TestNewCertificateService)
	t.Run("GetByID", TestCertificateServiceGetByID)
}

func TestNewCertificateService(t *testing.T) {
	serviceFunction := newCertificateService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceCertificateService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *certificateService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestCertificateServiceGetByID(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)

	for _, resource := range resourceList {
		resourceToCompare, err := service.GetByID(resource.ID)

		assert.NoError(err)
		assert.EqualValues(resource, *resourceToCompare)
	}
}

func TestCertificateServiceGetByPartialName(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)

	if len(resourceList) > 0 {
		resourcesToCompare, err := service.GetByPartialName(resourceList[0].Name)

		assert.NoError(err)
		assert.EqualValues(resourcesToCompare[0], resourceList[0])
	}
}

func TestCertificateServiceGetWithEmptyID(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	certificate, err := service.GetByID(emptyString)

	assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(certificate)

	certificate, err = service.GetByID(whitespaceString)

	assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(certificate)
}

func TestCertificateServiceGetWithEmptyName(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	certificates, err := service.GetByPartialName(emptyString)

	assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.Empty(certificates)

	certificates, err = service.GetByPartialName(whitespaceString)

	assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.Empty(certificates)
}

func TestCertificateServiceAddWithEmptyCertificate(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	certificate, err := service.Add(nil)

	assert.Equal(err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(certificate)
}

func TestCertificateServiceAddWithInvalidCertificate(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	certificate, err := service.Add(&model.Certificate{})

	assert.Error(err)
	assert.Nil(certificate)
}

func TestCertificateServiceDeleteWithEmptyID(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	err := service.DeleteByID(emptyString)
	assert.Equal(err, createInvalidParameterError(operationDeleteByID, parameterID))

	err = service.DeleteByID(whitespaceString)
	assert.Equal(err, createInvalidParameterError(operationDeleteByID, parameterID))
}

func TestCertificateServiceUpdateWithEmptyCertificate(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	certificate, err := service.Update(model.Certificate{})

	assert.Error(err)
	assert.Nil(certificate)
}

func TestCertificateServiceReplaceWithInvalidArguments(t *testing.T) {
	service := createCertificateService(t)
	assert := assert.New(t)

	certificate, err := service.Replace(emptyString, nil)

	assert.Equal(err, createInvalidParameterError(operationReplace, parameterCertificateID))
	assert.Nil(certificate)

	certificate, err = service.Replace(whitespaceString, nil)

	assert.Equal(err, createInvalidParameterError(operationReplace, parameterCertificateID))
	assert.Nil(certificate)

	certificate, err = service.Replace("fake-id-string", nil)

	assert.Equal(err, createInvalidParameterError(operationReplace, parameterReplacementCertificate))
	assert.Nil(certificate)

	replacementCertificate := model.NewReplacementCertificate("fake-name-string", "fake-password-string")

	assert.NotNil(replacementCertificate)

	certificate, err = service.Replace(whitespaceString, replacementCertificate)

	assert.Error(err)
	assert.Nil(certificate)
}

func createCertificateService(t *testing.T) *certificateService {
	service := newCertificateService(nil, TestURICertificates)
	testNewService(t, service, TestURICertificates, serviceCertificateService)
	return service
}
