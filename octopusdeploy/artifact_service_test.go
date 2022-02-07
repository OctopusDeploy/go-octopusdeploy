package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
)

func TestArtifactService(t *testing.T) {
	service := newArtifactService(nil, TestURIArtifacts)
	services.testNewService(t, service, TestURIArtifacts, ServiceArtifactService)

	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceArtifactService

	testCases := []struct {
		name        string
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", nil, uriTemplate},
		{"EmptyURITemplate", client, services.emptyString},
		{"URITemplateWithWhitespace", client, services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := newArtifactService(tc.client, tc.uriTemplate)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
