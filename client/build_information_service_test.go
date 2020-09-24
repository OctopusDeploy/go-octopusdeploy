package client

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewBuildInformationService(t *testing.T) {
	serviceFunction := newBuildInformationService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceBuildInformationService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *buildInformationService
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

func createBuildInformationService(t *testing.T) *buildInformationService {
	service := newBuildInformationService(nil, TestURIBuildInformation)
	testNewService(t, service, TestURIBuildInformation, serviceBuildInformationService)
	return service
}
