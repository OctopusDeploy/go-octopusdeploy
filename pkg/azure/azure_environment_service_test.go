package azure

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

func TestNewAzureEnvironmentService(t *testing.T) {
	ServiceFunction := NewAzureEnvironmentService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceAzureEnvironmentService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *AzureEnvironmentService
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
