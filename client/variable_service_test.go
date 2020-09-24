package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewVariableService(t *testing.T) {
	serviceFunction := newVariableService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceVariableService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *variableService
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

func TestVariableServiceGetAllWithEmptyID(t *testing.T) {
	service := newVariableService(&sling.Sling{}, emptyString)

	resource, err := service.GetAll(emptyString)

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.GetAll(whitespaceString)

	assert.Error(t, err)
	assert.Nil(t, resource)
}
