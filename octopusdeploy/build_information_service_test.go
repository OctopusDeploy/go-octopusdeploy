package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
)

func TestNewBuildInformationService(t *testing.T) {
	ServiceFunction := newBuildInformationService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	bulkPath := services.emptyString
	ServiceName := ServiceBuildInformationService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *buildInformationService
		client      *sling.Sling
		uriTemplate string
		bulkPath    string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, bulkPath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, bulkPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, bulkPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.bulkPath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func createBuildInformationService(t *testing.T) *buildInformationService {
	service := newBuildInformationService(nil, TestURIBuildInformation, TestURIBuildInformationBulk)
	services.testNewService(t, service, TestURIBuildInformation, ServiceBuildInformationService)
	return service
}
