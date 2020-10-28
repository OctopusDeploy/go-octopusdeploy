package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestArtifactService(t *testing.T) {
	service := newArtifactService(nil, TestURIArtifacts)
	testNewService(t, service, TestURIArtifacts, ServiceArtifactService)

	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceArtifactService

	testCases := []struct {
		name        string
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", nil, uriTemplate},
		{"EmptyURITemplate", client, emptyString},
		{"URITemplateWithWhitespace", client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := newArtifactService(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
