package runbooks

import (
	"testing"

	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

func TestRunbookProcessServiceNew(t *testing.T) {
	ServiceFunction := NewRunbookProcessService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceRunbookProcessService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *RunbookProcessService
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
