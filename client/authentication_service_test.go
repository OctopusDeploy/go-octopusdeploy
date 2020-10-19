package client

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewAuthenticationService(t *testing.T) {
	serviceFunction := newAuthenticationService
	client := &sling.Sling{}
	uriTemplate := emptyString
	loginInitiatedPath := emptyString
	serviceName := serviceAuthenticationService

	testCases := []struct {
		name               string
		f                  func(*sling.Sling, string, string) *authenticationService
		client             *sling.Sling
		uriTemplate        string
		loginInitiatedPath string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, loginInitiatedPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, loginInitiatedPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, loginInitiatedPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.loginInitiatedPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func createAuthenticationService(t *testing.T) *authenticationService {
	service := newAuthenticationService(nil, TestURIAuthentication, TestURILoginInitiated)
	testNewService(t, service, TestURIAuthentication, serviceAuthenticationService)
	return service
}
