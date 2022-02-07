package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectTriggerService(t *testing.T) {
	ServiceFunction := newProjectTriggerService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceProjectTriggerService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *projectTriggerService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectTriggerServiceGetWithEmptyID(t *testing.T) {
	service := newProjectTriggerService(&sling.Sling{}, services.emptyString)

	resource, err := service.GetByID(services.emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}
