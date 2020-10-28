package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewAzureDevOpService(t *testing.T) {
	ServiceFunction := newAzureEnvironmentService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceAzureEnvironmentService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *azureEnvironmentService
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
