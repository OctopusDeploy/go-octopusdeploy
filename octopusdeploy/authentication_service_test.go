package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
)

func createAuthenticationService(t *testing.T) *authenticationService {
	service := newAuthenticationService(nil, TestURIAuthentication, TestURILoginInitiated)
	testNewService(t, service, TestURIAuthentication, ServiceAuthenticationService)
	return service
}

func TestNewAuthenticationService(t *testing.T) {
	ServiceFunction := newAuthenticationService
	client := &sling.Sling{}
	uriTemplate := emptyString
	loginInitiatedPath := emptyString
	ServiceName := ServiceAuthenticationService

	testCases := []struct {
		name               string
		f                  func(*sling.Sling, string, string) *authenticationService
		client             *sling.Sling
		uriTemplate        string
		loginInitiatedPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, loginInitiatedPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, loginInitiatedPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, loginInitiatedPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.loginInitiatedPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
