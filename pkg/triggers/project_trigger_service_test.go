package triggers

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectTriggerService(t *testing.T) {
	ServiceFunction := NewProjectTriggerService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceProjectTriggerService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *ProjectTriggerService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectTriggerServiceGetWithEmptyID(t *testing.T) {
	service := NewProjectTriggerService(&sling.Sling{}, "")

	resource, err := service.GetByID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}
