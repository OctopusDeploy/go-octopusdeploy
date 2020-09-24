package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewLibraryVariableSetService(t *testing.T) {
	serviceFunction := newLibraryVariableSetService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceLibraryVariableSetService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *libraryVariableSetService
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

func TestLibraryVariableSetServiceGetWithEmptyID(t *testing.T) {
	service := createLibraryVariableSetService(t)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}

func TestLibraryVariableSetGetByPartialName(t *testing.T) {
	service := createLibraryVariableSetService(t)
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

func createLibraryVariableSetService(t *testing.T) *libraryVariableSetService {
	service := newLibraryVariableSetService(nil, TestURILibraryVariables)
	testNewService(t, service, TestURILibraryVariables, serviceLibraryVariableSetService)
	return service
}
