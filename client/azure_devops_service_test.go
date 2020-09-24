package client

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewAzureDevOpsService(t *testing.T) {
	serviceFunction := NewAzureDevOpsService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAzureDevOpsService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *AzureDevOpsService
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

func createAzureDevOpsService(t *testing.T) *AzureDevOpsService {
	service := NewAzureDevOpsService(&sling.Sling{}, TestURIAzureDevOpsConnectivityCheck)
	testNewService(t, service, TestURIAzureDevOpsConnectivityCheck, serviceAzureDevOpsService)
	return service
}
