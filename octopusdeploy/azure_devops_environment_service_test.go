package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
)

func TestNewAzureDevOpService(t *testing.T) {
	ServiceFunction := newAzureEnvironmentService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceAzureEnvironmentService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *azureEnvironmentService
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
