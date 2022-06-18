package authentication

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

func TestNewAuthenticationService(t *testing.T) {
	ServiceFunction := NewAuthenticationService
	client := &sling.Sling{}
	uriTemplate := ""
	loginInitiatedPath := ""
	ServiceName := constants.ServiceAuthenticationService

	testCases := []struct {
		name               string
		f                  func(*sling.Sling, string, string) *AuthenticationService
		client             *sling.Sling
		uriTemplate        string
		loginInitiatedPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, loginInitiatedPath},
		{"EmptyURITemplate", ServiceFunction, client, "", loginInitiatedPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", loginInitiatedPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.loginInitiatedPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
