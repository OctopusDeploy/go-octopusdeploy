package buildinformation

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

func TestNewBuildInformationService(t *testing.T) {
	serviceFunction := NewBuildInformationService
	client := &sling.Sling{}
	uriTemplate := ""
	bulkPath := ""
	serviceName := constants.ServiceBuildInformationService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *BuildInformationService
		client      *sling.Sling
		uriTemplate string
		bulkPath    string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, bulkPath},
		{"EmptyURITemplate", serviceFunction, client, "", bulkPath},
		{"URITemplateWithWhitespace", serviceFunction, client, " ", bulkPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.bulkPath)
			services.NewServiceTests(t, service, uriTemplate, serviceName)
		})
	}
}
