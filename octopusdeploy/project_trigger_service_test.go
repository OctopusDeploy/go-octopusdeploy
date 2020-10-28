package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectTriggerService(t *testing.T) {
	ServiceFunction := newProjectTriggerService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceProjectTriggerService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *projectTriggerService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectTriggerServiceGetWithEmptyID(t *testing.T) {
	service := newProjectTriggerService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}
