package service

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewBuildInformationService(t *testing.T) {
	ServiceFunction := newBuildInformationService
	client := &sling.Sling{}
	uriTemplate := emptyString
	bulkPath := emptyString
	ServiceName := ServiceBuildInformationService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *buildInformationService
		client      *sling.Sling
		uriTemplate string
		bulkPath    string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, bulkPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, bulkPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, bulkPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.bulkPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func createBuildInformationService(t *testing.T) *buildInformationService {
	service := newBuildInformationService(nil, TestURIBuildInformation, TestURIBuildInformationBulk)
	testNewService(t, service, TestURIBuildInformation, ServiceBuildInformationService)
	return service
}
