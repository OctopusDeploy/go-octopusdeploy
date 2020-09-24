package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestMachineService(t *testing.T) {
	t.Run("New", TestNewMachineService)
}

func TestNewMachineService(t *testing.T) {
	serviceFunction := newMachineService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceMachineService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *machineService
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

func TestMachineServiceGetWithEmptyID(t *testing.T) {
	service := newMachineService(&sling.Sling{}, emptyString)
	assert := assert.New(t)

	resource, err := service.GetByID(emptyString)

	assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(resource)
}
