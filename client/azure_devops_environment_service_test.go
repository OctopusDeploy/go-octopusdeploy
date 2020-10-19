package client

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewAzureDevOpService(t *testing.T) {
	serviceFunction := newAzureEnvironmentService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAzureEnvironmentService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *azureEnvironmentService
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
