package client

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewAuthenticationService(t *testing.T) {
	serviceFunction := newAuthenticationService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAuthenticationService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *authenticationService
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

func createAuthenticationService(t *testing.T) *authenticationService {
	service := newAuthenticationService(&sling.Sling{}, TestURIAuthentication)
	testNewService(t, service, TestURIAuthentication, serviceAuthenticationService)
	return service
}
