package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewBuildInformationService(t *testing.T) {
	serviceFunction := newBuildInformationService
	client := &sling.Sling{}
	uriTemplate := emptyString
	bulkPath := emptyString
	serviceName := serviceBuildInformationService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *buildInformationService
		client      *sling.Sling
		uriTemplate string
		bulkPath    string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, bulkPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, bulkPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, bulkPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.bulkPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func createBuildInformationService(t *testing.T) *buildInformationService {
	service := newBuildInformationService(nil, TestURIBuildInformation, TestURIBuildInformationBulk)
	testNewService(t, service, TestURIBuildInformation, serviceBuildInformationService)
	return service
}
